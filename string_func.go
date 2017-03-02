package foundation

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ---- Basic ----

// StrMD5 return md5 string of s
func StrMD5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// StrRandomMD5 return a md5 random string
func StrRandomMD5() string {
	prefix := StrRandom(64, OptionRandomStringAll)
	suffix := StrRandom(64, OptionRandomStringAll)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return StrMD5(prefix + timestamp + suffix)
}

// StrRandom return a random string
func StrRandom(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// StrEnumerate enumerate s by callback function f
func StrEnumerate(s string, sep string, reversal bool, f func(surplus string, current string, stop *bool)) {
	if f == nil {
		return
	}
	stop := false
	if s == "" || sep == "" {
		f("", s, &stop)
		return
	}

	surplus := s
	for !stop && surplus != "" {
		var idx int
		if reversal {
			idx = strings.LastIndex(surplus, sep)
		} else {
			idx = strings.Index(surplus, sep)
		}
		if idx < 0 {
			f("", surplus, &stop)
			break
		} else {
			var current string
			if reversal {
				current = surplus[idx+1:]
				surplus = surplus[:idx]
			} else {
				current = surplus[:idx]
				surplus = surplus[idx+1:]
			}
			f(surplus, current, &stop)
		}
	}
}

// StrSplitFirst return fist string item splited by sep
func StrSplitFirst(s string, sep string) string {
	if s == "" && sep == "" {
		return s
	}
	idx := strings.Index(s, sep)
	if idx > -1 {
		return s[:idx]
	}
	return s
}

// StrSplitLast return last string item splited by sep
func StrSplitLast(s string, sep string) string {
	if s == "" && sep == "" {
		return s
	}
	idx := strings.LastIndex(s, sep)
	if idx > -1 {
		return s[idx+1:]
	}
	return s
}

// StrTrim replace all string in trims to empty string
func StrTrim(s string, trims ...string) string {
	for _, item := range trims {
		s = strings.Replace(s, item, "", -1)
	}
	return s
}

// StrRegTrim replace all string then match regexp in trimRegs to empty string
func StrRegTrim(s string, trimRegs ...string) string {
	for _, item := range trimRegs {
		if aReg, err := regexp.Compile(item); err == nil {
			s = aReg.ReplaceAllString(s, "")
		}
	}
	return s
}

// --- Judge ---

// StrIsEmpty return the string argument is empty or not
func StrIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// StrMatchReg return s match reg or not
func StrMatchReg(s string, reg string) bool {
	return regexp.MustCompile(reg).MatchString(s)
}

// StrIsInt reutrn the string argument is int number or not
func StrIsInt(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

// StrIsFloat return the string argument is float or not
func StrIsFloat(s string) bool {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return false
	}
	return true
}

// StrIsNumber return the string argument is number or not
func StrIsNumber(s string) bool {
	return StrIsInt(s) || StrIsInt(s)
}

// StrIsHTTPURL return the string argument is http or https url or not
func StrIsHTTPURL(s string) bool {
	return StrMatchReg(s, cRegStrHttpURL)
}
