package bencode

import (
	"bufio"
	"errors"
)

func decodeInt(r *bufio.Reader) (Value, error) {
	_, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	num := int64(0)
	sign := int64(1)
	digits := 0

	b, err := r.Peek(1)
	if err != nil {
		return nil, err
	}

	if b[0] == '-' {
		sign = -1
		_, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	for {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		if b == 'e' {
			break
		}

		if b < '0' || b > '9' {
			return nil, errors.New("invalid integer")
		}

		num = num*10 + int64(b-'0')
		digits++
	}
	if digits == 0 {
		return nil, errors.New("empty integer")
	}
	if num == 0 && sign == -1 {
		return nil, errors.New("negative zero not allowed")
	}

	return num * sign, nil
}
