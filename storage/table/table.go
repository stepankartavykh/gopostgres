package table

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const DataPath = "/home/skartavykh/GolandProjects/gopostgres/data/"

func generateTableStorageName() (string, int) {
	return strconv.Itoa(rand.Int()), 0
}

func CreateTable(tableName string) {
	fmt.Println("DEBUG: Creating table...", tableName)
	d1 := []byte("hello world")
	a, _ := generateTableStorageName()
	storageNamePath := DataPath + tableName + a

	err := os.WriteFile(storageNamePath, d1, 0644)
	check(err)
	f, err := os.Create(DataPath + "/dat2")
	check(err)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("ERROR: closing file with error!")
		}
	}(f)
}
