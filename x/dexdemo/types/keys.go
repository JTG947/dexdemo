package types

const (
	// ModuleName defines the module name
	ModuleName = "dexdemo"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dexdemo"
)

var (
	ParamsKey = []byte("p_dexdemo")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
