package btree

import (
	"encoding/binary"

	"github.com/Mohammad-y-abbass/database-from-scratch/internal/helpers"
)

const (
	BNODE_INTERNAL = 1
	BNODE_LEAF     = 2
)

type BNode struct {
	data []byte
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data, btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data)
}

func (node BNode) nKeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[2:4])
}

func (node BNode) getPtr(index uint16) uint64 {
	helpers.Assert(index < node.nKeys())
	pos := HEADER + 8*index
	return binary.LittleEndian.Uint64(node.data[pos:])
}

func (node BNode) setPtr(index uint16, val uint64) {
	helpers.Assert(index < node.nKeys())
	pos := HEADER + 8*index
	binary.LittleEndian.PutUint64(node.data[pos:], val)
}

func offsetPos(node BNode, index uint16) uint16 {
	helpers.Assert(1 <= index && index <= node.nKeys())
	return HEADER + 8*node.nKeys() + 2*(index-1)
}

func (node BNode) getOffset(index uint16) uint16 {
	if index == 0 {
		return 0
	}

	return binary.LittleEndian.Uint16(node.data[offsetPos(node, index):])
}

func (node BNode) setOffset(index uint16, offset uint16) {
	binary.LittleEndian.PutUint16(node.data[offsetPos(node, index):], offset)
}

func (node BNode) kvPos(index uint16) uint16 {
	helpers.Assert(index <= node.nKeys())
	return HEADER + 8*node.nKeys() + 2*node.nKeys() + node.getOffset(index)
}

func (node BNode) getKey(index uint16) []byte {
	helpers.Assert(index < node.nKeys())
	pos := node.kvPos(index)
	klen := binary.LittleEndian.Uint16(node.data[pos:])
	return node.data[pos+4:][:klen]
}
func (node BNode) getVal(index uint16) []byte {
	helpers.Assert(index < node.nKeys())
	pos := node.kvPos(index)
	klen := binary.LittleEndian.Uint16(node.data[pos+0:])
	vlen := binary.LittleEndian.Uint16(node.data[pos+2:])
	return node.data[pos+4+klen:][:vlen]
}

func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nKeys())
}
