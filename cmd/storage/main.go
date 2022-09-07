package main

import (
	"fmt"
	"log"

	"github.com/skyris/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It's uploaded", file)

	newFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Same file    ", newFile)

}
