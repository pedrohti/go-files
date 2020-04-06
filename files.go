package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/artdarek/go-unzip"
	// "github.com/kjk/lzmadec"
)

func main() {
	files := readFiles()
	decompress(files)
}

func readFiles() []string {
	var files []string

	file, err := ioutil.ReadDir("./")
	check(err)

	for _, f := range file {
		// fmt.Println(f.Name())
		files = append(files, f.Name())
	}

	return files
}

func decompress(files []string) {
	for _, file := range files {
		if strings.Contains(file, ".zip") || strings.Contains(file, ".7z") || strings.Contains(file, ".rar") {
			noExtension := strings.TrimSuffix(file, filepath.Ext(file))
			uz := unzip.New(file, ""+noExtension)
			err := uz.Extract()
			check(err)

		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
