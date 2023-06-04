package block

import "github.com/boltdb/bolt"

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (b *BlockchainIterator) Next() *Block {
	var block *Block

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKSBUCKET))
		encodedBlock := bucket.Get(b.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		panic(err)
	}

	b.currentHash = block.PrevBlockHash

	return block
}
