package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: sha1sum [file ...]\n")
	os.Exit(2)
}

func error(s string) {
	fmt.Fprint(os.Stderr, s, "\n")
	os.Exit(1)
}

func sum(f *os.File, path string) {
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %s\n", path, err)
		return
	}
	if path == "" {
		fmt.Printf("%x\n", h.Sum(nil))
	} else {
		fmt.Printf("%x\t%s\n", h.Sum(nil), path)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		sum(os.Stdin, "")
	} else {
		for _, path := range args {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open %s: %s\n", path, err)
				continue
			}
			sum(f, path)
			f.Close()
		}
	}
	os.Exit(0)
}
