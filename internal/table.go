package internal

type ObjectAttributes = map[string]string
type Fields = map[int]string       // FieldID vs value
type UniqueValues = map[string]int // value vs count

type Table struct {
	ObjectHeader []string
	Rows         []Fields
}
