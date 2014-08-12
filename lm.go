package main

import (
	"os"
	"flag"
	"time"
	"github.com/kr/fs"
	"fmt"
)


type FilePath struct {
	MTime time.Time
	Path string
}


type ByMTime []FilePath

func (a ByMTime) Len() { return len(a) }
func (a ByMTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByMTime) Less(i, j int) bool { return a[i].MTime.Before(a[j].MTime) }


func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "."
	}
	filenodes := make([]FilePath, 0)

	walker := fs.Walk(root)
	for walker.Step() {
		filenodes = append(filenodes, FilePath{f.ModTime(), path})
	}

	fmt.Println(filenodes)
}
