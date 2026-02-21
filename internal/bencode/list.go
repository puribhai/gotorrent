package bencode

import "bufio"

func decodeList(r *bufio.Reader) (Value, error) {
	r.ReadByte()
	var list List
	for {
		b, _ := r.Peek(1)
		if b[0] == 'e' {
			r.ReadByte()
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
