package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wu8685/htmlizer.git/api"
)

var junitXmlPath string
var outputPath string

func init() {
	flag.StringVar(&junitXmlPath, "i", "", "The path pointing to junit.xml files.")
	flag.StringVar(&outputPath, "o", "", "The path pointing to a dir or a output file.")
}

func main() {
	flag.Parse()
	if junitXmlPath == "" {
		fmt.Printf("No junit.xml path is provided.\n")
		return
	}

	junitXmlPath, err := filepath.Abs(junitXmlPath)
	if err != nil {
		fmt.Printf("No valid junit.xml path is provided: %s.\n", err.Error())
		return
	}

	if outputPath == "" {
		fmt.Printf("No output path instructed. Use current dir as default.")
		_outputPath, _err := os.Getwd()
		if _err != nil {
			fmt.Printf("No valid output dir is provided: %s.\n", err.Error())
			return
		}
		outputPath = _outputPath
	}

	if _, err = os.Stat(outputPath); os.IsNotExist(err) {
		err = genDir(outputPath)
		if err != nil {
			fmt.Printf("No valid output dir is provided: %s.\n", err.Error())
			return
		}
	} else if stat, _err := os.Stat(outputPath); _err != nil || stat.Mode().IsRegular() {
		fmt.Printf("Output path should be a directory.\n")
		return
	}

	fmt.Printf("Search with pattern %s .\n", junitXmlPath)
	matches, err := filepath.Glob(junitXmlPath)
	if err != nil {
		fmt.Printf("No valid junit.xml path is provided: %s.\n", err.Error())
		return
	}

	fmt.Printf("Filter out files: %v.\n", matches)

	t := api.Template()

	for _, f := range matches {
		if !strings.HasSuffix(f, ".xml") {
			continue
		}

		fmt.Printf("Start working on file %s .\n", f)

		testSuite, err := api.ParseJunitXML(f)
		if err != nil {
			fmt.Println("err : " + err.Error())
			continue
		}

		report := api.Aggregate(testSuite)

		out, err := genFileUnderDir(outputPath, f)
		if err != nil {
			fmt.Printf("Fail to create output file for junit xml %s cause err: %s\n", f, err.Error())
			continue
		}
		defer out.Close()

		err = api.ApplyTemplate(t, out, report)
		if err != nil {
			fmt.Printf("Fail to generate junit html report: %s .\n", err.Error())
		} else {
			fmt.Printf("Succeed to generate junit html report: %s .\n", out.Name())
		}
	}
}

func genDir(path string) error {
	var mode os.FileMode = 0777
	return os.MkdirAll(path, mode)
}

func genFileUnderDir(dir, inputPath string) (*os.File, error) {
	base := filepath.Base(inputPath)
	fileName := base
	if strings.HasSuffix(base, ".xml") {
		fileName = base[:len(base)-4]
	}

	basePath := dir + string(filepath.Separator) + fileName
	for {
		path := basePath + ".html"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return os.Create(path)
		}
		basePath = basePath + "_1"
	}
	return nil, errors.New("Fail to create output file under dir " + dir)
}
