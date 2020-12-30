package waboorrt

import (
	"github.com/flipdot/waboorrt-go/types"
)

type EntityType string

const (
	EntityTypeBot = "BOT"
)

type Entity struct {
	types.Pos
	Type EntityType `json:"type"`
}
