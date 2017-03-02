package foundation

import "testing"
import "strconv"

func TestMapMerge(t *testing.T) {
	a := make(map[string]bool)
	b := make(map[string]bool)
	c := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		k := strconv.Itoa(i)
		a[k] = true
	}
	for i := 500; i < 1500; i++ {
		k := strconv.Itoa(i)
		b[k] = false
	}
	for i := 900; i < 1900; i++ {
		k := strconv.Itoa(i)
		c[k] = true
	}
	MapMerge(b, a)
	MapMerge(c, a)
}
