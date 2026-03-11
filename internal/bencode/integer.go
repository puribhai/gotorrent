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

	sign := int64(1)
	var digits []byte

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
		digits = append(digits, b)
	}
	if len(digits) == 0 {
		return nil, errors.New("empty integer")
	}
	if len(digits) > 1 && digits[0] == '0' {
		return nil, errors.New("leading zero not allowed")
	}
	if len(digits) == 1 && sign == -1 && digits[0] == '0' {
		return nil, errors.New("negative zero not allowed")
	}

	num := int64(0)
	for _, d := range digits {
		num = num*10 + int64(d-'0')
	}

	return num * sign, nil
}
