package day19

import "fmt"

const (
	TotalTime     = 24
	BotTypes      = 4
	ResourceTypes = 4
	OreIndex      = 0
	ClayIndex     = 1
	ObsidianIndex = 2
	GeodeIndex    = 3
)

type State struct {
	M             map[Key]Memo
	blueprint     *Blueprint
	maxProduction Supplies
	maxSeen       int
}

func NewState(blueprint *Blueprint) *State {
	s := &State{
		M:         make(map[Key]Memo),
		blueprint: blueprint,
	}
	for i := 0; i < ResourceTypes; i++ {
		for j := 0; j < BotTypes; j++ {
			s.maxProduction[i] = max(s.maxProduction[i], blueprint.Robots[j].Cost[i])
		}
	}
	s.maxProduction[optimizingIndex] = 127
	return s
}

var optimizingIndex = GeodeIndex

func (s *State) MaxGeodes(key *Key) (maxGeodes int) {
	defer func() {
		if maxGeodes > s.maxSeen {
			s.maxSeen = maxGeodes
		}
	}()

	timeRemaining := int(TotalTime - key.time)
	finalGeodesWithoutPurchasing := int(key.supplies[optimizingIndex]) + timeRemaining*int(key.botCounts[optimizingIndex])
	if timeRemaining < 2 {
		return finalGeodesWithoutPurchasing
	}

	finalGeodesBuyingGeodeBotEveryTurn := int(finalGeodesWithoutPurchasing) + timeRemaining*(timeRemaining-1)/2
	if finalGeodesBuyingGeodeBotEveryTurn < s.maxSeen {
		// Kill this branch that will never be best
		return 0
	}

	m, ok := s.M[*key]
	if ok {
		return m.MaxGeodes
	}

	var newSupplies Supplies
	for i := 0; i < BotTypes; i++ {
		newSupplies.add(mul(&s.blueprint.Robots[i].Output, key.botCounts[i]))
	}

	options := []Key{}
	for nextPurchase := 0; nextPurchase < BotTypes; nextPurchase++ {
		if key.botCounts[nextPurchase] >= s.maxProduction[nextPurchase] {
			continue
		}
		opt := *key
		opt.time++
		for opt.time < TotalTime {
			opt.supplies.sub(&s.blueprint.Robots[nextPurchase].Cost)
			if opt.supplies.valid() {
				opt.botCounts[nextPurchase]++
				break
			}
			opt.supplies.add(&s.blueprint.Robots[nextPurchase].Cost)
			opt.supplies.add(&newSupplies)
			opt.time++
		}
		if opt.time < TotalTime {
			options = append(options, opt)
		}
	}

	if len(options) == 0 {
		return int(key.supplies[optimizingIndex] + key.botCounts[optimizingIndex]*(TotalTime-key.time))
	}

	for i := range options {
		options[i].supplies.add(&newSupplies)
	}

	var NextKey Key
	for i := range options {
		if m := s.MaxGeodes(&options[i]); m >= maxGeodes {
			maxGeodes = m
			NextKey = options[i]
		}
	}

	s.M[*key] = Memo{MaxGeodes: maxGeodes, NextKey: NextKey}
	return maxGeodes
}

func max[T int8 | int](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type Memo struct {
	MaxGeodes int
	NextKey   Key
}

type Key struct {
	supplies, botCounts Supplies
	time                int8
}

func (k Key) String() string {
	return fmt.Sprintf("{supplies: %v bots: %v t: %d}", k.supplies, k.botCounts, k.time)
}

func NewKey() *Key {
	s := &Key{}
	s.botCounts[OreIndex] = 1
	return s
}

type Blueprint struct {
	Id     int
	Robots [BotTypes]Robot
}

type Supplies [ResourceTypes]int8

func (s *Supplies) valid() bool {
	for _, c := range *s {
		if c < 0 {
			return false
		}
	}
	return true
}

func (s *Supplies) sub(o *Supplies) {
	for i := range s {
		s[i] -= o[i]
	}
}

func (s *Supplies) add(o *Supplies) {
	for i := range s {
		s[i] += o[i]
	}
}

func mul(s *Supplies, n int8) *Supplies {
	var r Supplies
	for i := range s {
		r[i] = s[i] * n
	}
	return &r
}

type Robot struct {
	Cost, Output Supplies
}
