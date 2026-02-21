package bencode

import "bufio"

func decodeDict(r *bufio.Reader) (Value, error) {
	r.ReadByte()

	dict := make(Dict)

	for {
		b, _ := r.Peek(1)
		if b[0] == 'e' {
			r.ReadByte()
			break
		}
		keyVal, err := decodeString(r)
		if err != nil {
			return nil, err
		}
		key := string(keyVal.([]byte))

		val, err := Decode(r)
		if err != nil {
			return nil, err
		}
		dict[key] = val
	}
	return dict, nil
}
