package page

import "fmt"

type ResponseCodeOnWrite int

const (
	SUCCESS       ResponseCodeOnWrite = 1
	GENERAL_ERROR ResponseCodeOnWrite = 0
	PAGE_OVERFLOW ResponseCodeOnWrite = 2
)

type pageHeader struct {
	id                int64
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

type dataBlock struct {
	size int
}

func (page *pageLayout) writeInfo(data []dataBlock) ResponseCodeOnWrite {
	/*
		The question is how to write data in the page? What is stored in the dataBlock
	*/
	if len(data) == 0 {
		fmt.Println("No data to write to database.")
		return SUCCESS
	} else if len(data) > 1 {
		fmt.Printf("Number of blocks that would be writen is %v", len(data))
	}
	if len(data)*data[0].size > page.PageTotalSize {
		fmt.Printf("Can't write to page with page header %v", page.Header.id)
		return PAGE_OVERFLOW
	}
	writeProcedure()
	return SUCCESS
}

func writeProcedure() {
	panic("unimplemented")
}

func (page *pageLayout) readInfo() {

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
