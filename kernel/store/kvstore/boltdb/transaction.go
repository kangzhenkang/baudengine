package boltdb

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
	"github.com/tiglabs/baud/kernel/store/kvstore"
)

type Transaction struct {
	tx *bolt.Tx
	bucket *bolt.Bucket
	writable bool
}

func(bs *Store) NewTransaction(writable bool) (kvstore.Transaction, error) {
	tx, err := bs.db.Begin(false)
	if err != nil {
		return nil, err
	}
	bucket := tx.Bucket([]byte(bs.bucket))
	return &Transaction{tx: tx, bucket: bucket, writable: writable}, nil
}

func(tx *Transaction) Put(key, value []byte, ops ...*kvstore.Option) error {
	err := tx.bucket.Put(key, value)
	if err != nil {
		return err
	}
	if len(ops) > 0 {
		var buff [8]byte
		r := tx.tx.Bucket(raftBucket)
		binary.BigEndian.PutUint64(buff[:], ops[0].ApplyID)
		err = r.Put(RAFT_APPLY_ID, buff[:])
		if err != nil {
			return err
		}

	}
	return nil
}

func(tx *Transaction) Get(key []byte) ([]byte, error) {
	v := tx.bucket.Get(key)
	if v != nil {
		return cloneBytes(v), nil
	}
	return nil, nil
}

func(tx *Transaction) Delete(key []byte, ops ...*kvstore.Option) error {
	err := tx.bucket.Delete(key)
	if err != nil {
		return err
	}
	if len(ops) > 0 {
		var buff [8]byte
		r := tx.tx.Bucket(raftBucket)
		binary.BigEndian.PutUint64(buff[:], ops[0].ApplyID)
		err = r.Put(RAFT_APPLY_ID, buff[:])
		if err != nil {
			return err
		}

	}
	return nil
}

func(tx *Transaction) PrefixIterator(prefix []byte) kvstore.KVIterator {
	cursor := tx.bucket.Cursor()

	rv := &Iterator{
		cursor: cursor,
		prefix: prefix,
	}

	rv.Seek(prefix)
	return rv
}

func(tx *Transaction) RangeIterator(start, end []byte) kvstore.KVIterator {
	cursor := tx.bucket.Cursor()
	rv := &Iterator{
		tx:     tx,
		cursor: cursor,
		start:  start,
		end:    end,
	}

	rv.Seek(start)
	return rv
}

func(tx *Transaction) Commit() error {
	if tx == nil {
		return nil
	}
	if tx.writable {
		return tx.tx.Commit()
	} else {
		return tx.Rollback()
	}
}

func(tx *Transaction) Rollback() error {
	if tx == nil {
		return nil
	}
	return tx.Rollback()
}