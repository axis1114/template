package flags

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"template/models"
)

type Data struct {
	ID  *string         `json:"id"`
	Doc json.RawMessage `json:"doc"`
}

type ESIndexResponse struct {
	Index string `json:"index"`
	Data  []Data `json:"data"`
}

func EsIndexCreate(c *cli.Context) (err error) {
	var es models.ArticleItem
	es.CreateIndex()
	return nil
}
