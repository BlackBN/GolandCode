package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var originalPath = "/Users/bn/Documents/Technology/Linux性能优化实战"

func main() {
	SearchFile(originalPath)
}

func SearchFile(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fileName := file.Name()
		// if fileName == ".DS_Store" {
		// 	continue
		// }
		absoluteFilepath := path + "/" + fileName
		fileInfo, err := os.Stat(absoluteFilepath)
		if err != nil {
			break
		}
		if fileInfo.IsDir() {
			// fmt.Println(absoluteFilepath)
			// err = os.Remove(absoluteFilepath)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			SearchFile(path + "/" + fileName)
		} else {
			if strings.HasSuffix(fileInfo.Name(), ".pdf") {
				fmt.Println(fileName)
				//os.Rename(absoluteFilepath, originalPath+"/"+fileName)
			} else {
				// err = os.Remove(absoluteFilepath)
				// if err != nil {
				// 	log.Fatal(err)
				// }
			}
		}
	}
}
