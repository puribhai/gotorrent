package torrent

import (
	"crypto/sha1"

	"github.com/puribhai/gotorrent/internal/bencode"
)

func GenerateInfoHash(info bencode.Value) ([20]byte, error) {
	encoded, err := bencode.Encode(info)
	if err != nil {
		return [20]byte{}, err
	}

	hash := sha1.Sum(encoded)
	return hash, nil
}
