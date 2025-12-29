package btree

const HEADER = 4  // size of header in bytes
const BTREE_PAGE_SIZE = 4096	// size of page in bytes
const BTREE_MAX_KEY_SIZE = 1000	// max size of key in bytes
const BTREE_MAX_VAL_SIZE = 3000	// max size of value in bytes

type BTree struct {
	root uint64  // root of the tree

	get    func(uint64) BNode // get node by page number
	new    func() uint64	// create new node and return page number
	delete func(uint64)	// delete node by page number
}


// This function will run when the package is loaded and check if the node size is valid
func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	assert(node1max <= BTREE_PAGE_SIZE)
}

func assert(condition bool) {
	if !condition {
		panic("assertion failed")
	}
}
