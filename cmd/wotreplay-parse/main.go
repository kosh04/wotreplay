package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kosh04/wotreplay"
)

// App info
var (
	Name    = "wotreplay-parse"
	Version = "v0.1.1"
)

var opt struct {
	output  string
	version bool
}

func init() {
	flag.Usage = func() {
		println("Usage:", Name, "*.wotreplay")
		flag.PrintDefaults()
	}
	flag.StringVar(&opt.output, "out", "", "Write output replay data to `filename`")
	flag.BoolVar(&opt.version, "version", false, "Print version")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if opt.version {
		println(Name, Version)
		return
	}
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
