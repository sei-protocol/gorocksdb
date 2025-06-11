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
//
// The timestamp slice is treated as opaque; its length must match the
// TimestampSize configured for the column family (e.g., 8 for uint64 epoch).
// The slice is not retained by RocksDB after the call returns.
func (db *DB) PutCFWithTS(
	wo *WriteOptions,
	cf *ColumnFamilyHandle,
	key, ts, value []byte,
) error {
	var cKey *C.char
	if len(key) > 0 {
		cKey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	var cTs *C.char
	if len(ts) > 0 {
		cTs = (*C.char)(unsafe.Pointer(&ts[0]))
	}
	var cVal *C.char
	if len(value) > 0 {
		cVal = (*C.char)(unsafe.Pointer(&value[0]))
	}

	cErr := C.gorocksdb_put_cf_with_ts(
		db.c, wo.c, cf.c,
		cKey, C.size_t(len(key)),
		cTs, C.size_t(len(ts)),
		cVal, C.size_t(len(value)))
	return convertErr(cErr)
}

// PutWithTS is the default‑CF convenience wrapper around PutCFWithTS.
func (db *DB) PutWithTS(
	wo *WriteOptions,
	key, ts, value []byte,
) error {
	return db.PutCFWithTS(wo, db.defaultCF, key, ts, value)
}

// IncreaseFullHistoryTsLow raises RocksDB’s low‑water mark for full‑history
// retention. After the call, any keys whose most recent timestamp is below
// the supplied value may be dropped during compaction. The timestamp must be
// monotonically increasing between calls per column family.
func (db *DB) IncreaseFullHistoryTsLow(
	cf *ColumnFamilyHandle,
	ts []byte,
) error {
	var cTs *C.char
	if len(ts) > 0 {
		cTs = (*C.char)(unsafe.Pointer(&ts[0]))
	}

	cErr := C.gorocksdb_increase_full_history_ts_low(
		db.c, cf.c,
		cTs, C.size_t(len(ts)))
	return convertErr(cErr)
}
