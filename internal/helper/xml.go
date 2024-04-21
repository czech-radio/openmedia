package helper

import (
	"github.com/antchfx/xmlquery"
)

/**
 * Function: NodeGetParent___
 *
 * @param node:    ___
 * @param levelUp: ___
 * @return: [*xmlquery.Node, bool] ___
 */
func NodeGetParent(node *xmlquery.Node, levelUp int) (*xmlquery.Node, bool) {
	resultNode := node
	if levelUp == 0 {
		return resultNode, true
	}
	for i := 0; i <= levelUp; i++ {
		resultNode = resultNode.Parent
	}
	return resultNode, true
}
