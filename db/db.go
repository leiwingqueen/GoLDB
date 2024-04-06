package db

// LevelDB DB interface of kv store
type LevelDB interface {
	// Get value by key
	Get(key []byte) ([]byte, error)
	// Set value by key
	Set(key []byte, value []byte) error
	// Delete value by key
	Delete(key []byte) error
}

// LevelDBImpl implementation of DB interface
type LevelDBImpl struct {
}

func (db *LevelDBImpl) Set(key []byte, value []byte) error {
	//TODO implement me
	panic("implement me")
}

func (db *LevelDBImpl) Delete(key []byte) error {
	//TODO implement me
	panic("implement me")
}

// Get value by key
func (db *LevelDBImpl) Get(key []byte) ([]byte, error) {
	return []byte{}, nil
}
