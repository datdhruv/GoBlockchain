package main

import (
	"os"

	"github.com/datDhruvJain/GOBlockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
