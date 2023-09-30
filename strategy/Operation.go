package main

type Operation interface {
	execute(p *Processor, v float32)
}
