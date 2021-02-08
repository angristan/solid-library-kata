package main

import (
	"flag"
	"solid-library-kata/internal/cli"
)

func main() {
	f := cli.CommandFlags{}
	f.Action = flag.String("action", "", "action to do")
	f.Book = flag.String("book", "", "action to do")
	f.Author = flag.String("author", "", "action to do")
	f.User = flag.String("user", "", "action to do")
	f.Role = flag.String("role", "", "action to do")
	flag.Parse()
	cli.DispatchCommand(f)
}
