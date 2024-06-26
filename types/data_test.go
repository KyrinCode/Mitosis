package types

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestDataTemplate(t *testing.T) {
	data := NewDataTemplate(10000, "parameter2", make([]byte, 20))
	dataBytes := data.MarshalBinary()
	var n_data DataTemplate
	n_data.UnmarshalBinary(dataBytes)
	assert.Equal(t, data, &n_data)
}

func TestUnsignedTransaction(t *testing.T) {
	fromAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	toAddr := [20]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	tx := NewTransaction(1, 1, fromAddr, toAddr, uint64(1000), make([]byte, 0))

	fmt.Println(tx.Hash)

	tx_json := tx.MarshalJson()
	new_tx, err := NewTransactionFromJson(tx_json)
	assert.Nil(t, err)
	assert.Equal(t, tx, new_tx)

	tx_bytes := tx.MarshalBinary()
	var n_tx Transaction
	n_tx.UnmarshalBinary(tx_bytes)
	assert.Nil(t, err)
	assert.Equal(t, tx, &n_tx)
}

func TestOutboundChunk(t *testing.T) {
	fromAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	toAddr := [20]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	tx := NewTransaction(1, 1, fromAddr, toAddr, uint64(1000), make([]byte, 0))
	txs := []Transaction{*tx, *tx}
	chunkProof := [][]byte{[]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}
	outboundChunk := NewOutboundChunk(common.Hash{}, txs, chunkProof)

	c_bytes := outboundChunk.MarshalBinary()
	var n_chunk OutboundChunk
	n_chunk.UnmarshalBinary(c_bytes)
	assert.Equal(t, outboundChunk, &n_chunk)
}

/*

func TestBlock(t *testing.T) {
	fromAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	toAddr := [20]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	tx := NewTransaction(fromAddr, toAddr, uint16(1), uint16(1), uint64(1000), uint64(10000))
	txs := []Transaction{*tx, *tx}
	proof := [][]byte{[]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}

	intxs := NewInShardTransactions(txs, proof)

	proofbytes, _ := rlp.EncodeToBytes(proof)
	hashPreBlock := crypto.Keccak256Hash(proofbytes)
	hashPreRoot := crypto.Keccak256Hash([]byte{1, 2})

	h := NewHeader(uint16(1), uint32(1), hashPreBlock, hashPreRoot, hashPreBlock, hashPreRoot, uint64(time.Now().Unix()), hashPreRoot)

	b := NewBlock(*h, txs, []InShardTransactions{*intxs})

	h_bytes := h.MarshalBinary()
	var n_h Header
	n_h.UnmarshalBinary(h_bytes)
	assert.Equal(t, h, &n_h)

	b_bytes := b.MarshalBinary()
	var n_b Block
	n_b.UnmarshalBinary(b_bytes)
	assert.Equal(t, b, &n_b)
}

func TestBFTMessage(t *testing.T) {

	fromAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	toAddr := [20]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	tx := NewTransaction(fromAddr, toAddr, uint16(1), uint16(1), uint64(1000), uint64(10000))
	txs := []Transaction{*tx, *tx}
	proof := [][]byte{[]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}

	intxs := NewInShardTransactions(txs, proof)

	proofbytes, _ := rlp.EncodeToBytes(proof)
	hashPreBlock := crypto.Keccak256Hash(proofbytes)
	hashPreRoot := crypto.Keccak256Hash([]byte{1, 2})

	h := NewHeader(uint16(1), uint32(1), hashPreBlock, hashPreRoot, hashPreBlock, hashPreRoot, uint64(time.Now().Unix()), hashPreRoot)

	b := NewBlock(*h, txs, []InShardTransactions{*intxs})

	privateKey, _ := crypto.GenerateKey()
	signHash := b.Hash()
	signature, _ := crypto.Sign(signHash[:], privateKey)

	pro := NewProposalMsg(b.Height, b.Hash(), *b, signature)
	pro_bytes := pro.MarshalBinary()
	var n_pro BFTMessage
	n_pro.UnmarshalBinary(pro_bytes)
	assert.Equal(t, pro, &n_pro)

	pro2 := NewPreCommitMsg(b.Height, b.Hash(), signature)
	pro2_bytes := pro2.MarshalBinary()
	var n_pro2 BFTMessage
	n_pro2.UnmarshalBinary(pro2_bytes)
	assert.Equal(t, pro2, &n_pro2)

	pro3 := NewCommitMsg(b.Height, b.Hash(), signature)
	pro_bytes3 := pro3.MarshalBinary()
	var n_pro3 BFTMessage
	n_pro3.UnmarshalBinary(pro_bytes3)
	assert.Equal(t, pro3, &n_pro3)

}



func TestCrossMessage(t *testing.T) {

	fromAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	toAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	tx := NewTransaction(fromAddr, toAddr, uint16(1), uint16(1), uint64(1000), uint64(10000))

	//check tx
	data := tx.MarshalBinary()
	fmt.Println(data)
	var ntx Transaction
	err := ntx.UnmarshalBinary(data)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, tx, &ntx)
	nntx := tx.Copy().(Transaction)
	assert.Equal(t, tx, &nntx)

	//check inshardtx
	txs := []Transaction{*tx, *tx}
	proof := [][]byte{[]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}, []byte{0x09, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}, []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}
	hash := crypto.Keccak256Hash([]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08})
	root := crypto.Keccak256Hash([]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08})
	intxs := NewInShardTransactions(txs, proof, hash, root)

	data = intxs.MarshalBinary()
	fmt.Println(data)
	var nintxs InShardTransactions
	err = nintxs.UnmarshalBinary(data)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, intxs, &nintxs)
	nnintxs := intxs.Copy().(InShardTransactions)
	assert.Equal(t, intxs, &nnintxs)

	// check Header
	hashPreBlock := crypto.Keccak256Hash([]byte{1, 2})
	hashPreRoot := crypto.Keccak256Hash([]byte{1, 2})

	h := NewHeader(uint16(1), uint32(1), hashPreBlock, hashPreRoot, hashPreBlock, hashPreRoot, uint64(time.Now().Unix()), hashPreRoot)
	data = h.MarshalBinary()
	fmt.Println(data)
	var nintxs InShardTransactions
	err = nintxs.UnmarshalBinary(data)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, intxs, &nintxs)

	b := NewBlock(*h, txs, []InShardTransactions{*intxs})


		proofbytes, _ := rlp.EncodeToBytes(proof)
		hashPreBlock := crypto.Keccak256Hash(proofbytes)
		hashPreRoot := crypto.Keccak256Hash([]byte{1, 2})

		h := NewHeader(uint16(1), uint32(1), hashPreBlock, hashPreRoot, hashPreBlock, hashPreRoot, uint64(time.Now().Unix()), hashPreRoot)

		b := NewBlock(*h, txs, []InShardTransactions{*intxs})

		privateKey, _ := crypto.GenerateKey()
		signHash := b.Hash()
		signature, _ := crypto.Sign(signHash[:], privateKey)
		privateKey2, _ := crypto.GenerateKey()
		signature2, _ := crypto.Sign(signHash[:], privateKey2)

		cro := NewCrossShardMsg(uint16(2), *h, *intxs, [][]byte{signature, signature2})

		cro_bytes := cro.MarshalBinary()
		var n_cro CrossShardMsg
		n_cro.UnmarshalBinary(cro_bytes)
		assert.Equal(t, cro, &n_cro)




		cro2 := NewValidateResultMsg(b.Hash(), h.ShardId, h.Height, [][]byte{signature, signature2})
		cro2_bytes := cro2.MarshalBinary()
		var n_cro2 ValidateResultMsg
		n_cro2.UnmarshalBinary(cro2_bytes)
		assert.Equal(t, cro2, &n_cro2)

		cro3 := NewValidateRequestMsg(*b, [][]byte{signature, signature2})
		cro3_bytes := cro3.MarshalBinary()
		var n_cro3 ValidateRequestMsg
		n_cro3.UnmarshalBinary(cro3_bytes)
		assert.Equal(t, cro3, &n_cro3)

		cro4 := NewValidateRespondMsg(b.Hash(), h.ShardId, h.Height, signature)
		cro4_bytes := cro4.MarshalBinary()
		var n_cro4 ValidateRespondMsg
		n_cro4.UnmarshalBinary(cro4_bytes)
		assert.Equal(t, cro4, &n_cro4)



}

type ProofList [][]byte

func (n *ProofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}
func (n *ProofList) Delete(key []byte) error {
	panic("not supported")
}
*/
/*
func TestDeriveSha(t *testing.T) {

	fromAddr := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	toAddr := [20]byte{0x00, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	tx := NewTransaction(fromAddr, toAddr, uint16(1), uint16(1), uint64(1000), uint64(10000))
	txs := []Transaction{*tx, *tx}
	intxs := NewInShardTransactions(txs, nil)
	LatestOutTXs := []InShardTransactions{*intxs, *intxs}

	txTrie := new(trie.Trie)
	root := DeriveSha(LatestOutTXs, txTrie)
	keybuf := new(bytes.Buffer)
	for i := range LatestOutTXs {
		var px ProofList
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i+1))
		txTrie.Prove(keybuf.Bytes(), 0, &px)
		LatestOutTXs[i].TxProof = px
	}

	proof := LatestOutTXs[1].TxProof
	db := memorydb.New()
	for j := 0; j < len(proof); j++ {
		db.Put(crypto.Keccak256(proof[j]), proof[j])
	}

	key, _ := rlp.EncodeToBytes(uint(2))
	txvale, _ := rlp.EncodeToBytes(LatestOutTXs[0].Txs)
	fmt.Println(txvale)
	txhash := crypto.Keccak256Hash(txvale)
	fmt.Println(key, txhash.Bytes())

	hash, _ := trie.VerifyProof(root, key, db)
	fmt.Println(hash)
	assert.Equal(t, common.BytesToHash(hash), txhash)

}

*/

/*
func TestBlockHeader(t *testing.T) {
	prevHash := [32]byte{0xbf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	rootHash := [32]byte{0xba, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	prerootHash := [32]byte{0xbe, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	txHash := [32]byte{0xbd, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	h := NewBlockHeader(1, 1, prevHash, prerootHash, rootHash, txHash, uint64(time.Now().Unix()))

	h_byte := h.MarshalBinary()
	assert.NotNil(t, h_byte)
	var h_new BlockHeader
	err := h_new.UnmarshalBinary(h_byte)
	assert.Nil(t, err)
	assert.Equal(t, h, &h_new)

}

func TestBlock(t *testing.T) {
	fromAddr := [32]byte{0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	toAddr := [32]byte{0xbf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	tx1 := NewTransaction(fromAddr, toAddr, uint64(1000), uint64(10000))
	tx2 := NewTransaction(toAddr, fromAddr, uint64(1000), uint64(10000))

	txs := []Transaction{*tx1, *tx2}

	prevHash := [32]byte{0xbf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	prerootHash := [32]byte{0xbe, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	rootHash := [32]byte{0xba, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	txHash := [32]byte{0xba, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0xaf, 0xad, 0x12, 0x13, 0x14, 0x15, 0x16, 0x36}
	h := NewBlockHeader(1, 1, prevHash, prerootHash, rootHash, txHash, uint64(time.Now().Unix()))

	blk := NewBlock(*h, txs)

	blk_json := blk.MarshalJson()
	new_blk, err := NewBlockFromJson(blk_json)
	assert.Nil(t, err)
	assert.Equal(t, blk, new_blk)

	blk_bytes := blk.MarshalBinary()
	var blk_new Block
	err = blk_new.UnmarshalBinary(blk_bytes)
	assert.Nil(t, err)

	assert.Equal(t, blk, &blk_new)
}

*/
