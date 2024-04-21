package helper

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

// func TestNodeGetParent(t *testing.T) {
// 	reader := strings.NewReader(ExampleXmlLevels)
// 	baseNode, _ := xmlquery.Parse(reader)
// 	type args struct {
// 		node    *xmlquery.Node
// 		levelUp int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *xmlquery.Node
// 	}{
// 		{"A", args{baseNode, 0}, baseNode},
// 		{"B", args{baseNode.FirstChild, 1}, baseNode},
// 	}
// 	fmt.Println("kek", xmlquery.Find(baseNode, "/root"))
// for _, tt := range tests {
// 	t.Run(tt.name, func(t *testing.T) {
// 		got, _ := NodeGetParent(tt.args.node, tt.args.levelUp)
// 		ok := reflect.DeepEqual(got, tt.want)
// 		if !ok {
// 			t.Errorf(
// 				"NodeGetParent() = %v, want %v",
// 				got, tt.want)
// 		}
// 	})
// }
// }
