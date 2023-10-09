package main

import (
	"fmt"
	"log"
	"os"
)

func fileSize() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Size())

}
