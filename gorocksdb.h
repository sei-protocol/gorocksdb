#include <stdlib.h>
#include "rocksdb/c.h"

// This API provides convenient C wrapper functions for rocksdb client.

/* Base */

extern void gorocksdb_destruct_handler(void* state);

/* CompactionFilter */

extern rocksdb_compactionfilter_t* gorocksdb_compactionfilter_create(uintptr_t idx);

/* Comparator */

extern rocksdb_comparator_t* gorocksdb_comparator_create(uintptr_t idx);

/* Filter Policy */

extern rocksdb_filterpolicy_t* gorocksdb_filterpolicy_create(uintptr_t idx);
extern void gorocksdb_filterpolicy_delete_filter(void* state, const char* v, size_t s);

/* Merge Operator */

extern rocksdb_mergeoperator_t* gorocksdb_mergeoperator_create(uintptr_t idx);
extern void gorocksdb_mergeoperator_delete_value(void* state, const char* v, size_t s);

/* Slice Transform */

extern rocksdb_slicetransform_t* gorocksdb_slicetransform_create(uintptr_t idx);

/* gorocksdb.h */
extern rocksdb_status_t* gorocksdb_put_cf_with_ts(
    rocksdb_t* db,
    const rocksdb_writeoptions_t* options,
    rocksdb_column_family_handle_t* cf,
    const char* key, size_t keylen,
    const char* ts,  size_t tslen,
    const char* val, size_t vallen);

extern rocksdb_status_t* gorocksdb_increase_full_history_ts_low(
    rocksdb_t* db,
    rocksdb_column_family_handle_t* cf,
    const char* ts, size_t tslen);

extern rocksdb_status_t* gorocksdb_delete_cf_with_ts(
    rocksdb_t*, const rocksdb_writeoptions_t*,
    rocksdb_column_family_handle_t*,
    const char* key, size_t keylen,
    const char* ts,  size_t tslen);

