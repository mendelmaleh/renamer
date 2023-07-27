package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func start() error {
	// flags
	opts := struct {
		glob *string
		sort *string
		name *string
		yes  *bool
	}{
		flag.String("glob", "*", "files to rename"),
		flag.String("sort", "", "regex for file number"),
		flag.String("name", "", "regex for file name"),
		flag.Bool("yes", false, "rename"),
	}

	flag.Parse()

	resort := regexp.MustCompile(*opts.sort)
	rename := regexp.MustCompile(*opts.name)

	matches, err := filepath.Glob(*opts.glob)
	if err != nil {
		return err
	}

	for _, match := range matches {
		sorting := clean(resort.FindStringSubmatch(match)[1])
		naming := clean(rename.FindStringSubmatch(match)[1])
		name := sorting + "-" + naming + filepath.Ext(match)

		if *opts.yes {
			if err := os.Rename(match, name); err != nil {
				return err
			}
		} else {
			fmt.Println(name + " <- " + match)
		}
	}

	return nil
}

func clean(s string) string {
	s = strings.ToLower(s)
	s = strings.Join(strings.Fields(s), "_")

	return s
}
