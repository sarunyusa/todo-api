package model

import "encoding/json"

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *CommonResponse) ToJsonBytes() []byte {
	bs, err := json.Marshal(r)
	if err != nil {
		return []byte("")
	} else {
		return bs
	}
}
