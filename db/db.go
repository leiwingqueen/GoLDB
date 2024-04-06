package db

// interface of kv store
type DB interface {
	// get value by key
	Get(key []byte) ([]byte, error)
	// set value by key
	Set(key []byte, value []byte) error
	// delete value by key
	Delete(key []byte) error
}
