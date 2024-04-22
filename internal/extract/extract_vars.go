package extract

type OtputFileTypeCode string

const (
	OtputFileTypeCSV  = "csv"
	OtputFileTypeXLSX = "xlsx"
)

var OutputFileTypeMap = map[string]OtputFileTypeCode{}

// type FileOutputTypeCode int

// const (
// OutputFileTypeCSV =
// )
