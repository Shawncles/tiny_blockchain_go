package main

import(
  "fmt"
  "bytes"
  "crypto/sha256"
  "time"
  "strconv"
)

type Block struct {
  Timestamp     int64
  Data          []byte
  PreBlockHash  []byte
  Hash          []byte
}

func (b *Block) setHash() {
  Timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  header := bytes.Join([][]byte{b.PreBlockHash, b.Data, Timestamp}, []byte{})
  hash := sha256.Sum256(header)
  b.Hash = hash[:]
}

func NewBlock(data string, PreBlockHash []byte) (*Block) {
  block := &Block{time.Now().Unix(), []byte(data), PreBlockHash, []byte{}}
  block.setHash()
  return block
}

type Blockchain struct {
  block []*Block
}

func (bc *Blockchain) AddBlock(data string) {
  prevBlock := bc.block[len(bc.block) - 1]
  newBlock := NewBlock(data, prevBlock.Hash)
  bc.block = append(bc.block, newBlock)
}

func GenesisBlock() *Block {
  return NewBlock("GenesisBlock", []byte{})
}

func NewBlockchain() *Blockchain {
  return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {
  bc := NewBlockchain()
  bc.AddBlock("sent 1 BTC to alpha")
  bc.AddBlock("sent 2 BTC to beta")
  for _, b := range bc.block {
    fmt.Println(b)
  }
}
