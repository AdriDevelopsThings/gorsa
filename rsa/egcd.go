package rsa

import (
	"math/big"
)

type GCDData struct {
	a *big.Int
	b *big.Int
	q *big.Int
}

type EGCDData struct {
	GCDData
	x *big.Int
	y *big.Int
}

func GCD(a, b *big.Int, gcds []GCDData) (*big.Int, []GCDData) {
	q, r := big.NewInt(0).DivMod(a, b, big.NewInt(0)) // q = a / b and r = a mod b
	gcds = append(gcds, GCDData{a, b, q})
	if r.Cmp(big.NewInt(0)) == 0 { // if r = 0: end of algorithm
		return b, gcds // b is ggT(a,b)
	} else {
		return GCD(b, r, gcds)
	}
}

func EGCD(egcds []EGCDData, index uint) (*big.Int, *big.Int) {
	var x, y *big.Int
	if index == uint(len(egcds)-1) {
		// starting values
		x = big.NewInt(0)
		y = big.NewInt(1)
	} else {
		x = egcds[index+1].y                                                                          // x(i) = y(i + 1)
		y = egcds[index].y.Sub(egcds[index+1].x, big.NewInt(0).Mul(egcds[index].q, egcds[index+1].y)) // y(i) = x(i+1) - q(i) * y(i +1)
	}
	if index == 0 {
		return x, y // end of algorithm
	}
	egcds[index].x = x
	egcds[index].y = y
	return EGCD(egcds, index-1)
}

func ConvertGCDtoEGCD(gcds []GCDData) []EGCDData {
	egcds := make([]EGCDData, len(gcds))
	for index, element := range gcds {
		egcds[index] = EGCDData{element, big.NewInt(0), big.NewInt(0)}
	}
	return egcds
}
