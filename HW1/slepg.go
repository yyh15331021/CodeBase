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
	inFile      string
}

var choice string

func StartWord()
{
	fmt.Printf("\nUsage: %s -s(start) -e(end) [ -f | -l(line) ][filename]\n",choice)
}

func situation_args(sa *selpg_args)
{
	if len(os.Args)<3 || (os.Args[1][0] != '-' || os.Args[1][1] != 's')
	|| (os.Args[2][0] != '-' || os.Args[2][1] != 'e') || (sa.start > sa.end || sa.start <0 ||sa.end <0) 
	{
		fmt.Printf("Wrong input! Try again")
	}
	if sa.find == false
	{
		if sa.line < 0
		{
			fmt.Printf("Wrong input! Try again")
		}
		else if sa.line == -1
		{
			sa.line = 72
		}
	}
}

func main() {
	var sa selpg_args
	choice = os.Args[0]
	
	flag.Usage = StartWord
	flag.IntVar(&sa.start, "s", 0, "The start page")
	flag.IntVar(&sa.end, "e", 0, "The end page")
	flag.IntVar(&sa.line, "l", 72, "Lines per page")
	flag.StringVar(&sa.destination, "d", "", "The destination")
	flag.BoolVar(&sa.find, "f", false, "Find the delimiter")
	flag.Parse()

	situation_args(&sa)
	situation_cmd(&sa)

}
