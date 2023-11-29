package day17

import (
	"reflect"
	"testing"
)

func TestChamber_trimTop(t *testing.T) {
	tests := []struct {
		name string
		c    *Chamber
		want []string
	}{
		{
			name: "basic",
			c:    &Chamber{settled: []string{"..####.", "....#..", "...###.", ".......", ".......", "......."}},
			want: []string{"..####.", "....#..", "...###."},
		},
		{
			name: "invert",
			c:    &Chamber{settled: []string{".......", ".......", ".......", "..####.", "....#..", "...###."}},
			want: []string{".......", ".......", ".......", "..####.", "....#..", "...###."},
		},
		{
			name: "both sides",
			c:    &Chamber{settled: []string{".......", ".......", ".......", "..####.", "....#..", "...###.", ".......", ".......", "......."}},
			want: []string{".......", ".......", ".......", "..####.", "....#..", "...###."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.trimTop()
			if got := tt.c.settled; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chamber.trimTop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChamber_trimBottom(t *testing.T) {
	tests := []struct {
		name string
		c    *Chamber
		want []string
	}{
		{
			name: "basic",
			c:    &Chamber{linesToStore: 5, settled: []string{"0000000", "1111111", "2222222", "3333333", "4444444", "5555555", "6666666", "7777777", "8888888", "9999999"}},
			want: []string{"5555555", "6666666", "7777777", "8888888", "9999999"},
		},
		{
			name: "no trim",
			c:    &Chamber{linesToStore: 15, settled: []string{"0000000", "1111111", "2222222", "3333333", "4444444", "5555555", "6666666", "7777777", "8888888", "9999999"}},
			want: []string{"0000000", "1111111", "2222222", "3333333", "4444444", "5555555", "6666666", "7777777", "8888888", "9999999"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.trimBottom()
			if got := tt.c.settled; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chamber.trimBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
