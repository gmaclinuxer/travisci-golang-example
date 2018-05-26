package main

import (
	"flag"
	"fmt"
)


//var Usage = func() {
//	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
//	flag.PrintDefaults()
//}


func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("numb: ", *numbPtr)
	fmt.Println("fork: ", *boolPtr)
	fmt.Println("svar: ", svar)

	fmt.Println("tail:", flag.Args())

}
