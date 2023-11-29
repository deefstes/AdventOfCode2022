package day25

import (
	"testing"
)

func TestSnafuToDec(t *testing.T) {
	tests := []struct {
		name  string
		snafu string
		want  int
	}{
		{name: "test01", snafu: "2=-01", want: 976},
		{name: "test02", snafu: "1=-0-2", want: 1747},
		{name: "test03", snafu: "12111", want: 906},
		{name: "test04", snafu: "2=0=", want: 198},
		{name: "test05", snafu: "21", want: 11},
		{name: "test06", snafu: "2=01", want: 201},
		{name: "test07", snafu: "111", want: 31},
		{name: "test08", snafu: "20012", want: 1257},
		{name: "test09", snafu: "112", want: 32},
		{name: "test10", snafu: "1=-1=", want: 353},
		{name: "test11", snafu: "1-12", want: 107},
		{name: "test12", snafu: "12", want: 7},
		{name: "test13", snafu: "1=", want: 3},
		{name: "test14", snafu: "122", want: 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnafuToDec(tt.snafu); got != tt.want {
				t.Errorf("SnafuToDec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecToSnafu(t *testing.T) {
	tests := []struct {
		name string
		dec  int
		want string
	}{
		{name: "test02", dec: 1747, want: "1=-0-2"},
		{name: "test03", dec: 906, want: "12111"},
		{name: "test04", dec: 198, want: "2=0="},
		{name: "test05", dec: 11, want: "21"},
		{name: "test06", dec: 201, want: "2=01"},
		{name: "test07", dec: 31, want: "111"},
		{name: "test08", dec: 1257, want: "20012"},
		{name: "test09", dec: 32, want: "112"},
		{name: "test10", dec: 353, want: "1=-1="},
		{name: "test11", dec: 107, want: "1-12"},
		{name: "test12", dec: 7, want: "12"},
		{name: "test13", dec: 3, want: "1="},
		{name: "test14", dec: 37, want: "122"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecToSnafu(tt.dec); got != tt.want {
				t.Errorf("DecToSnafu() = %v, want %v", got, tt.want)
			}
		})
	}
}
