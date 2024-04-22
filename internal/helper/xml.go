package helper

import (
	"io"

	"github.com/antchfx/xmlquery"
)

// XMLnodeLevelUp go up in xmltree acording to level. If level greater than top most node, return top most node
func XMLnodeLevelUp(node *xmlquery.Node, levelUp int) (*xmlquery.Node, int) {
	var levelUpCount int
	resultNode := node
	if levelUp == 0 {
		return resultNode, levelUpCount
	}
	for i := 0; i < levelUp; i++ {
		subRes := resultNode.Parent
		if subRes.Parent == nil {
			break
		}
		levelUpCount++
		resultNode = subRes
	}
	return resultNode, levelUpCount
}

// XMLgetBaseNode get first significant node in xml tree
func XMLgetBaseNode(
	breader io.Reader) (*xmlquery.Node, error) {
	// Parse first xml tree
	xmlNode, err := xmlquery.Parse(breader)
	if err != nil {
		return nil, err
	}
	node := xmlNode.FirstChild.NextSibling
	return node, nil
}
