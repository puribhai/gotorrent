package bencode

import (
	"bufio"
	"errors"
	"io"
)

func decodeString(r *bufio.Reader) (Value, error) {
	length := 0

	for {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		if b == ':' {
			break
		}

		if b < '0' || b > '9' {
			return nil, errors.New("invalid string length")
		}

		length = length*10 + int(b-'0')
	}
	if length < 0 {
		return nil, errors.New("negative length")
	}

	if length > 10*1024*1024 {
		return nil, errors.New("string length exceeds max limit")
	}

	buff := make([]byte, length)
	n, err := io.ReadFull(r, buff)
	if err != nil {
		return nil, err
	}

	if n != length {
		return nil, io.ErrUnexpectedEOF
	}
	return string(buff), nil
}
