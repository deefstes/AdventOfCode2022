package day20

import (
	"strconv"

	"github.com/deefstes/AdventOfCode2022/helpers"
	"github.com/google/uuid"
)

type Element struct {
	id    string
	value int64
}

type CircBuf struct {
	Size int

	buffer []Element
}

func (c *CircBuf) Add(i int64) {
	id := uuid.NewString()
	c.buffer = append(c.buffer, Element{id: id, value: i})
	c.Size = len(c.buffer)
}

func (c *CircBuf) IndexOfID(id string) int {
	for i := 0; i < len(c.buffer); i++ {
		if c.buffer[i].id == id {
			return i
		}
	}

	return -1
}

func (c *CircBuf) DeWrap(val int64) int {
	for helpers.Abs64(val) > int64(c.Size) {
		val = val%int64(c.Size) + val/int64(c.Size)
	}

	return int(val)
}

func (c *CircBuf) MoveElement(e Element) {
	increment := c.DeWrap(e.value)
	oldIndex := c.IndexOfID(e.id)
	newIndex := oldIndex + int(increment)

	// Modulo causes off-by-one error, too lazy to figure out why
	for newIndex >= c.Size {
		newIndex = newIndex - c.Size + 1
	}
	for newIndex < 0 {
		newIndex = newIndex + c.Size - 1
	}

	if oldIndex == newIndex {
		return
	}

	var newBuffer []Element

	if oldIndex > newIndex {
		newBuffer = append(newBuffer, c.buffer[:newIndex]...)
		newBuffer = append(newBuffer, e)
		newBuffer = append(newBuffer, c.buffer[newIndex:oldIndex]...)
		newBuffer = append(newBuffer, c.buffer[oldIndex+1:]...)
	} else {
		newBuffer = append(newBuffer, c.buffer[:oldIndex]...)
		newBuffer = append(newBuffer, c.buffer[oldIndex+1:newIndex+1]...)
		newBuffer = append(newBuffer, e)
		newBuffer = append(newBuffer, c.buffer[newIndex+1:]...)
	}
	c.buffer = newBuffer
}

func (c *CircBuf) MoveAllElements(iterations int) {
	var ee []Element
	ee = append(ee, c.buffer...)

	for i := 0; i < iterations; i++ {
		for ndx, e := range ee {
			_ = ndx
			c.MoveElement(e)
		}
	}
}

func (c *CircBuf) ApplyKey(key int64) {
	for i := 0; i < len(c.buffer); i++ {
		c.buffer[i].value = c.buffer[i].value * int64(key)
	}
}

func (c *CircBuf) CalculateCoordinates(key int64, iterations int) int64 {
	c.ApplyKey(key)
	c.MoveAllElements(iterations)

	// Find zero index
	zeroIndex := 0
	for i, e := range c.buffer {
		if e.value == 0 {
			zeroIndex = i
			break
		}
	}

	i1000 := (zeroIndex + 1000) % c.Size
	i2000 := (zeroIndex + 2000) % c.Size
	i3000 := (zeroIndex + 3000) % c.Size

	return c.buffer[i1000].value + c.buffer[i2000].value + c.buffer[i3000].value
}

func NewCircBuf(input []string) CircBuf {
	var circbuf CircBuf

	for _, line := range input {
		i, _ := strconv.Atoi(line)
		circbuf.Add(int64(i))
	}

	return circbuf
}
