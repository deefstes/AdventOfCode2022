package day21

import (
	"testing"
)

func TestTroop_SolveAdvanced(t *testing.T) {
	tests := []struct {
		name    string
		tr      *Troop
		mname   string
		seedVal float64
		want    float64
	}{
		{
			name: "test01",
			tr: func() *Troop {
				t := NewTroop([]string{
					"root: aaaa + humn",
					"aaaa: 5",
					"humn: 1000",
				})
				return &t
			}(),
			mname:   "root",
			seedVal: 0,
			want:    5,
		},
		{
			name: "test02",
			tr: func() *Troop {
				t := NewTroop([]string{
					"root: aaaa + bbbb",
					"aaaa: cccc + humn",
					"cccc: 2",
					"humn: 1000",
					"bbbb: dddd - eeee",
					"dddd: 4",
					"eeee: 1",
				})
				return &t
			}(),
			mname:   "root",
			seedVal: 0,
			want:    1,
		},
		{
			name: "test03",
			tr: func() *Troop {
				t := NewTroop([]string{
					"root: pppw + sjmn",
					"dbpl: 5",
					"cczh: sllz + lgvd",
					"zczc: 2",
					"ptdq: humn - dvpt",
					"dvpt: 3",
					"lfqf: 4",
					"humn: 5",
					"ljgn: 2",
					"sjmn: drzm * dbpl",
					"sllz: 4",
					"pppw: cczh / lfqf",
					"lgvd: ljgn * ptdq",
					"drzm: hmdt - zczc",
					"hmdt: 32",
				})
				return &t
			}(),
			mname:   "root",
			seedVal: 0,
			want:    301,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.SolveAdvanced(false); got != tt.want {
				t.Errorf("Troop.SolveForHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseOperator(t *testing.T) {
	type args struct {
		operator      string
		seedVal       float64
		otherTree     float64
		unknownMonkey int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "left+",
			args: args{
				operator:      "+",
				seedVal:       10,
				otherTree:     4,
				unknownMonkey: 1,
			},
			want: 6,
		},
		{
			name: "right+",
			args: args{
				operator:      "+",
				seedVal:       600,
				otherTree:     4,
				unknownMonkey: 2,
			},
			want: 596,
		},
		{
			name: "left-",
			args: args{
				operator:      "-",
				seedVal:       5,
				otherTree:     6,
				unknownMonkey: 1,
			},
			want: 11,
		},
		{
			name: "right-",
			args: args{
				operator:      "-",
				seedVal:       1,
				otherTree:     4,
				unknownMonkey: 2,
			},
			want: 3,
		},
		{
			name: "left*",
			args: args{
				operator:      "*",
				seedVal:       15,
				otherTree:     5,
				unknownMonkey: 1,
			},
			want: 3,
		},
		{
			name: "right*",
			args: args{
				operator:      "*",
				seedVal:       5,
				otherTree:     2,
				unknownMonkey: 2,
			},
			want: 2.5,
		},
		{
			name: "left/",
			args: args{
				operator:      "/",
				seedVal:       3,
				otherTree:     5,
				unknownMonkey: 1,
			},
			want: 15,
		},
		{
			name: "right/",
			args: args{
				operator:      "/",
				seedVal:       5,
				otherTree:     20,
				unknownMonkey: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseOperator(tt.args.operator, tt.args.seedVal, tt.args.otherTree, tt.args.unknownMonkey); got != tt.want {
				t.Errorf("ReverseOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
