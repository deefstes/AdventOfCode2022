package day06

func FindNonRepeating(input []byte, len int) int {
	cb := MakeCircularBuffer(len)
	for i, b := range input {
		cb.Push(b)
		if cb.IsUnique() {
			return i + 1
		}
	}

	return 0
}

type CircularBuffer struct {
	size   int
	index  int
	buffer []byte
}

func MakeCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{
		size:   size,
		buffer: make([]byte, size),
	}
}

func (c *CircularBuffer) Push(b byte) {
	c.buffer[c.index] = b
	c.index = (c.index + 1) % c.size
}

func (c *CircularBuffer) IsUnique() bool {
	m := make(map[byte]bool)
	for _, b := range c.buffer {
		if b == 0 {
			return false
		}
		if m[b] {
			return false
		}
		m[b] = true
	}

	return true
}
