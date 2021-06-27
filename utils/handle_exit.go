package utils

import "os"

func GraceFullExit(err error) {
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(0)
		}
	}
}
