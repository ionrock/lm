package main

import (
	"os"
	"flag"
	"time"
	"github.com/kr/fs"
	"fmt"
	"sort"
)


type FilePath struct {
	MTime time.Time
	Path string
}


type ByMTime []FilePath

func (a ByMTime) Len() int {
	return len(a)
}

func (a ByMTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByMTime) Less(i, j int) bool {
	return a[i].MTime.Before(a[j].MTime)
}


func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "."
	}
	filenodes := make(ByMTime, 0)

	walker := fs.Walk(root)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		path := walker.Path()

		f, err := os.Stat(path)

		if err == nil && !f.IsDir() {
			filenodes = append(filenodes, FilePath{f.ModTime(), path})
		}
	}
	sort.Sort(sort.Reverse(filenodes))

	for _, node := range filenodes {
		fmt.Printf("[%s] %s\n", node.MTime, node.Path)
	}
}
