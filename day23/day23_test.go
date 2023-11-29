package day23

import (
	"reflect"
	"testing"
)

func TestElfMob_RotateMovePriority(t *testing.T) {
	tests := []struct {
		name string
		e    *ElfMob
		want []string
	}{
		{
			name: "test01",
			e:    &ElfMob{movePriority: []string{"N", "S", "E", "W"}},
			want: []string{"S", "E", "W", "N"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.RotateMovePriority()
			got := tt.e.movePriority
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElfMob.RotateMovePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
