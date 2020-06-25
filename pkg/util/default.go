package util

import (
	"encoding/json"
	"github.com/rs/xid"
)

func NewID() string {
	return xid.New().String()
}

func Stringify(obj interface{}) string {
	if obj == nil {
		return ""
	}

	if b, err := json.Marshal(obj); err != nil {
		return ""
	} else {
		return string(b)
	}
}
