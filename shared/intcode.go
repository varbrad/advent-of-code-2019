package shared

type Intcode struct {
	memory         []int
	originalMemory []int
	pointer        int
}

func NewIntcode(memory []int) *Intcode {
	originalMemory := make([]int, len(memory))
	copy(originalMemory, memory)
	i := &Intcode{
		memory:         []int{},
		originalMemory: originalMemory,
		pointer:        0,
	}
	i.SetMemory(memory)
	return i
}

func (i *Intcode) Reset() *Intcode {
	i.pointer = 0
	i.SetMemory(i.originalMemory)
	return i
}

func (i *Intcode) GetMemory() []int {
	return i.memory
}

func (i *Intcode) SetMemory(memory []int) *Intcode {
	i.memory = make([]int, len(memory))
	copy(i.memory, memory)
	return i
}

func (i *Intcode) GetAddress(index int) int {
	return i.memory[index]
}

func (i *Intcode) SetAddress(index, value int) *Intcode {
	i.memory[index] = value
	return i
}

func (i *Intcode) MovePointer(delta int) *Intcode {
	i.pointer += delta
	return i
}

func (i *Intcode) SetPointer(value int) *Intcode {
	i.pointer = value
	return i
}

func (i *Intcode) GetPointer() int {
	return i.pointer
}

func (i *Intcode) Run() *Intcode {
	for {
		_, halted := i.Step()
		if halted {
			return i
		}
	}
}

func (i *Intcode) Step() (intcode *Intcode, halted bool) {
	instruction := i.memory[i.pointer]

	if instruction == 1 {
		return i.Add(), false
	} else if instruction == 2 {
		return i.Multiply(), false
	}

	return i, true
}

func (i *Intcode) Add() *Intcode {
	a := i.memory[i.pointer+1]
	b := i.memory[i.pointer+2]
	c := i.memory[i.pointer+3]

	sum := i.GetAddress(a) + i.GetAddress(b)
	i.SetAddress(c, sum)
	i.MovePointer(4)

	return i
}

func (i *Intcode) Multiply() *Intcode {
	a := i.memory[i.pointer+1]
	b := i.memory[i.pointer+2]
	c := i.memory[i.pointer+3]

	product := i.GetAddress(a) * i.GetAddress(b)
	i.SetAddress(c, product)
	i.MovePointer(4)

	return i
}
