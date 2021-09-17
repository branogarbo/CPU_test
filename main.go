package main

import "sync"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0; i < 25; i++ {
		go RunTests(i)
	}

	wg.Wait()
}
