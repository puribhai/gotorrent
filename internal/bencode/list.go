package bencode

import "bufio"

func decodeList(r *bufio.Reader) (Value, error) {
	_, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	var list List
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
		val, err := Decode(r)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
	}
	return list, nil
}
