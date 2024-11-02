package table

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type ErrorWithCreatingTableInStorage struct {
	arg     int
	message string
}

func (err *ErrorWithCreatingTableInStorage) Error() string {
	return fmt.Sprintf("%d - %s", err.arg, err.message)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var DataPath = os.Getenv("DATA_PATH")

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
