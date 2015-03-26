package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var rflag = flag.Bool("r", false, "recursive")

func main() {
	elog := log.New(os.Stderr, "rm: ", 0)
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: rm [options] file ...")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	for _, path := range flag.Args() {
		if *rflag {
			err := os.RemoveAll(path)
			if err != nil {
				elog.Print(err)
			}
		} else {
			err := os.Remove(path)
			if err != nil {
				elog.Print(err)
			}
		}
	}
}
