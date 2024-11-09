package highload

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init function...")
}

func ConnectSeveralClientsToDatabase(t *testing.T) {
	/*
		Testing gopostgres behaviour when two clients are trying to connect and execute some query.
	*/
}
