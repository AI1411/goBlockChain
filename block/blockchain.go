package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFICULTY = 3
	MINING_SENDER    = "THE BLOCK CHAIN"
	MINING_REWORD    = 1.0
)

//Blockのstructを定義
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

//新しいブロックを生成
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

//block表示した時に見やすくするためのメソッド
func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previous_hash         %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

//hash化するためのメソッド
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

//jsonにしてからhash化するための記述
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

//新規でblockchainを生成する
func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

//生成したblockを元にblockchainを生成するためのメソッド
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	//chainに生成したBlockを追加していく
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

//最後のブロックを求めるメソッド
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

//blockchainを表示した時に見やすくするためのメソッド
func (bc *Blockchain) Print() {
	//chainの中身を取り出す
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

//transactionをpoolに追加していく
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(
			t.senderBlockChainAddress,
			t.recipientBlockChainAddress,
			t.value))
	}
	return transactions
}

//oがdifficultyの数だけ並ぶhash値を求めるメソッド
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	//正しいnonceになるまでforで回す
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFICULTY) {
		nonce += 1
	}
	return nonce
}

//miningできたかどうか判定
func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWORD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

//送受信があった金額を計算する処理
func (bc *Blockchain) CalculateTotalAmount(blockchainAdress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if blockchainAdress == t.recipientBlockChainAddress {
				totalAmount += value
			}
			if blockchainAdress == t.senderBlockChainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}


//transactionのstruct
type Transaction struct {
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}

//新規トランザクションを生成
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address     %s\n", t.senderBlockChainAddress)
	fmt.Printf(" sender_blockchain_address     %s\n", t.recipientBlockChainAddress)
	fmt.Printf(" value                         %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockChainAddress    string  `json:"sender_block_chain_address"`
		RecipientBlockChainAddress string  `json:"recipient_block_chain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockChainAddress:    t.senderBlockChainAddress,
		RecipientBlockChainAddress: t.recipientBlockChainAddress,
		Value:                      t.value,
	})
}
