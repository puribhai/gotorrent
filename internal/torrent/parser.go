package torrent

import (
	"bufio"
	"os"

	"github.com/puribhai/gotorrent/internal/bencode"
)

func ParseTorrent(path string) (*Torrent, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	value, err := bencode.Decode(reader)
	if err != nil {
		return nil, err
	}

	root := value.(bencode.Dict)

	info := root["info"].(bencode.Dict)

	t := &Torrent{}

	t.Announce = string(root["announce"].(string))
	t.Name = string(info["name"].(string))
	t.Length = int(info["length"].(int64))
	t.PieceLength = int(info["piece length"].(int64))

	piecesRaw := []byte(info["pieces"].(string))

	for i := 0; i < len(piecesRaw); i += 20 {
		t.Pieces = append(t.Pieces, piecesRaw[i:i+20])
	}
	return t, nil
}
