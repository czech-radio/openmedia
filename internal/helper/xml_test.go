package helper

import (
	"reflect"
	"strings"
	"testing"

	"github.com/antchfx/xmlquery"
)

var ExampleXmlLevels = `
<root attr1="value1">
	<level1 attr2="value2">
		<level2 attr3="value3">
			<level3 attr4="value4">
				<level4 attr5="value5">Content of level 4</level4>
			</level3>
		</level2>
  </level1>
</root>
`

func TestNodeGetParent(t *testing.T) {
	// Prepare test pairs
	reader := strings.NewReader(ExampleXmlLevels)
	baseNode, err := XMLgetBaseNode(reader)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		node    *xmlquery.Node
		levelUp int
	}
	tests := []struct {
		name string
		args args
		want *xmlquery.Node
	}{
		{"up_zero_level", args{baseNode, 0}, baseNode},
		{"up_one_level", args{baseNode.FirstChild.NextSibling, 1}, baseNode},
		{"up_two_levels", args{baseNode.FirstChild.NextSibling.FirstChild, 2}, baseNode},
		{"up_three_levels", args{baseNode.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild, 3}, baseNode},
		{"up_more_levels", args{baseNode.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild, 10}, baseNode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NodeGetParent(tt.args.node, tt.args.levelUp)
			ok := reflect.DeepEqual(got, tt.want)
			if !ok {
				t.Errorf(
					"NodeGetParent() = %v, want %v",
					got, tt.want)
			}
		})
	}
}
