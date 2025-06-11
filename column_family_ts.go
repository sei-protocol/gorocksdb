// column_family_ts.go
//
// Build only when RocksDB was compiled with USE_USER_DEFINED_TIMESTAMP=1.
//go:build rocksdb_user_timestamp
// +build rocksdb_user_timestamp

package gorocksdb

/*
#cgo CFLAGS: -DUSE_USER_DEFINED_TIMESTAMP
#include "rocksdb/c.h"
*/
import "C"

// EnableUserDefinedTimestamp turns the feature on/off for this column‑family
// *before* the CF is created. It maps straight to
// rocksdb_column_family_options_enable_user_defined_timestamp().
func (opts *ColumnFamilyOptions) EnableUserDefinedTimestamp(enable bool) {
	var flag C.uchar
	if enable {
		flag = 1
	}
	C.rocksdb_column_family_options_enable_user_defined_timestamp(opts.c, flag)
}

// SetTimestampSize sets the fixed length (in bytes) of the timestamp slice
// RocksDB will expect. Most users pick 8 for uint64 epoch microseconds.
func (opts *ColumnFamilyOptions) SetTimestampSize(size int) {
	C.rocksdb_column_family_options_set_timestamp_size(opts.c, C.size_t(size))
}
