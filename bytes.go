package foundation

import (
	"bufio"
	"bytes"
	"io"
)

type tEnumerableBytes struct {
	Bytes      []byte
	FilterFunc func(item interface{}) bool
}

func (s *tEnumerableBytes) Filter(f func(item interface{}) bool) IEnumerable {
	s.FilterFunc = f
	return s
}

func (s *tEnumerableBytes) Enumerate(f func(item interface{}, err error, stop *bool)) {
	if f == nil {
		return
	}

	aBufReader := bufio.NewReader(bytes.NewReader(s.Bytes))
	stop := false
	for !stop {
		b, _, err := aBufReader.ReadLine()
		if err != nil {
			if err != io.EOF {
				f("", err, &stop)
			}
			break
		}
		line := string(b)
		if s.FilterFunc != nil && s.FilterFunc(line) {
			continue
		}
		f(string(b), nil, &stop)
	}
}

func NewEnumerableBytes(b []byte) IEnumerable {
	aEnumerableBytes := tEnumerableBytes{Bytes: b}
	return &aEnumerableBytes
}
