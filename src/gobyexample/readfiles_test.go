package gobyexample

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"io"
	"bufio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetPwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
	//return strings.Replace(dir, "\\", "/", -1)
}

const (
	projectDir = "/tmp/"
)

func TestReadFiles(t *testing.T) {
	var inDatFile = filepath.Join(projectDir, "in.dat")

	f, err := os.OpenFile(inDatFile, os.O_WRONLY|os.O_CREATE, 0666)
	check(err)
	f.WriteString("hello\ngo\n")
	f.Close()

	f, err = os.Open(inDatFile)
	check(err)

	defer f.Close()

	fmt.Println(GetPwd())
	dat, err := ioutil.ReadFile(inDatFile)
	check(err)
	fmt.Print(string(dat))

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

}
