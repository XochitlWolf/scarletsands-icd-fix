// Code generated by "pipeline"; DO NOT EDIT.
package oldmercspal

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.WeaponData

func init() {
	base = &model.WeaponData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}
