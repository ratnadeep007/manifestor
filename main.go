package main

import (
	"fmt"
	"manifest_creator/cli"
	"os"
)

func main() {
	if err := os.Args[1:]; err != nil {
		if os.Args[1] == "--version" {
			fmt.Println("Manifestor Alpha v1")
			os.Exit(0)
		}
	}

	cli.Run()
}
