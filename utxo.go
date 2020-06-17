package bench

type TxID [32]byte
type txoFlags uint8
type Outpoint struct {
	TxID  TxID
	Index uint32
}

type UTXOEntry struct {
	amount         uint64
	scriptPubKey   []byte
	blockBlueScore uint64
	packedFlags    txoFlags
}

type utxoCollection map[Outpoint]*UTXOEntry

type keysArray []Outpoint
type OrderedUtxoCollection struct {
	m    utxoCollection
	keys keysArray
}

func removeIndex(s keysArray, index int) keysArray {
	return append(s[:index], s[index+1:]...)
}

func NewOrderedUtxoCollection() *OrderedUtxoCollection {
	return &OrderedUtxoCollection{
		m:    make(utxoCollection),
		keys: make(keysArray, 0),
	}
}
func (uc *OrderedUtxoCollection) findIndex(lookFor Outpoint) (int, bool) {
	for i, value := range uc.keys {
		if value == lookFor {
			return i, true
		}
	}
	return 0, false
}

// add adds a new UTXO entry to this collection
// if exists it will modify the existing value and will move the key to the end of the list.
func (uc *OrderedUtxoCollection) Add(outpoint Outpoint, entry *UTXOEntry) {
	if uc.Contains(outpoint) {
		// Delete the old index
		index, _ := uc.findIndex(outpoint)
		uc.keys = removeIndex(uc.keys, index)
	}
	uc.m[outpoint] = entry
	uc.keys = append(uc.keys, outpoint)
}

// remove removes a UTXO entry from this collection if it exists
func (uc *OrderedUtxoCollection) Remove(outpoint Outpoint) {
	delete(uc.m, outpoint)
	index, found := uc.findIndex(outpoint)
	if found {
		uc.keys = removeIndex(uc.keys, index)
	}
}

// get returns the UTXOEntry represented by provided outpoint,
// and a boolean value indicating if said UTXOEntry is in the set or not
func (uc *OrderedUtxoCollection) Get(outpoint Outpoint) (*UTXOEntry, bool) {
	entry, ok := uc.m[outpoint]
	return entry, ok
}

// contains returns a boolean value indicating whether a UTXO entry is in the set
func (uc *OrderedUtxoCollection) Contains(outpoint Outpoint) bool {
	_, ok := uc.m[outpoint]
	return ok
}
