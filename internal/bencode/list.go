package bencode

import (
	"bufio"
	"errors"
)

func decodeList(r *bufio.Reader) (Value, error) {
	_, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	var list List
	const maxListElements = 10000
	for {
		b, err := r.Peek(1)
		if err != nil {
			return nil, err
		}
		if b[0] == 'e' {
			_, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			break
		}
		if len(list) >= maxListElements {
			return nil, errors.New("list exceeds maximum element limit")
		}
		val, err := Decode(r)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
	}
	return list, nil
}
