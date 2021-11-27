package shared

type Intcode struct {
	memory  []int
	pointer int
}

func NewIntcode(memory []int) *Intcode {
	return &Intcode{
		memory:  memory,
		pointer: 0,
	}
}

func (i *Intcode) GetMemory() []int {
	return i.memory
}

func (i *Intcode) GetValue(index int) int {
	return i.memory[index]
}

func (i *Intcode) SetValue(index, value int) *Intcode {
	i.memory[index] = value
	return i
}

func (i *Intcode) MovePointer(delta int) *Intcode {
	i.pointer += delta
	return i
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

	sum := i.GetValue(a) + i.GetValue(b)
	i.SetValue(c, sum)
	i.MovePointer(4)

	return i
}

func (i *Intcode) Multiply() *Intcode {
	a := i.memory[i.pointer+1]
	b := i.memory[i.pointer+2]
	c := i.memory[i.pointer+3]

	product := i.GetValue(a) * i.GetValue(b)
	i.SetValue(c, product)
	i.MovePointer(4)

	return i
}
