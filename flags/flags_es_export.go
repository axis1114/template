package flags

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_gorm/global"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"os"
	"time"
)

func EsExport(index string) {
	res, err := global.Es.Search().Index(index).Query(&types.Query{
		MatchAll: &types.MatchAllQuery{},
	}).Do(context.Background())
	if err != nil {
		global.Log.Error("EsExport err:", err.Error())
	}

	var data ESIndexResponse
	data.Index = index
	for _, hit := range res.Hits.Hits {
		item := Data{
			ID:  hit.Id_,
			Doc: hit.Source_,
		}
		data.Data = append(data.Data, item)
	}

	fileName := fmt.Sprintf("%s_%s.json", index, time.Now().Format("20060102"))
	file, _ := os.Create(fileName)

	byteData, _ := json.Marshal(data)
	_, err = file.Write(byteData)
	if err != nil {
		global.Log.Error("EsExport err:", err.Error())
		return
	}
	err = file.Close()
	if err != nil {
		global.Log.Error("EsExport err:", err.Error())
	}

	global.Log.Infof("索引 %s 导出成功  %s", index, fileName)
}
