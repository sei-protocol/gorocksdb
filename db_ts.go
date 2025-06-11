//go:build rocksdb_user_timestamp
// +build rocksdb_user_timestamp

package gorocksdb

/*
#cgo CFLAGS: -DUSE_USER_DEFINED_TIMESTAMP
#include "rocksdb/c.h"
#include "gorocksdb.h"
*/
import "C"
import (
	"unsafe"
)

// PutCFWithTS stores the given key/value pair in the specified column family
// while attaching an explicit timestamp (ts).
func (db *DB) PutCFWithTS(
	wo *WriteOptions,
	cf *ColumnFamilyHandle,
	key, ts, value []byte,
) error {
	cErr := C.gorocksdb_put_cf_with_ts(
		db.c, wo.c, cf.c,
		byteSliceToChar(key), C.size_t(len(key)),
		byteSliceToChar(ts), C.size_t(len(ts)),
		byteSliceToChar(value), C.size_t(len(value)))
	return convertErr(cErr)
}

// PutWithTS is the default‑CF convenience wrapper.
func (db *DB) PutWithTS(wo *WriteOptions, key, ts, value []byte) error {
	return db.PutCFWithTS(wo, db.defaultCF, key, ts, value)
}

// DeleteCFWithTS marks the given key as deleted with the supplied timestamp.
// The semantics mirror PutCFWithTS but create a tombstone instead of a value.
func (db *DB) DeleteCFWithTS(
	wo *WriteOptions,
	cf *ColumnFamilyHandle,
	key, ts []byte,
) error {
	cErr := C.gorocksdb_delete_cf_with_ts(
		db.c, wo.c, cf.c,
		byteSliceToChar(key), C.size_t(len(key)),
		byteSliceToChar(ts), C.size_t(len(ts)))
	return convertErr(cErr)
}

// DeleteWithTS is the default‑CF convenience wrapper.
func (db *DB) DeleteWithTS(wo *WriteOptions, key, ts []byte) error {
	return db.DeleteCFWithTS(wo, db.defaultCF, key, ts)
}

// IncreaseFullHistoryTsLow raises RocksDB’s low‑water mark for full‑history
// retention for the specified column family.
func (db *DB) IncreaseFullHistoryTsLow(cf *ColumnFamilyHandle, ts []byte) error {
	cErr := C.gorocksdb_increase_full_history_ts_low(
		db.c, cf.c,
		byteSliceToChar(ts), C.size_t(len(ts)))
	return convertErr(cErr)
}

// byteSliceToChar safely converts a nil‑able Go []byte to *C.char.
func byteSliceToChar(b []byte) *C.char {
	if len(b) == 0 {
		return nil
	}
	return (*C.char)(unsafe.Pointer(&b[0]))
}
