package main

import (
	"flag"
	"fmt"
	"os"
)

var name1 string

func init() {
	// 深层定制
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name1, "name", "everyone", "The greeting object.")
}

func main() {
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name1)
}
