package flags

import "encoding/json"

type Data struct {
	ID  *string         `json:"id"`
	Doc json.RawMessage `json:"doc"`
}

type ESIndexResponse struct {
	Index string `json:"index"`
	Data  []Data `json:"data"`
}
