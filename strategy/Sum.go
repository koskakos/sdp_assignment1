package main

type Sum struct {
}

func (s *Sum) execute(p *Processor, v float32) {
	p.value += v
}
