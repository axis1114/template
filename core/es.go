package core

import (
	"fmt"
	"gin_gorm/global"
	"github.com/elastic/go-elasticsearch/v8"
)

func InitEs() *elasticsearch.TypedClient {
	dsn := global.Config.Es.Dsn()
	cfg := elasticsearch.Config{
		Addresses: []string{
			global.Config.Es.Dsn(),
		},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		lg.Fatal(fmt.Sprintf("[%s] es连接失败", dsn))
	}
	return es
}
