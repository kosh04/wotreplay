package wotreplay

// PacketReader is wotreplay packet reader
// https://github.com/evido/wotreplay-parser/blob/master/src/packet_reader_80.cpp
type PacketReader struct {
	Buffer  string
	Version string
	Title   string
	Pos     int
}

func (pr PacketReader) next()    {}
func (pr PacketReader) hasNext() {}
