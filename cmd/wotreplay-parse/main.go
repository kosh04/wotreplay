package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kosh04/wotreplay"
)

var opt struct {
	output string
}

func init() {
	flag.Usage = func() {
		progname := filepath.Base(os.Args[0])
		println("Usage:", progname, "*.wotreplay")
		flag.PrintDefaults()
	}
	flag.StringVar(&opt.output, "out", "", "Write output replay data to `filename`")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		return
	}

	for _, file := range args {
		r, err := wotreplay.ParseFile(file)
		if err != nil {
			log.Println(file, err)
			continue
		}
		if opt.output != "" {
			err := ioutil.WriteFile(opt.output, r.File.Replay, os.ModePerm)
			if err != nil {
				log.Fatal(opt.output, err)
			}
		}
		fmt.Println(r)
	}
}
