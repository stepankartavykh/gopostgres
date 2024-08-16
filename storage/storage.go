package storage

import (
	"fmt"
	"mypostgres/storage/table"
)

const StoragePath string = "storage/GOPOSTGRES_DATA"

type Storage struct{}

type Page struct {
	name    string
	path    string
	storage *Storage
}

type SystemFile struct {
}

type processType struct {
	debugMessage string
}

func storageWorker(process processType) {
	fmt.Println(process.debugMessage, "[DEBUG]: STORAGE location:", StoragePath, "\n[DEBUG]:----Worker for storage element is running...")
}

func CreateTable() {
	table.CreateTable("testTableName")
}

func TestStorage() {
	testProcess := processType{debugMessage: "TEST message for storage package"}
	storageWorker(testProcess)
}
