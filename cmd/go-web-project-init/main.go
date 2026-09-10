package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jasontconnell/go-web-project-init/data"
	"github.com/jasontconnell/go-web-project-init/process"
)

var live bool = false

func main() {
	repo := flag.String("file", "https://github.com/jasontconnell/go-web-project/archive/refs/heads/master.zip", "the zip file location")
	project := flag.String("project", "", "the project name")
	keyValues := flag.String("kvfile", "", "the key value file for replacements")
	dest := flag.String("dest", "", "output destination")
	flag.Parse()

	if *project == "" || *keyValues == "" || *dest == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var b []byte
	var err error
	if live {
		b, err = process.DownloadFile(*repo)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(len(b))
	} else {
		b, err = os.ReadFile("./go-web-project-master.zip")
		if err != nil {
			log.Fatal(err)
		}
	}

	folder, err := process.UnzipBytes(b)
	if err != nil {
		log.Fatal("unzip error", err)
	}

	replacements := []data.KeyValue{}
	replacements = append(replacements, data.KeyValue{Key: "go-web-project-master", Value: *project})
	replacements = append(replacements, data.KeyValue{Key: "go-web-project", Value: *project})

	fileKv, err := process.ParseKeyValue(*keyValues, "%")
	if err != nil {
		log.Fatal("key value error", err)
	}

	replacements = append(replacements, fileKv...)

	process.ReplaceAll(folder, replacements)

	err = process.WriteProject(folder, *dest)
	if err != nil {
		log.Fatal("couldn't write project", err)
	}

	printdir(folder, 0)
}

func printdir(folder *data.Folder, depth int) {
	prefix := strings.Repeat(" ", depth)
	fmt.Println(prefix + folder.Name + "/")

	for _, p := range folder.Folders {
		printdir(p, depth+1)
		for _, f := range p.Files {
			printfile(f, depth+2)
		}
	}

}

func printfile(file *data.File, depth int) {
	prefix := strings.Repeat(" ", depth)
	fmt.Println(prefix+file.Name, len(file.Contents))
}
