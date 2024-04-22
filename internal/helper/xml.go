package helper

import (
	"fmt"
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

// XMLgetBaseNodes get first significant node in xml tree
func XMLgetBaseNodes(
	reader io.Reader, nodePath string) ([]*xmlquery.Node, error) {
	// Parse first xml tree
	xmlTree, err := xmlquery.Parse(reader)
	if err != nil {
		return nil, err
	}
	return xmlquery.Find(xmlTree, nodePath), nil
}

// XMLgetBaseNode get base node fro a simple xml tree (no multiple top nodes)
func XMLgetBaseNode(reader io.Reader, nodePath string) (*xmlquery.Node, error) {
	nodes, err := XMLgetBaseNodes(reader, nodePath)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, fmt.Errorf("no node found for path: %s", nodePath)
	}
	if len(nodes) > 1 {
		return nil, fmt.Errorf("not a simple xml, multiple nodes found for path: %s", nodePath)
	}
	return nodes[0], nil
}
