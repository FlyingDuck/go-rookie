package main

import (
	"fmt"
	"github.com/FlyingDuck/go-rookie/alg"
)

func main() {
	//q := alg.CutRod(alg.Prices, 4)
	//q := alg.MemoCutRod(alg.Prices, 4)
	//q := alg.BottomUpCutRod(alg.Prices, 4)

	r, sols := alg.ExtendedBottomUpCutRod(alg.Prices, 10)
	fmt.Println(r)
	fmt.Println(sols)


	//sentence := "“Other than 中国 is good"
	//ascii, replaces := alg.ConvertSentence2ASCII(context.Background(), sentence)
	//
	//rStc := []rune(ascii)
	//for idx, r := range replaces {
	//
	//}
}
