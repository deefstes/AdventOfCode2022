package day21

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	monkey1  string
	monkey2  string
	operator string
}

func (o *Operation) Noop() bool {
	return o.operator == ""
}

func NewOperation(input string) Operation {
	w := strings.Split(input, " ")
	return Operation{
		monkey1:  w[0],
		monkey2:  w[2],
		operator: w[1],
	}
}

type Monkey struct {
	name      string
	number    float64
	operation Operation
	solved    bool
}

func NewMonkey(input string) Monkey {
	var monkey Monkey

	monkey.name = input[0:4]
	num, err := strconv.Atoi(input[6:])
	if err != nil {
		monkey.operation = NewOperation(input[6:])
	} else {
		monkey.number, monkey.solved = float64(num), true
	}

	return monkey
}

type Troop struct {
	monkeys map[string]*Monkey
}

func (t *Troop) MonkeyFormula(m *Monkey) string {
	if m.solved {
		return fmt.Sprintf("%.f", m.number)
	}

	if m.operation.Noop() {
		return m.name
	}

	return fmt.Sprintf("(%s%s%s)", t.MonkeyFormula(t.monkeys[m.operation.monkey1]), m.operation.operator, t.MonkeyFormula(t.monkeys[m.operation.monkey2]))
}

func (t *Troop) SolveMonkey(m *Monkey) (float64, bool) {
	if m.solved {
		return m.number, true
	}

	if m.operation.Noop() {
		return m.number, false
	}

	monkey1 := t.monkeys[m.operation.monkey1]
	monkey2 := t.monkeys[m.operation.monkey2]
	operator := m.operation.operator

	num1, ok := t.SolveMonkey(monkey1)
	if !ok {
		return num1, false
	}
	num2, ok := t.SolveMonkey(monkey2)
	if !ok {
		return num2, false
	}

	switch operator {
	case "+":
		m.number, m.solved = num1+num2, true
	case "-":
		m.number, m.solved = num1-num2, true
	case "*":
		m.number, m.solved = num1*num2, true
	case "/":
		m.number, m.solved = num1/num2, true
	}

	return m.number, m.solved
}

func (t *Troop) SolveBasic() float64 {
	rootMonkey := t.monkeys["root"]

	v, ok := t.SolveMonkey(rootMonkey)
	if ok {
		return v
	}

	return 0
}

func (t *Troop) SolveAdvanced(showSteps bool) float64 {
	m := *t.monkeys["root"]
	m.operation.operator = "-"
	human := t.monkeys["humn"]
	human.number = 0
	human.solved = false
	var seedVal float64

	for {
		if m.name == "humn" {
			return seedVal
		}

		m1 := t.monkeys[m.operation.monkey1]
		m2 := t.monkeys[m.operation.monkey2]
		v1, solved1 := t.SolveMonkey(m1)
		v2, _ := t.SolveMonkey(m2)

		if solved1 {
			if showSteps {
				fmt.Printf("%.f%s%s=%.f\n\n", v1, m.operation.operator, t.MonkeyFormula(m2), seedVal)
			}
			seedVal = ReverseOperator(m.operation.operator, seedVal, v1, 2)
			m = *m2
		} else {
			if showSteps {
				fmt.Printf("%s%s%.f=%.f\n\n", t.MonkeyFormula(m1), m.operation.operator, v2, seedVal)
			}
			seedVal = ReverseOperator(m.operation.operator, seedVal, v2, 1)
			m = *m1
		}
	}
}

func NewTroop(input []string) Troop {
	troop := Troop{make(map[string]*Monkey)}

	for _, line := range input {
		monkey := NewMonkey(line)
		troop.monkeys[monkey.name] = &monkey
	}

	return troop
}

func ReverseOperator(operator string, seedVal, otherTree float64, unknownMonkey int) float64 {
	switch operator {
	case "+":
		return seedVal - otherTree
	case "-":
		if unknownMonkey == 1 {
			return seedVal + otherTree
		} else {
			return otherTree - seedVal
		}
	case "*":
		return seedVal / otherTree
	case "/":
		if unknownMonkey == 1 {
			return seedVal * otherTree
		} else {
			return otherTree / seedVal
		}
	}

	return 0
}
