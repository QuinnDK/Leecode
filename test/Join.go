package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Println(filepath.Join("a", "b", "c"))
	f:=FindPhoneNumber("1.txt")
	fmt.Printf("%s",f)
}
func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)

}
