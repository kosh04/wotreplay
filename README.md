# WoTReplay

[![Build Status](https://travis-ci.org/kosh04/wotreplay.svg?branch=master)](https://travis-ci.org/kosh04/wotreplay)
[![GoDoc](https://godoc.org/github.com/kosh04/wotreplay?status.svg)](https://godoc.org/github.com/kosh04/wotreplay)

[World of Tanks](https://worldoftanks.com/) Replay Parser (written in Go)

*This package is work in progress*

## Usage

### command-line

    $ wotreplay-parse [-out FILE] *.wotreplay

    $ wotreplay-parse somefile.wotreplay
    {
	    "Date": "08.06.2017 04:52:35",
	    "Map": "聖なる谷",
	    "Gamemode": "Random",
	    "Reason": "Extermination",
	    "Tank": "ussr-R41_T-50",
	    "Player": "kosh04",
	    "Version": "World of Tanks v.0.9.19.0.1 #450"
    }

### Go code

    import (
        "fmt"
    	"github.com/kosh04/go-wotreplay"
    )
    
    r, err := wotreplay.ParseFile("FILE.wotreplay")
    if err != nil {
        // ParseError
    }
    fmt.Println(r)


## TODO

- [ ] Define JSON data type of Matchstart, ButtleResult, etc
- [ ] Analyze packet
- [ ] Friendly command-line
- [ ] [wotrp2j.py](https://github.com/Phalynx/WoT-Replay-To-JSON) compatible JSON
- [ ] Enable parse of old replays

## Similar Projects

- https://github.com/evido/wotreplay-parser (C++)
- https://github.com/Aimdrol/WoT-Replay-Analyzer (exe only)
- https://github.com/Phalynx/WoT-Replay-To-JSON (Python)
- https://gist.github.com/benvanstaveren/2402016 (Perl)

## License

MIT License
