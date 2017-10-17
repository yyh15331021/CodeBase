package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type selpg_args struct {
	start       int
	end         int
	line        int
	find        bool
	destination string
}

func main() {
	var sa selpg_args
	flag.Usage = func() {
		fmt.Printf("Usage of %s: \n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.IntVar(&sa.start, "s", 0, "The start page")
	flag.IntVar(&sa.end, "e", 0, "The end page")
	flag.IntVar(&sa.line, "l", 72, "Lines per page")
	flag.StringVar(&sa.destination, "d", "", "The destination")
	flag.BoolVar(&sa.find, "f", false, "Find the delimiter")
	flag.Parse()

	situation_args(&sa)
	situation_cmd(&sa)

}
