package CPU_test

import (
	"sync"
)

type Config struct {
	Concurrency int
	Operation   func(goNum int)
}

func Run(c Config) {
	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0; i < c.Concurrency; i++ {
		go func(i int) {
			for {
				c.Operation(i)
			}
		}(i)
	}

	wg.Wait()
}
