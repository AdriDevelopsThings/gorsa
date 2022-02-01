package tests

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/adridevelopsthings/gorsa/rsa"
)

func TestRSAKeys(t *testing.T) {
	_, _, err := rsa.GenerateKeyPair(1024)
	if err != nil {
		t.Errorf("Error while generating rsa keys %v\n", err)
	}
}

func TestRSAEncryption(t *testing.T) {
	pk, sk, err := rsa.GenerateKeyPair(1024)
	if err != nil {
		t.Errorf("Error while generating rsa keys %v\n", err)
	}
	message, err := rand.Int(rand.Reader, pk.E)
	if err != nil {
		t.Errorf("Error while generating random message %v\n", err)
	}
	chiffre := rsa.EncryptInt(pk, message)
	newMessage := rsa.DecryptInt(sk, chiffre)
	if message.Cmp(newMessage) != 0 {
		t.Errorf("Error while encrypting and decrypting. message is not equal to decrypted message. message was %x and new message is %x.\n", message.Bytes(), newMessage.Bytes())
	}
}

func TestRSAKeysBytes(t *testing.T) {
	pk, sk, err := rsa.GenerateKeyPair(64)
	if err != nil {
		t.Errorf("Error while generating rsa keys %v\n", err)
	}
	pk_n_o := pk.N.Bytes()
	pk_e_o := pk.E.Bytes()
	pk_buf := new(bytes.Buffer)
	pk.Bytes(pk_buf)
	pk = &rsa.PublicKey{}
	pk_buf_reader := bytes.NewReader(pk_buf.Bytes())
	pk.SetBytes(pk_buf_reader)
	pk_n_n := pk.N.Bytes()
	pk_e_n := pk.E.Bytes()
	if bytes.Compare(pk_n_o, pk_n_n) != 0 || bytes.Compare(pk_e_o, pk_e_n) != 0 {
		t.Errorf("Error while decoding pk and encoding it again. Old values N=%x E=%x, new values N=%x E=%x\n", pk_n_o, pk_e_o, pk_n_n, pk_e_n)
	}

	sk_n_o := sk.N.Bytes()
	sk_d_o := sk.D.Bytes()
	sk_buf := new(bytes.Buffer)
	sk.Bytes(sk_buf)
	sk = &rsa.SecretKey{}
	sk_buf_reader := bytes.NewReader(sk_buf.Bytes())
	sk.SetBytes(sk_buf_reader)
	sk_n_n := sk.N.Bytes()
	sk_d_n := sk.D.Bytes()
	if bytes.Compare(sk_n_o, sk_n_n) != 0 || bytes.Compare(sk_d_o, sk_d_n) != 0 {
		t.Errorf("Error while decoding sk and encoding it again. Old values N=%x D=%x, new values N=%x D=%x\n", sk_n_o, sk_d_o, sk_n_n, sk_d_n)
	}
}
