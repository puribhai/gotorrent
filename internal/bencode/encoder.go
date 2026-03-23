package bencode

import (
	"bytes"
	"fmt"
	"sort"
)

func Encode(v Value) ([]byte, error) {
	var buf bytes.Buffer

	err := encodeValue(&buf, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encodeValue(buf *bytes.Buffer, v Value) error {
	switch val := v.(type) {
	case string:
		return encodeString(buf, val)
	case int:
		return encodeInt(buf, int64(val))
	case int64:
		return encodeInt(buf, val)
	case List:
		return encodeList(buf, val)
	case Dict:
		return encodeDict(buf, val)
	default:
		return fmt.Errorf("unsupported type %T", v)
	}
}

func encodeInt(buf *bytes.Buffer, v int64) error {
	_, err := fmt.Fprintf(buf, "i%de", v)
	return err
}

func encodeString(buf *bytes.Buffer, s string) error {
	_, err := fmt.Fprintf(buf, "%d:%s", len(s), s)
	return err
}

func encodeList(buf *bytes.Buffer, l List) error {
	buf.WriteByte('l')

	for _, v := range l {
		if err := encodeValue(buf, v); err != nil {
			return err
		}
	}

	buf.WriteByte('e')
	return nil
}

func encodeDict(buf *bytes.Buffer, d Dict) error {
	buf.WriteByte('d')

	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if err := encodeString(buf, k); err != nil {
			return err
		}

		if err := encodeValue(buf, d[k]); err != nil {
			return err
		}
	}

	buf.WriteByte('e')
	return nil
}
