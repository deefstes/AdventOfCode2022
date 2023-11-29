package day20

import (
	"fmt"
	"testing"
)

func CreateTestElement(val int64) Element {
	return Element{id: fmt.Sprintf("%d", val), value: val}
}

func CreateTestElementSlice(vals []int64) []Element {
	var retval []Element
	for _, val := range vals {
		retval = append(retval, CreateTestElement(val))
	}

	return retval
}

func CompareTestElementSlices(s1, s2 []Element) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i].value != s2[i].value {
			return false
		}
	}

	return true
}

func TestCircBuf_MoveElement(t *testing.T) {
	tests := []struct {
		name       string
		c          *CircBuf
		e          Element
		key        int64
		iterations int
		wantBuf    []Element
	}{
		{
			name: "no movement",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, 0, 50, 60, 70, 80})},
			e:    CreateTestElement(0),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, 0, 50, 60, 70, 80}),
		},
		{
			name: "right (full wrap)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, 7, 50, 60, 70, 80})},
			e:    CreateTestElement(7),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, 7, 50, 60, 70, 80}),
		},
		{
			name: "right (no wrap)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, 2, 50, 60, 70, 80})},
			e:    CreateTestElement(2),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, 50, 60, 2, 70, 80}),
		},
		{
			name: "right (with wrap)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, 6, 50, 60, 70, 80})},
			e:    CreateTestElement(6),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 6, 30, 50, 60, 70, 80}),
		},
		{
			name: "right (overtake)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, 9, 50, 60, 70, 80})},
			e:    CreateTestElement(9),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, 50, 60, 9, 70, 80}),
		},
		{
			name: "left (full wrap)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, -7, 50, 60, 70, 80})},
			e:    CreateTestElement(-7),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, -7, 50, 60, 70, 80}),
		},
		{
			name: "left (no wrap)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, -2, 50, 60, 70, 80})},
			e:    CreateTestElement(-2),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, -2, 20, 30, 50, 60, 70, 80}),
		},
		{
			name: "left (with wrap)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, -6, 50, 60, 70, 80})},
			e:    CreateTestElement(-6),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, 50, -6, 60, 70, 80}),
		},
		{
			name: "left (overtake)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, -9, 50, 60, 70, 80})},
			e:    CreateTestElement(-9),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, -9, 20, 30, 50, 60, 70, 80}),
		},
		{
			name: "right (to edge)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, 4, 50, 60, 70, 80})},
			e:    CreateTestElement(4),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{10, 20, 30, 50, 60, 70, 80, 4}),
		},
		{
			name: "left (to edge)",
			c:    &CircBuf{Size: 8, buffer: CreateTestElementSlice([]int64{10, 20, 30, -3, 50, 60, 70, 80})},
			e:    CreateTestElement(-3),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{-3, 10, 20, 30, 50, 60, 70, 80}),
		},
		{
			name: "example 1.1",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			e:    CreateTestElement(1),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{2, 1, -3, 3, -2, 0, 4}),
		},
		{
			name: "example 1.2",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{2, 1, -3, 3, -2, 0, 4})},
			e:    CreateTestElement(2),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{1, -3, 2, 3, -2, 0, 4}),
		},
		{
			name: "example 1.3",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, -3, 2, 3, -2, 0, 4})},
			e:    CreateTestElement(-3),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{1, 2, 3, -2, -3, 0, 4}),
		},
		{
			name: "example 1.4",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, 3, -2, -3, 0, 4})},
			e:    CreateTestElement(3),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{1, 2, -2, -3, 0, 3, 4}),
		},
		{
			name: "example 1.5",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -2, -3, 0, 3, 4})},
			e:    CreateTestElement(-2),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{-2, 1, 2, -3, 0, 3, 4}),
		},
		{
			name: "example 1.6",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 0, 3, 4, -2})},
			e:    CreateTestElement(0),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{1, 2, -3, 0, 3, 4, -2}),
		},
		{
			name: "example 1.7",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 0, 3, 4, -2})},
			e:    CreateTestElement(4),
			key:  1, iterations: 1,
			wantBuf: CreateTestElementSlice([]int64{1, 2, -3, 4, 0, 3, -2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.MoveElement(tt.e)
			if !CompareTestElementSlices(tt.c.buffer, tt.wantBuf) {
				t.Errorf("buffers mismatch:\n got -> %v\nwant -> %v", tt.c.buffer, tt.wantBuf)
			}
		})
	}
}

func TestCircBuf_ApplyKey(t *testing.T) {
	tests := []struct {
		name    string
		c       *CircBuf
		key     int64
		wantBuf []Element
	}{
		{
			name:    "example 2.1",
			c:       &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			key:     811589153,
			wantBuf: CreateTestElementSlice([]int64{811589153, 1623178306, -2434767459, 2434767459, -1623178306, 0, 3246356612}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.ApplyKey(tt.key)
			if !CompareTestElementSlices(tt.c.buffer, tt.wantBuf) {
				t.Errorf("buffers mismatch:\n got -> %v\nwant -> %v", tt.c.buffer, tt.wantBuf)
			}
		})
	}
}

func TestCircBuf_DeWrap(t *testing.T) {
	tests := []struct {
		name string
		c    *CircBuf
		val  int64
		want int
	}{
		{
			name: "right wrap",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			val:  8,
			want: 2,
		},
		{
			name: "right double wrap",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			val:  16,
			want: 4,
		},
		{
			name: "right crazy wrap",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			val:  12345,
			want: 3,
		},
		{
			name: "left wrap",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			val:  -8,
			want: -2,
		},
		{
			name: "left double wrap",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			val:  -16,
			want: -4,
		},
		{
			name: "left crazy wrap",
			c:    &CircBuf{Size: 7, buffer: CreateTestElementSlice([]int64{1, 2, -3, 3, -2, 0, 4})},
			val:  -12345,
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.DeWrap(tt.val); got != tt.want {
				t.Errorf("CircBuf.DeWrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
