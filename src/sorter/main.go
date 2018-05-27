package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"github.com/gmaclinuxer/travis-golang-example/src/algs/qsort"
	"github.com/gmaclinuxer/travis-golang-example/src/algs/bubblesort"
	"time"
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
func writeDatFile(values []int, outfile string) error {

	file, err := os.Create(outfile)

	if err != nil {
		fmt.Printf("Fail to create output file: %s", outfile)
		return err
	}

	defer file.Close()

	//bw := bufio.NewWriter(file)

	for _, v := range values {
		str := strconv.Itoa(v)
		//bw.WriteString(str + "\n")
		file.WriteString(str + "\n")
	}

	//bw.Flush()

	return nil

}

func readDatFile(infile string) (values []int, err error) {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Printf("Failed to open input file: %s", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}

	return

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

	values, err := readDatFile(*inFile)

	if err == nil {
		fmt.Println("Read Values: ", values)

		t1 := time.Now()

		//values = []int{33, 3, 2, 6, 4, 3}
		fmt.Println("values: ", values, " len of values: ", len(values))

		switch *algName {
		case "qsort":
			qsort.QuickSort(values, 0, (len(values) - 1))
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Not support sort alg: ", algName)
		}

		t2 := time.Now()

		fmt.Println("Sort cost: ", t2.Sub(t1), "to complete")
		fmt.Println("Sort result: ", values)

		writeDatFile(values, *outFile)

	} else {
		fmt.Println("Read Error: ", err)
	}

}
