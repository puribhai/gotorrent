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
	var prevKey string
	const maxDictSize = 10000

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
		if prevKey != "" && key < prevKey {
			return nil, errors.New("dictionary keys not sorted")
		}
		prevKey = key
		if _, exists := dict[key]; exists {
			return nil, errors.New("duplicate dictionary key")
		}
		if len(dict) > maxDictSize {
			return nil, errors.New("dictionary too large")
		}

		val, err := Decode(r)
		if err != nil {
			return nil, err
		}
		dict[key] = val
	}
	return dict, nil
}
