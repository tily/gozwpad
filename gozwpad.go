package gozwpad

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

var ZERO_WIDTH_CHARS = []int{
	0x200b,
	0x200c,
	0x200d,
	0x200e,
	0x200f,
	0x202a,
	0x202b,
	0x202c,
	0x202d,
	0x202e,
	0x2060,
	0x2061,
	0x2062,
	0x2063,
	0x2064,
	0x2066,
	0x2067,
	0x2068,
	0x2069,
	0x206a,
	0x206b,
	0x206c,
	0x206d,
	0x206e,
	0x206f,
}

func sample(list []int) int {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

func Pad(str string, size int) (ret string) {
	orgSize := utf8.RuneCountInString(str)
	for i := 0; i < (size - orgSize); i++ {
		ret = ret + string(sample(ZERO_WIDTH_CHARS))
	}
	ret = ret + str
	return ret
}
