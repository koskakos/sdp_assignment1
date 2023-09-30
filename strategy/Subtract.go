package main

type Subtract struct {
}

func (s *Subtract) execute(p *Processor, v float32) {
	p.value -= v
}
