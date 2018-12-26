package main

import (
	"fmt"
	"sync"
	"time"
)

func usarCerrojo(c *sync.Mutex, i int, wg *sync.WaitGroup) {
	c.Lock()
	defer c.Unlock()
	time.Sleep(time.Second)
	fmt.Println("Padawan", i, "usando el cerrojo")
	wg.Done()
}

func main() {
	cerrojo := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go usarCerrojo(cerrojo, 1, wg)
	go usarCerrojo(cerrojo, 2, wg)
	go usarCerrojo(cerrojo, 3, wg)
	go usarCerrojo(cerrojo, 4, wg)
	wg.Wait()
}
