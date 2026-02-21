package bencode

import "bufio"

func decodeString(r *bufio.Reader) (Value, error) {
	length := 0

	for {
		b, _ := r.ReadByte()
		if b == ':' {
			break
		}
		length = length*10 + int(b-'0')
	}
	buff := make([]byte, length)
	_, err := r.Read(buff)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}
