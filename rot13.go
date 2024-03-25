package rot13

import "io"

func Rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func Rot13Bytes(b []byte) []byte {
	var rot []byte
	for _, si := range b {
		rot = append(rot, Rot13(si))
	}
	return rot
}

type Rot13Reader struct {
	r io.Reader
}

func (r Rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = Rot13(p[i])
	}
	return
}

func Rot13String(s string) string {
	return string(Rot13Bytes([]byte(s)))
}
