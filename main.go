package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type logWriter struct{}

func main() {

	dat, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(dat))

}
