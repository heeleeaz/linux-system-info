package main

import (
	"data"
	"fmt"
)

func main() {
	repository := data.NewSystemRepository()
	fmt.Println(repository.DiskPartition())
}
