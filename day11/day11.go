package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Troop struct {
	monkeys    []Monkey
	calmingDiv int64
}

func NewTroop() Troop {
	return Troop{calmingDiv: 1}
}

func (t *Troop) AddMonkey(m Monkey) {
	t.monkeys = append(t.monkeys, m)
	t.calmingDiv = t.calmingDiv * m.testDiv
}

func (t *Troop) String() string {
	var troop string
	for _, m := range t.monkeys {
		troop = fmt.Sprintf("%s\n%s", troop, m.String())
	}

	return strings.TrimLeft(troop, "\n")
}

func (t *Troop) SortInspections() []Monkey {
	sort.Slice(t.monkeys, func(i, j int) bool { return t.monkeys[i].Inspections > t.monkeys[j].Inspections })
	return t.monkeys
}

func (t *Troop) PlayRound(withRelief, withCalming bool) {
	for i := range t.monkeys {
		t.monkeys[i].PlayRound(t, withRelief, withCalming)
	}
}

type Monkey struct {
	Name        string
	items       []int64
	operation   string
	amount      string
	testDiv     int64
	monkeyTrue  int
	monkeyFalse int
	Inspections int
}

func NewMonkey(lines []string) Monkey {
	var ii []int64
	for _, v := range strings.Split(strings.TrimPrefix(lines[1], "  Starting items: "), ", ") {
		i, _ := strconv.ParseInt(v, 10, 64)
		ii = append(ii, i)
	}
	op := strings.Split(strings.TrimPrefix(lines[2], "  Operation: new = old "), " ")
	div, _ := strconv.ParseInt(strings.TrimPrefix(lines[3], "  Test: divisible by "), 10, 64)
	mt, _ := strconv.Atoi(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
	mf, _ := strconv.Atoi(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

	m := Monkey{
		Name:        strings.TrimRight(lines[0], ":"),
		items:       ii,
		operation:   op[0],
		amount:      op[1],
		testDiv:     div,
		monkeyTrue:  mt,
		monkeyFalse: mf,
	}

	return m
}

func (m *Monkey) AddItem(i int64, withRelief bool, withCalming bool, troop *Troop) {
	if withCalming {
		i = i % troop.calmingDiv
	}

	m.items = append(m.items, i)
}

func (m *Monkey) String() string {
	var retval string
	for _, v := range m.items {
		retval = fmt.Sprintf("%s, %d", retval, v)
	}

	return fmt.Sprintf("%s: %s", m.Name, strings.TrimPrefix(retval, ", "))
}

func (m *Monkey) PlayRound(t *Troop, withRelief, withCalming bool) {
	for _, v := range m.items {
		m.Inspections = m.Inspections + 1

		var amnt int64
		if m.amount == "old" {
			amnt = v
		} else {
			amnt, _ = strconv.ParseInt(m.amount, 10, 64)
		}
		// Apply inspection operation
		switch m.operation {
		case "*":
			v = v * amnt
		case "+":
			v = v + amnt
		}

		// Apply bored operation
		if withRelief {
			v = v / 3
		}

		if v%(m.testDiv) == 0 {
			t.monkeys[m.monkeyTrue].AddItem(v, withRelief, withCalming, t)
		} else {
			t.monkeys[m.monkeyFalse].AddItem(v, withRelief, withCalming, t)
		}
	}
	m.items = make([]int64, 0)
}
