package util

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/branogarbo/mtrix/inverse"
	u "github.com/branogarbo/mtrix/util"
)

func PopulateMat(rows, cols int) u.MatVal {
	m := make(u.MatVal, rows)

	for r := 0; r < rows; r++ {
		m[r] = make(u.Row, cols)

		for c := 0; c < cols; c++ {
			m[r][c] = rand.Float64()*200 - 100
		}
	}

	return m
}

func UnitOperation(num int) {
	m := PopulateMat(3, 3)

	det, err := inverse.MatInv(u.Matrix{
		Value: m,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("goroutine number", num, ":", det)
}

func RunOps(num int) {
	for {
		UnitOperation(num)
	}
}

func Run(goroutines int) {
	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0; i < goroutines; i++ {
		go RunOps(i)
	}

	wg.Wait()
}
