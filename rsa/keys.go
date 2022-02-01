package rsa

import (
	"encoding/binary"
	"io"
	"math/big"
)

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type SecretKey struct {
	N *big.Int
	D *big.Int
}

func (pk *PublicKey) Bytes(buf io.Writer) {
	pk_n := pk.N.Bytes()
	pk_e := pk.E.Bytes()
	binary.Write(buf, binary.LittleEndian, uint32(len(pk_n)))
	binary.Write(buf, binary.LittleEndian, uint32(len(pk_e)))
	binary.Write(buf, binary.LittleEndian, pk_n)
	binary.Write(buf, binary.LittleEndian, pk_e)
}

func (pk *PublicKey) SetBytes(reader io.Reader) {
	var N_Length uint32
	binary.Read(reader, binary.LittleEndian, &N_Length)
	var E_Length uint32
	binary.Read(reader, binary.LittleEndian, &E_Length)
	N := make([]byte, N_Length)
	reader.Read(N)

	pk.N = big.NewInt(0)
	pk.N.SetBytes(N)
	N = nil
	E := make([]byte, E_Length)
	reader.Read(E)
	pk.E = big.NewInt(0)
	pk.E.SetBytes(E)
	E = nil
}

func (sk *SecretKey) Bytes(buf io.Writer) {
	sk_n := sk.N.Bytes()
	sk_d := sk.D.Bytes()
	binary.Write(buf, binary.LittleEndian, uint32(len(sk_n)))
	binary.Write(buf, binary.LittleEndian, uint32(len(sk_d)))
	binary.Write(buf, binary.LittleEndian, sk_n)
	binary.Write(buf, binary.LittleEndian, sk_d)
}

func (sk *SecretKey) SetBytes(reader io.Reader) {
	var N_Length uint32
	binary.Read(reader, binary.LittleEndian, &N_Length)
	var D_Length uint32
	binary.Read(reader, binary.LittleEndian, &D_Length)
	N := make([]byte, N_Length)
	reader.Read(N)

	sk.N = big.NewInt(0)
	sk.N.SetBytes(N)
	N = nil
	D := make([]byte, D_Length)
	reader.Read(D)
	sk.D = big.NewInt(0)
	sk.D.SetBytes(D)
	D = nil
}
