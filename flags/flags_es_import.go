package flags

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/urfave/cli/v2"
	"os"
	"template/global"
	"template/models"
)

func EsImport(c *cli.Context) (err error) {
	path := c.String("path")
	byteData, err := os.ReadFile(path)
	if err != nil {
		global.Log.Error("EsImport ReadFile err:", err)
		return err
	}

	var response ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		global.Log.Fatalf("%s err: %s", string(byteData), err.Error())
		return err
	}
	var esClient models.ArticleItem
	esClient.CreateIndexByJson(response.Index)
	var request bulk.Request
	for _, data := range response.Data {
		request = append(request, map[string]interface{}{
			"index": map[string]interface{}{
				"_index": response.Index,
				"_id":    data.ID,
			},
		})
		request = append(request, data.Doc)
	}
	_, err = global.Es.Bulk().Index(response.Index).Request(&request).Do(context.Background())
	if err != nil {
		global.Log.Error("EsImport Bulk err:", err)
		return err
	}
	global.Log.Infof("Es数据添加成功,共添加 %d 条", len(response.Data))
	return nil
}
