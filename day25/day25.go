package day25

import (
	"math"
)

func SnafuToDec(snafu string) int {
	var dec int
	for i, d := range snafu {
		exp := len(snafu) - i - 1
		var digit int
		switch d {
		case '=':
			digit = -2
		case '-':
			digit = -1
		case '0':
			digit = 0
		case '1':
			digit = 1
		case '2':
			digit = 2
		}
		dec += digit * int(math.Pow(5, float64(exp)))
	}

	return dec
}

type bt []int8

func (b bt) String() string {
	if len(b) == 0 {
		return "0"
	}
	last := len(b) - 1
	r := make([]byte, len(b))
	for i, d := range b {
		r[last-i] = "=-012"[d+1]
		// r[last-i] = "-0+"[d+1]
	}
	return string(r)
}

func DecToSnafu_bak(i int) string {
	if i == 0 {
		return new(bt).String()
	}
	var b bt
	var btDigit func(int)
	btDigit = func(digit int) {
		m := int8(i % 5)
		i /= 5
		switch m {
		case 2:
			m = -1
			i++
		case -2:
			m = 1
			i--
		case 3: //?
			m = -2
			i += 2
		case -3: //?
			m = 2
			i -= 2
		}
		if i == 0 {
			b = make(bt, digit+1)
		} else {
			btDigit(digit + 1)
		}
		b[digit] = m
	}
	btDigit(0)
	return b.String()
}

func DecToSnafu(i int) (snafu string) {
	for i > 0 {
		snafu = string("=-012"[(i+2)%5]) + snafu
		i = (i + 2) / 5
	}
	return snafu
}
