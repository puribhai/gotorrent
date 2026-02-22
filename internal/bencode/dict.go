package bencode

import (
	"bufio"
	"errors"
)

func decodeDict(r *bufio.Reader) (Value, error) {
	_, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	dict := make(Dict)

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
		keyVal, err := decodeString(r)
		if err != nil {
			return nil, err
		}
		key, ok := keyVal.(string)
		if !ok {
			return nil, errors.New("dictionary key must be string")
		}

		val, err := Decode(r)
		if err != nil {
			return nil, err
		}
		dict[key] = val
	}
	return dict, nil
}
