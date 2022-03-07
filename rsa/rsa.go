package rsa

import (
	"crypto/rand"
	"math/big"
)

const (
	CHUNK_SIZE = 64 / 8
)

func CalculatePrivateKeyFromPhi(e, phi_N *big.Int) *big.Int {
	gcds := make([]GCDData, 0)
	_, gcds = GCD(e, phi_N, gcds)
	// ggt == 1
	egcds := ConvertGCDtoEGCD(gcds)
	x, _ := EGCD(egcds, uint(len(egcds)-1))
	return x
}

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
	N = big.NewInt(0)
	N.Mul(p, q) // N = p * q
	phi_N := big.NewInt(0)
	phi_N.Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1))) // phi(N) = (p - 1) * (q - 1)
	p, q = nil, nil                                             // remove because we don't need p & q now
	e, err := rand.Prime(rand.Reader, phi_N.BitLen()-1)         // e < phi(N)
	if err != nil {
		return nil, nil, err
	}
	publickey := PublicKey{N, e}
	secretkey := SecretKey{N, CalculatePrivateKeyFromPhi(e, phi_N)}
	return &publickey, &secretkey, nil
}

func EncryptInt(pubkey *PublicKey, message *big.Int) *big.Int {
	return big.NewInt(0).Exp(message, pubkey.E, pubkey.N)
}

func DecryptInt(secretkey *SecretKey, chiffre *big.Int) *big.Int {
	return big.NewInt(0).Exp(chiffre, secretkey.D, secretkey.N)
}
