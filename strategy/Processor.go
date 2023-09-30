package main

type Processor struct {
	value     float32
	operation Operation
}

func initProcessor(o Operation, value float32) *Processor {
	return &Processor{
		value:     value,
		operation: o,
	}
}

func (p *Processor) setOperation(o Operation) {
	p.operation = o
}

func (p *Processor) execute(value float32) {
	p.operation.execute(p, value)
}
