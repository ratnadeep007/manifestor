package utils

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(name string, data []byte) {
	f, err := os.Create(name + ".yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done creating file")
}
