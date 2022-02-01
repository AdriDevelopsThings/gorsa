package gorsa

import "math/big"

type PublicKey struct {
	N *big.Int
	e *big.Int
}

type SecretKey struct {
	N *big.Int
	d *big.Int
}
