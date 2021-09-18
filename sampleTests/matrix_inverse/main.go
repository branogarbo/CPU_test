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
	c.Run(5, func(gorNum int) {
		m := u.PopulateNewMat(u.MatPopConfig{
			MainMat: u.InitMat(9, 9),
			Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
				return rand.Float64()*200 - 100
			},
		})

		result, err := inverse.MatInv(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("goroutine number", gorNum, ":", result)
	})
}
