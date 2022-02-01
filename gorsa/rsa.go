package gorsa

import (
	"crypto/rand"
	"math/big"
)

func GenerateKeyPair(keylength int) (*PublicKey, *SecretKey, error) {
	// generate p, q, N, phi(N)
	var p, q, N *big.Int
	var err error
	p, err = rand.Prime(rand.Reader, keylength)
	if err != nil {
		return nil, nil, err
	}
	q, err = rand.Prime(rand.Reader, keylength)
	if err != nil {
		return nil, nil, err
	}
	N.Mul(p, q)
	var phi_N *big.Int
	phi_N.Mul(p.Sub(p, big.NewInt(1)), p.Sub(q, big.NewInt(1)))
	p, q = nil, nil // remove because we don't need p & q now
	e, err := rand.Prime(rand.Reader, keylength-1)
	if err != nil {
		return nil, nil, err
	}
	publickey := PublicKey{N, e}

}
