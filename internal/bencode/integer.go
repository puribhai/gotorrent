package bencode

import "bufio"

func decodeInt(r *bufio.Reader) (Value, error) {
	r.ReadByte()

	var num int64
	sign := int64(1)
	b, _ := r.Peek(1)

	if b[0] == '-' {
		sign = -1
		r.ReadByte()
	}

	for {
		b, _ := r.ReadByte()
		if b == 'e' {
			break
		}
		num = num*10 + int64(b-'0')
	}
	return num * sign, nil
}
