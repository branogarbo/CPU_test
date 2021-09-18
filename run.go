package CPU_test

import (
	"sync"
)

func Run(concurrencies int, unitOperation func(gorNum int)) {
	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0; i < concurrencies; i++ {
		go func(i int) {
			for {
				unitOperation(i)
			}
		}(i)
	}

	wg.Wait()
}
