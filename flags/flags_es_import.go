package flags

import (
	"context"
	"encoding/json"
	"gin_gorm/global"
	"gin_gorm/models"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/sirupsen/logrus"
	"os"
)

func EsImport(path string) {
	byteData, err := os.ReadFile(path)
	if err != nil {
		global.Log.Error("EsImport ReadFile err:", err)
	}

	var response ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		logrus.Fatalf("%s err: %s", string(byteData), err.Error())
	}
	var esClient models.ArticleItem
	esClient.CreateIndexByJson(response.Index)
	var request bulk.Request
	for _, data := range response.Data {
		request = append(request, data.ID)
		request = append(request, data.Doc)
	}
	_, err = global.Es.Bulk().Index(response.Index).Request(&request).Do(context.Background())
	if err != nil {
		global.Log.Error("EsImport Bulk err:", err)
	}
	global.Log.Infof("Es数据添加成功,共添加 %d 条", len(response.Data))
}
