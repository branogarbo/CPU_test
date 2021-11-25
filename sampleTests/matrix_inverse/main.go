package main

import (
	"fmt"
	"log"
	"math/rand"

	c "github.com/branogarbo/CPU_test"
	"github.com/branogarbo/mtrix/inverse"
	u "github.com/branogarbo/mtrix/util"
)

func main() {
	config := c.Config{
		Concurrency: 10,
		Operation: func(goNum int) {
			m := u.PopulateNewMat(u.MatPopConfig{
				MainMat: u.InitMat(9, 9),
				Action: func(m u.Matrix, r, c int, secMs []u.Matrix) float64 {
					return rand.Float64()*200 - 100
				},
			})

			result, err := inverse.MatInv(m)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("goroutine number", goNum, ":", result)
		},
	}

	c.Run(config)
}
