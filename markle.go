package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// Read transactions from input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var transactions [][]byte
	for scanner.Scan() {
		transaction, err := hex.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}

	// Build Merkle Tree
	rootHash := buildMerkleTree(transactions)

	// Output root hash
	fmt.Println(hex.EncodeToString(rootHash))
}

func buildMerkleTree(transactions [][]byte) []byte {
	var merkleTree [][]byte

	// Add transactions to the first level of the Merkle Tree
	for _, transaction := range transactions {
		hash := sha256.Sum256(transaction)
		merkleTree = append(merkleTree, hash[:])
	}

	// Repeat until there is only one hash left (the root hash)
	for len(merkleTree) > 1 {
		var level []byte

		// If there is an odd number of hashes, duplicate the last one
		if len(merkleTree)%2 != 0 {
			merkleTree = append(merkleTree, merkleTree[len(merkleTree)-1])
		}

		// Combine adjacent hashes in pairs and hash them together
		for i := 0; i < len(merkleTree); i += 2 {
			combined := append(merkleTree[i], merkleTree[i+1]...)
			hash := sha256.Sum256(combined)
			level = append(level, hash[:]...)
		}

		// merkleTree = level
	}

	return merkleTree[0]
}
