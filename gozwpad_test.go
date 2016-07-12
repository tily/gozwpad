package gozwpad

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_Pad_PadsWithZeroWidthChars(t *testing.T) {
	var ss []string
	for _, c := range ZERO_WIDTH_CHARS {
		ss = append(ss, string(c))
	}
	r, _ := regexp.Compile(fmt.Sprintf("^[%s]{139}a$", strings.Join(ss, "")))
	if !r.MatchString(Pad("a", 140)) {
		t.Fatal("String is not padded correctly")
	}
}

func Test_Pad_PadsWithRandomChars(t *testing.T) {
	a := Pad("a", 140)
	b := Pad("a", 140)
	if a == b {
		t.Fatal("Padding is not random")
	}
}
