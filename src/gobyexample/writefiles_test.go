package gobyexample

import (
	"path/filepath"
	"testing"
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

// TestWriteFiles https://gobyexample.com/writing-files
func TestWriteFiles(t *testing.T) {
	var outDatFile = filepath.Join("/tmp", "in.dat")
	d1 := []byte("hello\nworld\n")
	err := ioutil.WriteFile(outDatFile, d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{'h', 'e', 'l', 'o'}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("write %d bytes \n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("write %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("write %d bytes\n", n4)

	w.Flush()
}













