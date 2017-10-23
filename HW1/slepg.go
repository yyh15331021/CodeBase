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
func process_input(sa *selpg_args)
{
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd
	if sa.destination != ""
	{
		cmd = exec.Command("Cat","-n")
		stdin, err = cmd.StdinPipe()
		if err != nil
		{
			fmt.Println(err)
			os.Exit(1)
		}
	}
	else
	{
		stdin = os.Stdout
	}
	if flag.NArg() == 0
	{
		fmt.Printf("\nUsage: %s Now the input will read from stdin, press (Ctrl+D)(Linux) | (Ctrl+C)(Windows) to exit\n",choice)
		ReadFromStdin(sa, stdin)
	}
	else
	{
		sa.inFile = flag.Arg(0)
		input_stream, err := os.Open(sa.inFile)
		if err != nil
		{
			fmt.Println(err)
			os.Exit(1)
		}
		input_buffer := bufio.NewReader(input_stream)
		if sa.find
		{
			type_f_process(sa, input_buffer, stdin)
		}
		else
		{
			type_l_process(sa, input_buffer, stdin)
		}
	}
	if sa.destination != ""
	{
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func type_f_process(sa *selpg_args, reader *bufio.Reader, stdin io.WriteCloser)
{
	for pageNum := 0; pageNum <= sa.end; pageNum++
	{
		Output, err := reader.ReadString('\f')
		if err == io.EOF
		{
			break
		}
		else if err != nil
		{
			fmt.Println(err)
			os.Exit(1)
		}
		if pageNum >= sa.start
		{
			OutputProcess(sa, string(Output), stdin)
		}
	}
}

func type_l_process(sa *selpg_args, reader *bufio.Reader, stdin io.WriteCloser)
{
	num := 0
	for
	{
		Output, _, err := reader.ReadLine()
		if err == io.EOF
		{
			break
		}
		else if err != nil
		{
			fmt.Println(err)
			os.Exit(1)
		}
		if num/sa.line >= sa.start
		{
			if num/sa.line <= sa.end
			{
				OutputProcess(sa, string(Output), stdin)
			}
			else
			{
				break
			}
		}
		num++
	}
}

func ReadFromStdin(sa *selpg_args, stdin io.WriteCloser)
{
	stdinInput := bufio.NewScanner(os.stdin)
	num := 0
	Output := ""
	temp := ""
	for StdinInput.Scan()
	{
		temp = StdinInput.Text() + "\n"
		if num/sa.line >= sa.start && num/sa.line < sa.end
		{
			Output += temp
		}
		num++
	}
	OutputProcess(sa, string(Output), stdin)
}

func OutputProcess(sa *selpg_args, Output string, stdin io.WriteCloser)
{
	if sa.destination != ""
	{
		stdin.Write([]byte(Output + "\n"))
	}
	else
	{
		fmt.Println(Output)
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
