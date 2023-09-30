package main

import "fmt"

func main() {
	sum := &Sum{}
	processor := initProcessor(sum, 0)

	processor.execute(32)
	fmt.Println(processor.value)
	processor.execute(53)
	fmt.Println(processor.value)

	subtract := &Subtract{}
	processor.setOperation(subtract)

	processor.execute(52)
	fmt.Println(processor.value)

	multiply := &Multiply{}
	processor.setOperation(multiply)

	processor.execute(2)
	fmt.Println(processor.value)

}
