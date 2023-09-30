package main

type Controller struct {
	observerList []Observer
}

func (c *Controller) update() {
	c.notifyAll()
}

func (c *Controller) register(observer Observer) {
	c.observerList = append(c.observerList, observer)
}

func (c *Controller) deregister(o Observer) {
	c.observerList = removeFromSlice(c.observerList, o)
}

func (c *Controller) notifyAll() {
	for _, observer := range c.observerList {
		observer.update("Updated status")
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
