package mvdb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/vitelabs/go-vite/chain/pending"
	"github.com/vitelabs/go-vite/chain/utils"
	"github.com/vitelabs/go-vite/common/types"
	"path"
)

type MultiVersionDB struct {
	db *leveldb.DB

	pending *chain_pending.MemDB

	//latestKeyId   uint64
	//latestValueId uint64
}

func NewMultiVersionDB(chainDir string) (*MultiVersionDB, error) {

	dbDir := path.Join(chainDir, "state_bak")

	db, err := leveldb.OpenFile(dbDir, nil)
	if err != nil {
		return nil, err
	}

	return &MultiVersionDB{
		db:      db,
		pending: chain_pending.NewMemDB(),
	}, nil
}

func (mvDB *MultiVersionDB) Destroy() error {
	if err := mvDB.db.Close(); err != nil {
		return err
	}

	mvDB.db = nil

	mvDB.pending.Clean()
	mvDB.pending = nil

	return nil
}

func (mvDB *MultiVersionDB) LatestKeyId() uint64 {
	return mvDB.latestKeyId
}

func (mvDB *MultiVersionDB) GetKeyId(key []byte) (uint64, error) {
	keyIdKey := chain_utils.CreateKeyIdKey(key)

	keyIdBytes, ok := mvDB.pending.Get(keyIdKey)
	if !ok {
		var err error
		keyIdBytes, err = mvDB.db.Get(keyIdKey, nil)
		if err != nil {
			if err == leveldb.ErrNotFound {
				return 0, nil
			}
			return 0, err
		}
	}

	return chain_utils.BytesToUint64(keyIdBytes), nil
}

func (mvDB *MultiVersionDB) GetValueId(keyId uint64) (uint64, error) {
	latestValueKey := chain_utils.CreateLatestValueKey(keyId)

	valueIdBytes, ok := mvDB.pending.Get(latestValueKey)
	if !ok {
		var err error
		valueIdBytes, err = mvDB.db.Get(latestValueKey, nil)
		if err != nil {
			if err == leveldb.ErrNotFound {
				return 0, nil
			}

			return 0, err
		}
	}

	return chain_utils.BytesToUint64(valueIdBytes), nil
}

func (mvDB *MultiVersionDB) GetValue(key []byte) ([]byte, error) {
	keyId, err := mvDB.GetKeyId(key)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	if keyId <= 0 {
		return nil, nil
	}

	valueId, err := mvDB.GetValueId(keyId)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	if valueId <= 0 {
		return nil, nil
	}

	value, err := mvDB.GetValueByValueId(valueId)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return value, nil
}

func (mvDB *MultiVersionDB) HasValue(key []byte) (bool, error) {
	keyId, err := mvDB.GetKeyId(key)
	if err != nil {
		return false, err
	}
	if keyId <= 0 {
		return false, nil
	}

	valueId, err := mvDB.GetValueId(keyId)
	if err != nil {
		return false, err
	}
	if valueId <= 0 {
		return false, nil
	}
	return true, nil
}

func (mvDB *MultiVersionDB) GetValueByValueId(valueId uint64) ([]byte, error) {
	valueIdKey := chain_utils.CreateValueIdKey(valueId)

	value, ok := mvDB.pending.Get(valueIdKey)
	if !ok {
		var err error
		value, err = mvDB.db.Get(valueIdKey, nil)
		if err != nil {
			if err == leveldb.ErrNotFound {
				return nil, nil
			}

			return nil, err
		}
	}

	return value, nil
}

func (mvDB *MultiVersionDB) Insert(blockHash *types.Hash, addr *types.Address, keyList [][]byte, valueList [][]byte) error {
	//keySize := len(keyList)

	//prevValueIdList := make([]uint64, keySize)

	//endValueId := atomic.AddUint64(&mvDB.latestValueId, uint64(keySize))
	//startValueId := endValueId - uint64(keySize)

	for index, key := range keyList {

		storageKey := chain_utils.CreateStorageValueKey(addr, key)

		mvDB.pending.Put(blockHash, storageKey, valueList[index])

		historyKey := chain_utils.CreateHistoryStorageValueKey(addr, key, 0)

		mvDB.pending.Put(blockHash, historyKey, valueList[index])

		//keyId, err := mvDB.GetKeyId(key)
		//if err != nil {
		//	return err
		//}
		//if keyId <= 0 {
		//	keyId = atomic.AddUint64(&mvDB.latestKeyId, 1)
		//	// insert key id
		//	mvDB.pending.Put(blockHash, chain_utils.CreateKeyIdKey(key), chain_utils.Uint64ToBytes(keyId))
		//} else {
		//	prevValueId, err := mvDB.GetValueId(keyId)
		//	if err != nil {
		//		return err
		//	}
		//	prevValueIdList[index] = prevValueId
		//}
		//
		//valueId := startValueId + uint64(index+1)
		//
		//// update latest value index
		//mvDB.pending.Put(blockHash, chain_utils.CreateLatestValueKey(keyId), chain_utils.Uint64ToBytes(valueId))
		//
		//// insert value
		//mvDB.pending.Put(blockHash, chain_utils.CreateValueIdKey(valueId), valueList[index])
	}

	// insert undo log
	//if err := mvDB.writeUndoLog(blockHash, keyList, prevValueIdList); err != nil {
	//	return err
	//}
	return nil
}

func (mvDB *MultiVersionDB) Flush(blockHashList []*types.Hash) error {
	batch := new(leveldb.Batch)

	mvDB.pending.FlushList(batch, blockHashList)
	//mvDB.updateLatestLocation(batch, latestLocation)

	if err := mvDB.db.Write(batch, nil); err != nil {
		return err
	}

	mvDB.pending.DeleteByBlockHashList(blockHashList)
	return nil
}

//func (mvDB *MultiVersionDB) Flush(blockHashList []*types.Hash, latestLocation *chain_block.Location) error {
//	batch := new(leveldb.Batch)
//
//	mvDB.pending.FlushList(batch, blockHashList)
//	mvDB.updateLatestLocation(batch, latestLocation)
//
//	if err := mvDB.db.Write(batch, nil); err != nil {
//		return err
//	}
//
//	mvDB.pending.DeleteByBlockHashList(blockHashList)
//	return nil
//}

func (mvDB *MultiVersionDB) DeletePendingBlock(blockHash *types.Hash) {
	mvDB.pending.DeleteByBlockHash(blockHash)
}

//func (mvDB *MultiVersionDB) QueryLatestLocation() (*chain_block.Location, error) {
//	key := chain_utils.CreateMvDbLatestLocationKey()
//	value, err := mvDB.db.Get(key, nil)
//	if err != nil {
//		if err == leveldb.ErrNotFound {
//			return nil, nil
//		}
//		return nil, err
//	}
//
//	return chain_utils.DeserializeLocation(value), nil
//}

//func (mvDB *MultiVersionDB) updateLatestLocation(batch *leveldb.Batch, latestLocation *chain_block.Location) {
//	batch.Put(chain_utils.CreateMvDbLatestLocationKey(), chain_utils.SerializeLocation(latestLocation))
//
//}
//
//func (mvDB *MultiVersionDB) updateLatestValueId(batch *leveldb.Batch, keyId uint64, valueId uint64) {
//	batch.Put(chain_utils.CreateLatestValueKey(keyId), chain_utils.Uint64ToBytes(valueId))
//}
//
//func (mvDB *MultiVersionDB) deleteValue(batch *leveldb.Batch, valueId uint64) {
//	batch.Delete(chain_utils.CreateValueIdKey(valueId))
//}
//
//func (mvDB *MultiVersionDB) deleteValueId(batch *leveldb.Batch, keyId uint64) {
//	batch.Delete(chain_utils.CreateLatestValueKey(keyId))
//}
