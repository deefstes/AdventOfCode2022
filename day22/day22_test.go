package day22

import (
	"reflect"
	"testing"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

func Test_newPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Path
	}{
		{name: "test01", input: "10R5L5R10L4R5L5", want: Path{steps: []any{10, "R", 5, "L", 5, "R", 10, "L", 4, "R", 5, "L", 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPath(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_Left(t *testing.T) {
	tests := []struct {
		name   string
		d      Direction
		number int
		want   Direction
	}{
		{name: "North", d: north, number: 1, want: west},
		{name: "South", d: south, number: 1, want: east},
		{name: "East", d: east, number: 1, want: north},
		{name: "West", d: west, number: 1, want: south},

		{name: "North", d: north, number: 2, want: south},
		{name: "South", d: south, number: 2, want: north},
		{name: "East", d: east, number: 2, want: west},
		{name: "West", d: west, number: 2, want: east},

		{name: "North", d: north, number: 4, want: north},
		{name: "South", d: south, number: 4, want: south},
		{name: "East", d: east, number: 4, want: east},
		{name: "West", d: west, number: 4, want: west},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Left(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Direction.Left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_Right(t *testing.T) {
	tests := []struct {
		name   string
		d      Direction
		number int
		want   Direction
	}{
		{name: "North", d: north, number: 1, want: east},
		{name: "South", d: south, number: 1, want: west},
		{name: "East", d: east, number: 1, want: south},
		{name: "West", d: west, number: 1, want: north},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Right(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Direction.Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Move(t *testing.T) {
	testFlatBoard := NewBoard([]string{
		"        ...#",
		"        .#..",
		"        #...",
		"        ....",
		"...#.......#",
		"........#...",
		"..#....#....",
		"..........#.",
		"        ...#....",
		"        .....#..",
		"        .#......",
		"        ......#.",
		"",
		"10R5L5R10L4R5L5",
	}, false)

	testCubeBoard := NewBoard([]string{
		"        ...#",
		"        .#..",
		"        #...",
		"        ....",
		"...#.......#",
		"........#...",
		"..#....#....",
		"..........#.",
		"        ...#....",
		"        .....#..",
		"        .#......",
		"        ......#.",
		"",
		"10R5L5R10L4R5L5",
	}, true)

	type args struct {
		cell helpers.Coords
		dir  Direction
		dist int
	}
	tests := []struct {
		name       string
		b          *Board
		args       args
		wantCoords helpers.Coords
		wantNewDir Direction
	}{
		{
			name: "test_flat_01", b: &testFlatBoard,
			args:       args{cell: helpers.NewCoords(11, 6), dir: east, dist: 1},
			wantCoords: helpers.NewCoords(0, 6),
			wantNewDir: east,
		},
		{
			name: "test_flat_02", b: &testFlatBoard,
			args:       args{cell: helpers.NewCoords(5, 7), dir: south, dist: 1},
			wantCoords: helpers.NewCoords(5, 4),
			wantNewDir: south,
		},
		{
			name: "test_flat_03", b: &testFlatBoard,
			args:       args{cell: helpers.NewCoords(5, 7), dir: south, dist: 10},
			wantCoords: helpers.NewCoords(5, 5),
			wantNewDir: south,
		},
		{
			name: "test_flat_04", b: &testFlatBoard,
			args:       args{cell: helpers.NewCoords(5, 5), dir: west, dist: 12},
			wantCoords: helpers.NewCoords(9, 5),
			wantNewDir: west,
		},
		{
			name: "test_cube_01", b: &testCubeBoard,
			args:       args{cell: helpers.NewCoords(11, 6), dir: east, dist: 1},
			wantCoords: helpers.NewCoords(13, 8),
			wantNewDir: south,
		},
		{
			name: "test_cube_02", b: &testCubeBoard,
			args:       args{cell: helpers.NewCoords(10, 11), dir: south, dist: 1},
			wantCoords: helpers.NewCoords(1, 7),
			wantNewDir: north,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCoords, gotNewDir := tt.b.Move(tt.args.cell, tt.args.dir, tt.args.dist)
			if !reflect.DeepEqual(gotCoords, tt.wantCoords) {
				t.Errorf("Board.Move() gotCoords = %v, want %v", gotCoords, tt.wantCoords)
			}
			if !reflect.DeepEqual(gotNewDir, tt.wantNewDir) {
				t.Errorf("Board.Move() gotNewDir = %v, want %v", gotNewDir, tt.wantNewDir)
			}
		})
	}
}

func TestBoard_Advance(t *testing.T) {
	testCubeBoard := NewBoard([]string{
		"        ...#",
		"        .#..",
		"        #...",
		"        ....",
		"...#.......#",
		"........#...",
		"..#....#....",
		"..........#.",
		"        ...#....",
		"        .....#..",
		"        .#......",
		"        ......#.",
		"",
		"10R5L5R10L4R5L5",
	}, true)

	type args struct {
		cell helpers.Coords
		dir  Direction
	}
	tests := []struct {
		name        string
		b           *Board
		args        args
		wantNewCell helpers.Coords
		wantNewDir  Direction
		wantCanMove bool
	}{
		{
			name: "test01",
			b:    &testCubeBoard,
			args: args{
				cell: helpers.NewCoords(10, 3),
				dir:  south,
			},
			wantNewCell: helpers.NewCoords(10, 4),
			wantNewDir:  south,
			wantCanMove: true,
		},
		{
			name: "test02",
			b:    &testCubeBoard,
			args: args{
				cell: helpers.NewCoords(11, 5),
				dir:  east,
			},
			wantNewCell: helpers.NewCoords(14, 8),
			wantNewDir:  south,
			wantCanMove: true,
		},
		{
			name: "test04",
			b:    &testCubeBoard,
			args: args{
				cell: helpers.NewCoords(3, 5),
				dir:  east,
			},
			wantNewCell: helpers.NewCoords(4, 5),
			wantNewDir:  east,
			wantCanMove: true,
		},
		{
			name: "test05",
			b:    &testCubeBoard,
			args: args{
				cell: helpers.NewCoords(6, 4),
				dir:  north,
			},
			wantNewCell: helpers.NewCoords(6, 4),
			wantNewDir:  north,
			wantCanMove: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewCell, gotNewDir, gotCanMove := tt.b.Advance(tt.args.cell, tt.args.dir)
			if !reflect.DeepEqual(gotNewCell, tt.wantNewCell) {
				t.Errorf("Board.Advance() gotNewCell = %v, want %v", gotNewCell, tt.wantNewCell)
			}
			if !reflect.DeepEqual(gotNewDir, tt.wantNewDir) {
				t.Errorf("Board.Advance() gotNewDir = %v, want %v", gotNewDir, tt.wantNewDir)
			}
			if gotCanMove != tt.wantCanMove {
				t.Errorf("Board.Advance() gotCanMove = %v, want %v", gotCanMove, tt.wantCanMove)
			}
		})
	}
}
