package main

import (
	"fmt"
	"log"
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

//表示を見やすくするためのメソッド
func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previous_hash         %s\n", b.previousHash)
	fmt.Printf("transactions          %s\n", b.transactions)
}

func init()  {
	//logを生成するための処理
	log.SetPrefix("Blockchain:  ")
}

func main()  {
	b := NewBlock(0, "init hash")
	b.Print()
}
