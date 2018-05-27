package gobyexample

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
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
	projectDir = "/Users/admin/go/src/github.com/gmaclinuxer/travis-golang-example/"
)

func TestReadFiles(t *testing.T) {
	var inDatFile = filepath.Join(projectDir, "in.dat")

	fmt.Println(GetPwd())
	dat, err := ioutil.ReadFile(inDatFile)
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open(inDatFile)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

}

func Test_check(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.err)
		})
	}
}

func TestGetPwd(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPwd(); got != tt.want {
				t.Errorf("GetPwd() = %v, want %v", got, tt.want)
			}
		})
	}
}
