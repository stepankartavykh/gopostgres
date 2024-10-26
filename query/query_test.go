package query

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init function...")
}

func TestQueryParsing(t *testing.T) {
	query := "SELECT * from samples"
	response := HandleQuery(query)
	fmt.Println(response)
}

func TestQueryWrongSelects(t *testing.T) {
	fmt.Println("test wrong selects")
}
