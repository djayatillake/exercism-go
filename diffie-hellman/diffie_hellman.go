// package diffiehellman implements the diffie-hellman-merkely key exchange
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey creates a private key less than p
func PrivateKey(p *big.Int) *big.Int {
	private, err := rand.Int(rand.Reader, p)
	if err != nil {
		panic(err)
	}
	if private.BitLen() <= 1 {
		return private.Add(private, big.NewInt(2))
	}
	return private
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	g_big := big.NewInt(g)
	return new(big.Int).Exp(g_big, private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	priv := PrivateKey(p)
	return priv, PublicKey(priv, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
