package btree

type BNode struct {
	data []byte
}

const (
	BNODE_INTERNAL = 1
	BNODE_LEAF     = 2
)
