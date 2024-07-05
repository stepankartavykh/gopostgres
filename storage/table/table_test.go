package table

import (
	"testing"
)

func TestTableCreation(t *testing.T) {
	tablesToCreate := []string{"a", "b", "c", "d"}
	for _, t := range tablesToCreate {
		CreateTable(t)
	}
}
