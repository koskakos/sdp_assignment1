package main

func main() {
	controller := Controller{}

	firstObserver := &User{id: 1}
	secondObserver := &User{id: 2}

	controller.register(firstObserver)
	controller.register(secondObserver)

	controller.update()
}
