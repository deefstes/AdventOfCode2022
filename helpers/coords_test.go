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

func TestCoords_Rotate(t *testing.T) {
	type args struct {
		dir    int
		width  int
		height int
	}
	tests := []struct {
		name string
		c    Coords
		args args
		want Coords
	}{
		{
			name: "0,0 in 3x3 right",
			c:    NewCoords(0, 0),
			args: args{dir: 1, width: 3, height: 3},
			want: NewCoords(2, 0),
		},
		{
			name: "1,0 in 3x3 right",
			c:    NewCoords(1, 0),
			args: args{dir: 1, width: 3, height: 3},
			want: NewCoords(2, 1),
		},
		{
			name: "2,0 in 3x3 right",
			c:    NewCoords(2, 0),
			args: args{dir: 1, width: 3, height: 3},
			want: NewCoords(2, 2),
		},
		{
			name: "1,1 in 3x3 right",
			c:    NewCoords(1, 1),
			args: args{dir: 1, width: 3, height: 3},
			want: NewCoords(1, 1),
		},

		{
			name: "0,0 in 4x4 right",
			c:    NewCoords(0, 0),
			args: args{dir: 1, width: 4, height: 4},
			want: NewCoords(3, 0),
		},
		{
			name: "1,0 in 4x4 right",
			c:    NewCoords(1, 0),
			args: args{dir: 1, width: 4, height: 4},
			want: NewCoords(3, 1),
		},
		{
			name: "3,0 in 4x4 right",
			c:    NewCoords(3, 0),
			args: args{dir: 1, width: 4, height: 4},
			want: NewCoords(3, 3),
		},
		{
			name: "1,1 in 4x4 right",
			c:    NewCoords(1, 1),
			args: args{dir: 1, width: 4, height: 4},
			want: NewCoords(2, 1),
		},

		{
			name: "0,0 in 3x3 left",
			c:    NewCoords(0, 0),
			args: args{dir: -1, width: 3, height: 3},
			want: NewCoords(0, 2),
		},
		{
			name: "1,0 in 3x3 left",
			c:    NewCoords(1, 0),
			args: args{dir: -1, width: 3, height: 3},
			want: NewCoords(0, 1),
		},
		{
			name: "2,0 in 3x3 left",
			c:    NewCoords(2, 0),
			args: args{dir: -1, width: 3, height: 3},
			want: NewCoords(0, 0),
		},
		{
			name: "1,1 in 3x3 left",
			c:    NewCoords(1, 1),
			args: args{dir: -1, width: 3, height: 3},
			want: NewCoords(1, 1),
		},

		{
			name: "0,0 in 4x4 left",
			c:    NewCoords(0, 0),
			args: args{dir: -1, width: 4, height: 4},
			want: NewCoords(0, 3),
		},
		{
			name: "1,0 in 4x4 left",
			c:    NewCoords(1, 0),
			args: args{dir: -1, width: 4, height: 4},
			want: NewCoords(0, 2),
		},
		{
			name: "3,0 in 4x4 left",
			c:    NewCoords(3, 0),
			args: args{dir: -1, width: 4, height: 4},
			want: NewCoords(0, 0),
		},
		{
			name: "1,1 in 4x4 left",
			c:    NewCoords(1, 1),
			args: args{dir: -1, width: 4, height: 4},
			want: NewCoords(1, 2),
		},

		{
			name: "0,0 in 4x4 rightx2",
			c:    NewCoords(0, 0),
			args: args{dir: 2, width: 4, height: 4},
			want: NewCoords(3, 3),
		},
		{
			name: "0,0 in 4x4 leftx2",
			c:    NewCoords(0, 0),
			args: args{dir: -2, width: 4, height: 4},
			want: NewCoords(3, 3),
		},
		{
			name: "0,0 in 4x4 rightx4",
			c:    NewCoords(0, 0),
			args: args{dir: 4, width: 4, height: 4},
			want: NewCoords(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Rotate(tt.args.dir, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coords.Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
