package main

import (
	"flag"
	"fmt"
)

var (
	Version     string
	BuildCommit string
	BuildTime   string
	GoVersion   string
)

// printVersion print version info from makefile
func printVersion() {
	fmt.Printf(`Version:      %s
Go version:   %s
Git commit:   %s
Built:        %s
`, Version, GoVersion, BuildCommit, BuildTime)
}

func main() {

	inFile := flag.String("i", "in.dat", "input file")
	outFile := flag.String("o", "out.dat", "output file")
	algName := flag.String("a", "qsort", "qsort/bsort")
	version := flag.Bool("v", false, "print version info")

	flag.Parse()

	if *version {
		printVersion()
		return
	}

	fmt.Printf("inFile: %s, outFile: %s, algName: %s\n",
		*inFile, *outFile, *algName)

}
