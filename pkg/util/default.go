package util

import (
	"github.com/rs/xid"
)

func NewID() string {
	return xid.New().String()
}
