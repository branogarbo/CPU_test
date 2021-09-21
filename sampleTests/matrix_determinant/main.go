package main

import (
	"fmt"
	"log"
	"math/rand"

	c "github.com/branogarbo/CPU_test"
	"github.com/branogarbo/mtrix/determinant"
	u "github.com/branogarbo/mtrix/util"
)

func main() {
	config := c.Config{
		Concurrency: 30,
		Operation: func(goNum int) {
			m := u.PopulateNewMat(u.MatPopConfig{
				MainMat: u.InitMat(5, 5),
				Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
					return rand.Float64()*200 - 100
				},
			})

			result, err := determinant.MatDet(m)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("goroutine number", goNum, ":", result)
		},
	}

	c.Run(config)
}
