package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(1)
	//if err != nil {
	//	panic(err)
	//}
}

func MakeUUID() string {

	return node.Generate().String()

}
