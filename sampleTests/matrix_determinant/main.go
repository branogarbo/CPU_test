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
	c.Run(30, func(gorNum int) {
		m := u.PopulateNewMat(u.MatPopConfig{
			MainMat: u.InitMat(5, 5),
			Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
				return rand.Float64()*200 - 100
			},
		})

		fmt.Println(m)

		result, err := determinant.MatDet(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("goroutine number", gorNum, ":", result)
	})
}
