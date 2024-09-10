package page

type pageHeader struct {
	logSequenceNumber int64
	checkSum          int16
}

type pageItems struct {
}

type pageLayout struct {
	Header        *pageHeader `json:"Header of the page"`
	Items         *pageItems  `json:"All rows are here"`
	PageTotalSize int         `json:"size of page in bytes json"`
}
