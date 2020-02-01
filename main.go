package main

import (
	"fmt"
	"goblockchain/wallet"
)

func init() {
	//logを生成するための処理
	//log.SetPrefix("Blockchain:  ")
}

func main() {
	w := wallet.NewWallet()
	//publicKey privateKey生成して表示
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	//myBlockchainAddress := "my block chain"
	//blockChain := NewBlockchain(myBlockchainAddress)
	//blockChain.Print()
	//
	//blockChain.AddTransaction("A", "B", 1.0)
	//blockChain.Mining()
	//blockChain.Print()
	//
	//blockChain.AddTransaction("C", "D", 2.0)
	//blockChain.AddTransaction("X", "Y", 3.0)
	//blockChain.Mining()
	//blockChain.Print()
	//
	//fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my block chain"))
	//fmt.Printf("c %.1f\n", blockChain.CalculateTotalAmount("C"))
	//fmt.Printf("d %.1f\n", blockChain.CalculateTotalAmount("D"))
}
