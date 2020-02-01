package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"math/big"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	//1,ESCSAで３２バイトのprivate key と６４バイトのpublic keyを生成する
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	//2,３２バイトのpublic keyをsha256でハッシュ化する
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	//3,sha3256でハッシュ化した文字列をRIPEMD-160でハッシュ化(20バイト)
	//ripemdの方が短いハッシュを作れる
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	//4,RIPEMD-160でハッシュ化した結果の先頭にバージョンのバイトを加える
	vd4 := make([]byte, 21)
	vd4[0] = 0X00
	copy(vd4[1:], digest3[:])
	//5,さらにsha256でハッシュ化する
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	//6,前回の結果をさらにハッシュ化する
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	//7,前回の結果からチェックサムとして先頭の４バイトを取り出す
	chsum := digest6[:4]
	//8,ripemd160で生成した２１バイトの文字列のチェックサムを加える(25バイト)
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	//9,base58にエンコードする
	address := base58.Encode(dc8)
	w.blockchainAddress = address

	return w
}

//type PrivateKey struct {
//	PublicKey
//	D *big.Int
//}
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

//type PublicKey struct {
//	elliptic.Curve
//	X, Y *big.Int
//}
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAdderss    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipent string, value float32) *Transaction {
	return &Transaction{privateKey, publicKey, sender, recipent, value}
}

func (t *Transaction) GenerateSignature() *Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r , s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &Signature{r,s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"_blockchain_address"`
		Recipient string  `json:"_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAdderss,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}