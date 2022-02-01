package tests

import (
	"math/big"
	"testing"

	"github.com/adridevelopsthings/gorsa/rsa"
)

func TestGCD(t *testing.T) {
	gcds := make([]rsa.GCDData, 0)
	ggt, _ := rsa.GCD(big.NewInt(156), big.NewInt(66), gcds)
	if ggt.Cmp(big.NewInt(6)) != 0 {
		t.Errorf("GGT of 156 and 66 is 6 but GCD algorithm says %d", ggt.Int64())
	}
}

func TestEGCD(t *testing.T) {
	gcds := make([]rsa.GCDData, 0)
	ggt, gcds := rsa.GCD(big.NewInt(31), big.NewInt(26), gcds)
	if ggt.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("GGT of 156 and 66 is 6 but GCD algorithm says %d", ggt.Int64())
	}
	egcds := rsa.ConvertGCDtoEGCD(gcds)
	x, y := rsa.EGCD(egcds, uint(len(egcds)-1))
	if x.Cmp(big.NewInt(-5)) != 0 {
		t.Errorf("x,y of 31 * x + 26 * y = ggT(31,26) should be x=-5, y=6 but x=%d\n", x.Int64())
	}
	if y.Cmp(big.NewInt(6)) != 0 {
		t.Errorf("x,y of 31 * x + 26 * y = ggT(31,26) should be x=-5, y=6 but y=%d\n", y.Int64())
	}
}
