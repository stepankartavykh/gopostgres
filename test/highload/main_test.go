package highload

import (
	"fmt"
	"os/exec"
	"testing"
)

func init() {
	fmt.Println("init function...")
	/*
		Launch several clients in init function? Via telnet? Is it possible?
	*/
}

func TestConnectSeveralClientsToDatabase(t *testing.T) {
	/*
		Testing gopostgres behaviour when two clients are trying to connect and execute some query.
	*/

	fmt.Println("test print")
	app := "telnet"
	q := "localhost"
	w := "5555"

	cmd := exec.Command(app, q, w)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// go func() {
	// 	time.Sleep(time.Hour)
	// 	fmt.Print("qwe qw")
	// }()

	fmt.Println(string(stdout))
}
