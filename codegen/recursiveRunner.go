package codegen

import (
	"fmt"
	"go/ast"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

type recursiveRunner struct {
	maxTry          int
	SubstringByNode func(n ast.Node) (string, error)
}

func (it *recursiveRunner) TypeName(n ast.Node) string {
	typeName := fmt.Sprintf("%T", n)

	if len(typeName) > 3 {
		return typeName[1 : len(typeName)-2]
	}

	return typeName
}

func (it *recursiveRunner) recursiveMapEntry(
	rawErrSlice errcore.RawErrCollection,
	curMap args.Map,
	n ast.Node,
) args.Map {
	if it.maxTry <= 0 || n == nil {
		return curMap
	}

	it.maxTry--
	toString, subsErr := it.SubstringByNode(n)

	if subsErr != nil {
		rawErrSlice.Add(subsErr)

		return curMap
	}

	typeName := it.TypeName(n)
	m, isFound := curMap[typeName]

	if isFound {
		m.(map[string]interface{})[toString] = make(map[string]interface{})
		toMap := m.(map[string]interface{})[toString].(map[string]interface{})

		ast.Inspect(
			n, func(node ast.Node) bool {
				if it.maxTry <= 0 {
					return false
				}

				if n == nil {
					return true
				}

				it.recursiveMapEntry(
					rawErrSlice,
					toMap,
					node,
				)

				return true
			},
		)
	} else {
		curMap[typeName] = make(map[string]interface{})
		toMap := curMap[typeName].(map[string]interface{})
		toMap[toString] = make(map[string]interface{})
		toMap = toMap[toString].(map[string]interface{})
		ast.Inspect(
			n, func(node ast.Node) bool {
				if it.maxTry <= 0 {
					return false
				}

				if n == nil {
					return true
				}

				it.recursiveMapEntry(
					rawErrSlice,
					toMap,
					node,
				)

				return true
			},
		)
	}

	return curMap
}
