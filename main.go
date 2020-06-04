package main

import (
	"fmt"
	"math"
	"flag"
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Filename must be passed as an argument.")
		os.Exit(0)
	}
	filename := os.Args[1]
	imageFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}


}