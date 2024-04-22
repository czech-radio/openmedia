package helper

import (
	"io"
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
	baseNode, err := XMLgetBaseNode(reader, "root")
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
			got, _ := XMLnodeLevelUp(tt.args.node, tt.args.levelUp)
			ok := reflect.DeepEqual(got, tt.want)
			if !ok {
				t.Errorf(
					"NodeGetParent() = %v, want %v",
					got, tt.want)
			}
		})
	}
}

var ExampleXMLrundown = `<?xml version="1.0" encoding="utf-8" standalone="no"?>
<!DOCTYPE OPENMEDIA SYSTEM "ann_objects.dtd">
<OPENMEDIA>
  <OM_SERVER/>
  <OM_OBJECT SystemID="3fc88f5c-ef6b-44fa-bdef-002c69855f16" ObjectID="0000000200da1b00" DocumentURN="urn:openmedia:3fc88f5c-ef6b-44fa-bdef-002c69855f16:0000000200DA1B00" DirectoryID="00000002007dbae5" InternalType="1" TemplateID="fffffffa00001022" TemplateType="1" TemplateName="Radio Rundown">
    <OM_HEADER>
    </OM_HEADER>
  </OM_OBJECT>
</OPENMEDIA>
`

func TestXMLgetBaseNode(t *testing.T) {
	type args struct {
		breader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"openmedia_xml",
			args{strings.NewReader(ExampleXMLrundown)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := XMLgetBaseNode(
				tt.args.breader, "/OPENMEDIA")
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"XMLgetBaseNode() error = %v, wantErr %v",
					err, tt.wantErr,
				)
				return
			}
		})
	}
}
