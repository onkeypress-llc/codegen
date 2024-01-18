package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

func StringOutput[D NodeWithStringOutput](data D) *NodeOutput[D] {
	instance := Output(cgtmp.New("string"), data)
	instance.toStringFunction = NodeObjectDataToString
	return instance
}
