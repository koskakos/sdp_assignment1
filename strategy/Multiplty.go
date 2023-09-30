package main

type Multiply struct {
}

func (s *Multiply) execute(p *Processor, v float32) {
	p.value *= v
}
