package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//Blockのstructを定義
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

//新しいブロックを生成
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

//block表示した時に見やすくするためのメソッド
func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previous_hash         %s\n", b.previousHash)
	fmt.Printf("transactions          %s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

//新規でblockchainを生成する
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

//生成したblockを元にblockchainを生成するためのメソッド
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	//chainに生成したBlockを追加していく
	bc.chain = append(bc.chain, b)
	return b
}

//blockchainを表示した時に見やすくするためのメソッド
func (bc *Blockchain) Print() {
	//chainの中身を取り出す
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=",25),i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	//logを生成するための処理
	log.SetPrefix("Blockchain:  ")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
	blockChain.CreateBlock(3, "hash 3")

}
