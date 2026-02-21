package bencode

import (
	"bufio"
	"fmt"
)

func Decode(r *bufio.Reader) (Value, error) {
	b, err := r.Peek(1)
	if err != nil {
		return nil, err
	}
	switch b[0] {
	case 'i':
		return decodeInt(r)
	case 'l':
		return decodeList(r)
	case 'd':
		return decodeDict(r)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return decodeString(r)
	default:
		return nil, fmt.Errorf("invalid bencode type: %c", b[0])
	}
}
