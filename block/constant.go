package block

import "github.com/ethereum/go-ethereum/common/math"

const (
	DIFFICULTY = 24
	maxNonce   = math.MaxInt64

	dbPath       = "../assets/"
	dbFile       = "blockchain.db"
	BLOCKSBUCKET = "blocks"
)
