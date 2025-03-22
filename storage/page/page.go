package page

type pageHeader struct {
	logSequenceNumber int64
	checkSum          int16
}

type pageItems struct {
}

type pageLayout struct {
	Header        *pageHeader //Header of the page
	Items         *pageItems  //All rows are here
	PageTotalSize int         //Size of page in bytes
}

func DefineType(i interface{}) string {
	switch i.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	default:
		return "unknown"
	}
}
