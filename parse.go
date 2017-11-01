package wotreplay

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/blowfish"
)

var (
	errInvalidFile = errors.New("wotreplay: invalid file")
)

// ParseFile decode wotreplay file from FILENAME
func ParseFile(filename string) (*Replay, error) {
	// if !IsReplayFile(filename) {
	// 	return errInvalidFile
	// }

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parse(f)
}

func parse(in io.Reader) (*Replay, error) {
	r := &Replay{}
	r.Logger = log.New(os.Stderr, "", log.Lshortfile)

	if err := binary.Read(in, binary.LittleEndian, &r.File.Header); err != nil {
		return nil, err
	}
	if r.File.Header.Magic != magic {
		return nil, errInvalidFile
	}

	r.File.Blocks = make([]dataBlock, r.File.Header.BlockCount)

	if r.File.Header.BlockCount >= 1 {
		if err := parseBlock(in, &r.File.Blocks[0], &r.MatchStart); err != nil {
			return nil, err
		}
	}
	if r.File.Header.BlockCount >= 2 {
		var dat []json.RawMessage
		if err := parseBlock(in, &r.File.Blocks[1], &dat); err != nil {
			return nil, err
		}
		json.Unmarshal(dat[0], &r.BattleResults.Results)
		json.Unmarshal(dat[1], &r.BattleResults.Vehicles)
		json.Unmarshal(dat[2], &r.BattleResults.Data)
	}
	if r.File.Header.BlockCount >= 3 {
		// FIXME: How is "Match End" block used?
		if err := binary.Read(in, binary.LittleEndian, &r.File.Blocks[2].Length); err != nil {
			return nil, err
		}
		r.File.Blocks[2].Data = make([]byte, r.File.Blocks[2].Length)
		if err := binary.Read(in, binary.LittleEndian, &r.File.Blocks[2].Data); err != nil {
			return nil, err
		}
	}

	// FIXME: Unknown fields (4 + 4 bytes)
	if err := binary.Read(in, binary.LittleEndian, &r.File.ReplayLen); err != nil {
		return nil, err
	}
	if err := binary.Read(in, binary.LittleEndian, &r.File.XXX); err != nil {
		return nil, err
	}

	// Read real replay data
	raw, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	p, err := unpackReplay(raw)
	if err != nil {
		return nil, err
	}
	if r.File.ReplayLen != uint32(len(p)) {
		r.Logger.Printf("WARNING: Unpack %d bytes (expected %d bytes)\n", len(p), r.File.ReplayLen)
	}
	r.File.Replay = make([]byte, len(p))
	copy(r.File.Replay, p)

	return r, nil
}

func parseBlock(in io.Reader, b *dataBlock, v interface{}) error {
	if err := binary.Read(in, binary.LittleEndian, &b.Length); err != nil {
		return err
	}
	b.Data = make([]byte, b.Length)
	if err := binary.Read(in, binary.LittleEndian, &b.Data); err != nil {
		return err
	}
	return json.Unmarshal(b.Data, v)
}

var encryptionKey = []byte("\xDE\x72\xBE\xA0\xDE\x04\xBE\xB1\xDE\xFE\xBE\xEF\xDE\xAD\xBE\xEF")

func unpackReplay(raw []byte) ([]byte, error) {
	block, err := blowfish.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	var prev struct {
		decrypted bool
		data      [8]byte
	}
	out := make([]byte, 0, len(raw))
	r := bytes.NewReader(raw)
	for {
		var b [8]byte
		n, err := r.Read(b[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for i := n; i < 8; i++ {
			b[i] = 0x0
		}

		block.Decrypt(b[:], b[:])

		if prev.decrypted {
			for i := range b {
				b[i] ^= prev.data[i]
			}
		}
		copy(prev.data[:], b[:])
		prev.decrypted = true

		out = append(out, b[:]...)
	}

	return decompress(out)
}

func decompress(b []byte) ([]byte, error) {
	// First two bytes should be : 78 DA. ("Best Compression" Header)
	if !bytes.HasPrefix(b, []byte("\x78\xDA")) {
		log.Printf("WARNING: Unexpected Compressed Header: %#v (want 78 DA) \n", b[0:2])
	}

	z, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	return ioutil.ReadAll(z)
}
