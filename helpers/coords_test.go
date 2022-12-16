package helpers

import (
	"reflect"
	"sort"
	"testing"
)

func TestCoords_ManhattanNeighbourhood(t *testing.T) {
	tests := []struct {
		name string
		c    Coords
		dist int
		want []Coords
	}{
		{"3around0,0", NewCoords(0, 0), 3, []Coords{
			NewCoords(0, 0),

			NewCoords(0, 0).Right(1),
			NewCoords(0, 0).Right(1).Right(1),
			NewCoords(0, 0).Right(1).Right(1).Right(1),
			NewCoords(0, 0).Right(1).Down(1),
			NewCoords(0, 0).Right(1).Right(1).Down(1),
			NewCoords(0, 0).Right(1).Down(1).Down(1),

			NewCoords(0, 0).Up(1),
			NewCoords(0, 0).Up(1).Up(1),
			NewCoords(0, 0).Up(1).Up(1).Up(1),
			NewCoords(0, 0).Up(1).Right(1),
			NewCoords(0, 0).Up(1).Up(1).Right(1),
			NewCoords(0, 0).Up(1).Right(1).Right(1),

			NewCoords(0, 0).Left(1),
			NewCoords(0, 0).Left(1).Left(1),
			NewCoords(0, 0).Left(1).Left(1).Left(1),
			NewCoords(0, 0).Left(1).Up(1),
			NewCoords(0, 0).Left(1).Left(1).Up(1),
			NewCoords(0, 0).Left(1).Up(1).Up(1),

			NewCoords(0, 0).Down(1),
			NewCoords(0, 0).Down(1).Down(1),
			NewCoords(0, 0).Down(1).Down(1).Down(1),
			NewCoords(0, 0).Down(1).Left(1),
			NewCoords(0, 0).Down(1).Down(1).Left(1),
			NewCoords(0, 0).Down(1).Left(1).Left(1),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.ManhattanNeighbourhood(tt.dist)
			sort.Slice(got, func(i, j int) bool { return got[i].String() < got[j].String() })
			sort.Slice(tt.want, func(i, j int) bool { return tt.want[i].String() < tt.want[j].String() })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coords.ManhattanNeighbourhood() = %v, want %v", got, tt.want)
			}
		})
	}
}
