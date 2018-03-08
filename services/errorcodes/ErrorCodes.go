// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package errorcodes

import (
	"bytes"
	"reflect"
	"database/sql/driver"
	"errors"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type TErrorCode int64
const (
  TErrorCode_OK TErrorCode = 0
  TErrorCode_UNUSED TErrorCode = 1
  TErrorCode_GENERAL TErrorCode = 2
  TErrorCode_CANCELLED TErrorCode = 3
  TErrorCode_ANALYSIS_ERROR TErrorCode = 4
  TErrorCode_NOT_IMPLEMENTED_ERROR TErrorCode = 5
  TErrorCode_RUNTIME_ERROR TErrorCode = 6
  TErrorCode_MEM_LIMIT_EXCEEDED TErrorCode = 7
  TErrorCode_INTERNAL_ERROR TErrorCode = 8
  TErrorCode_RECOVERABLE_ERROR TErrorCode = 9
  TErrorCode_PARQUET_MULTIPLE_BLOCKS TErrorCode = 10
  TErrorCode_PARQUET_COLUMN_METADATA_INVALID TErrorCode = 11
  TErrorCode_PARQUET_HEADER_PAGE_SIZE_EXCEEDED TErrorCode = 12
  TErrorCode_PARQUET_HEADER_EOF TErrorCode = 13
  TErrorCode_PARQUET_GROUP_ROW_COUNT_ERROR TErrorCode = 14
  TErrorCode_PARQUET_GROUP_ROW_COUNT_OVERFLOW TErrorCode = 15
  TErrorCode_PARQUET_MISSING_PRECISION TErrorCode = 16
  TErrorCode_PARQUET_WRONG_PRECISION TErrorCode = 17
  TErrorCode_PARQUET_BAD_CONVERTED_TYPE TErrorCode = 18
  TErrorCode_PARQUET_INCOMPATIBLE_DECIMAL TErrorCode = 19
  TErrorCode_SEQUENCE_SCANNER_PARSE_ERROR TErrorCode = 20
  TErrorCode_SNAPPY_DECOMPRESS_INVALID_BLOCK_SIZE TErrorCode = 21
  TErrorCode_SNAPPY_DECOMPRESS_INVALID_COMPRESSED_LENGTH TErrorCode = 22
  TErrorCode_SNAPPY_DECOMPRESS_UNCOMPRESSED_LENGTH_FAILED TErrorCode = 23
  TErrorCode_SNAPPY_DECOMPRESS_RAW_UNCOMPRESS_FAILED TErrorCode = 24
  TErrorCode_SNAPPY_DECOMPRESS_DECOMPRESS_SIZE_INCORRECT TErrorCode = 25
  TErrorCode_FRAGMENT_EXECUTOR TErrorCode = 26
  TErrorCode_PARTITIONED_HASH_JOIN_MAX_PARTITION_DEPTH TErrorCode = 27
  TErrorCode_PARTITIONED_AGG_MAX_PARTITION_DEPTH TErrorCode = 28
  TErrorCode_MISSING_BUILTIN TErrorCode = 29
  TErrorCode_RPC_GENERAL_ERROR TErrorCode = 30
  TErrorCode_RPC_RECV_TIMEOUT TErrorCode = 31
  TErrorCode_UDF_VERIFY_FAILED TErrorCode = 32
  TErrorCode_PARQUET_CORRUPT_RLE_BYTES TErrorCode = 33
  TErrorCode_AVRO_DECIMAL_RESOLUTION_ERROR TErrorCode = 34
  TErrorCode_AVRO_DECIMAL_METADATA_MISMATCH TErrorCode = 35
  TErrorCode_AVRO_SCHEMA_RESOLUTION_ERROR TErrorCode = 36
  TErrorCode_AVRO_SCHEMA_METADATA_MISMATCH TErrorCode = 37
  TErrorCode_AVRO_UNSUPPORTED_DEFAULT_VALUE TErrorCode = 38
  TErrorCode_AVRO_MISSING_FIELD TErrorCode = 39
  TErrorCode_AVRO_MISSING_DEFAULT TErrorCode = 40
  TErrorCode_AVRO_NULLABILITY_MISMATCH TErrorCode = 41
  TErrorCode_AVRO_NOT_A_RECORD TErrorCode = 42
  TErrorCode_PARQUET_DEF_LEVEL_ERROR TErrorCode = 43
  TErrorCode_PARQUET_NUM_COL_VALS_ERROR TErrorCode = 44
  TErrorCode_PARQUET_DICT_DECODE_FAILURE TErrorCode = 45
  TErrorCode_SSL_PASSWORD_CMD_FAILED TErrorCode = 46
  TErrorCode_SSL_CERTIFICATE_PATH_BLANK TErrorCode = 47
  TErrorCode_SSL_PRIVATE_KEY_PATH_BLANK TErrorCode = 48
  TErrorCode_SSL_CERTIFICATE_NOT_FOUND TErrorCode = 49
  TErrorCode_SSL_PRIVATE_KEY_NOT_FOUND TErrorCode = 50
  TErrorCode_SSL_SOCKET_CREATION_FAILED TErrorCode = 51
  TErrorCode_MEM_ALLOC_FAILED TErrorCode = 52
  TErrorCode_PARQUET_REP_LEVEL_ERROR TErrorCode = 53
  TErrorCode_PARQUET_UNRECOGNIZED_SCHEMA TErrorCode = 54
  TErrorCode_COLLECTION_ALLOC_FAILED TErrorCode = 55
  TErrorCode_TMP_DEVICE_BLACKLISTED TErrorCode = 56
  TErrorCode_TMP_FILE_BLACKLISTED TErrorCode = 57
  TErrorCode_RPC_CLIENT_CONNECT_FAILURE TErrorCode = 58
  TErrorCode_STALE_METADATA_FILE_TOO_SHORT TErrorCode = 59
  TErrorCode_PARQUET_BAD_VERSION_NUMBER TErrorCode = 60
  TErrorCode_SCANNER_INCOMPLETE_READ TErrorCode = 61
  TErrorCode_SCANNER_INVALID_READ TErrorCode = 62
  TErrorCode_AVRO_BAD_VERSION_HEADER TErrorCode = 63
  TErrorCode_UDF_MEM_LIMIT_EXCEEDED TErrorCode = 64
  TErrorCode_UNUSED_65 TErrorCode = 65
  TErrorCode_COMPRESSED_FILE_MULTIPLE_BLOCKS TErrorCode = 66
  TErrorCode_COMPRESSED_FILE_BLOCK_CORRUPTED TErrorCode = 67
  TErrorCode_COMPRESSED_FILE_DECOMPRESSOR_ERROR TErrorCode = 68
  TErrorCode_COMPRESSED_FILE_DECOMPRESSOR_NO_PROGRESS TErrorCode = 69
  TErrorCode_COMPRESSED_FILE_TRUNCATED TErrorCode = 70
  TErrorCode_DATASTREAM_SENDER_TIMEOUT TErrorCode = 71
  TErrorCode_KUDU_IMPALA_TYPE_MISSING TErrorCode = 72
  TErrorCode_IMPALA_KUDU_TYPE_MISSING TErrorCode = 73
  TErrorCode_KUDU_NOT_SUPPORTED_ON_OS TErrorCode = 74
  TErrorCode_KUDU_NOT_ENABLED TErrorCode = 75
  TErrorCode_PARTITIONED_HASH_JOIN_REPARTITION_FAILS TErrorCode = 76
  TErrorCode_UNUSED_77 TErrorCode = 77
  TErrorCode_AVRO_TRUNCATED_BLOCK TErrorCode = 78
  TErrorCode_AVRO_INVALID_UNION TErrorCode = 79
  TErrorCode_AVRO_INVALID_BOOLEAN TErrorCode = 80
  TErrorCode_AVRO_INVALID_LENGTH TErrorCode = 81
  TErrorCode_SCANNER_INVALID_INT TErrorCode = 82
  TErrorCode_AVRO_INVALID_RECORD_COUNT TErrorCode = 83
  TErrorCode_AVRO_INVALID_COMPRESSED_SIZE TErrorCode = 84
  TErrorCode_AVRO_INVALID_METADATA_COUNT TErrorCode = 85
  TErrorCode_SCANNER_STRING_LENGTH_OVERFLOW TErrorCode = 86
  TErrorCode_PARQUET_CORRUPT_PLAIN_VALUE TErrorCode = 87
  TErrorCode_PARQUET_CORRUPT_DICTIONARY TErrorCode = 88
  TErrorCode_TEXT_PARSER_TRUNCATED_COLUMN TErrorCode = 89
  TErrorCode_SCRATCH_LIMIT_EXCEEDED TErrorCode = 90
  TErrorCode_BUFFER_ALLOCATION_FAILED TErrorCode = 91
  TErrorCode_PARQUET_ZERO_ROWS_IN_NON_EMPTY_FILE TErrorCode = 92
  TErrorCode_NO_REGISTERED_BACKENDS TErrorCode = 93
  TErrorCode_KUDU_KEY_ALREADY_PRESENT TErrorCode = 94
  TErrorCode_KUDU_NOT_FOUND TErrorCode = 95
  TErrorCode_KUDU_SESSION_ERROR TErrorCode = 96
  TErrorCode_AVRO_UNSUPPORTED_TYPE TErrorCode = 97
  TErrorCode_AVRO_INVALID_DECIMAL TErrorCode = 98
  TErrorCode_KUDU_NULL_CONSTRAINT_VIOLATION TErrorCode = 99
  TErrorCode_PARQUET_TIMESTAMP_OUT_OF_RANGE TErrorCode = 100
  TErrorCode_SCRATCH_ALLOCATION_FAILED TErrorCode = 101
  TErrorCode_SCRATCH_READ_TRUNCATED TErrorCode = 102
  TErrorCode_KUDU_TIMESTAMP_OUT_OF_RANGE TErrorCode = 103
  TErrorCode_MAX_ROW_SIZE TErrorCode = 104
  TErrorCode_IR_VERIFY_FAILED TErrorCode = 105
  TErrorCode_MINIMUM_RESERVATION_UNAVAILABLE TErrorCode = 106
  TErrorCode_ADMISSION_REJECTED TErrorCode = 107
  TErrorCode_ADMISSION_TIMED_OUT TErrorCode = 108
  TErrorCode_THREAD_CREATION_FAILED TErrorCode = 109
  TErrorCode_DISK_IO_ERROR TErrorCode = 110
  TErrorCode_DATASTREAM_RECVR_CLOSED TErrorCode = 111
  TErrorCode_BAD_PRINCIPAL_FORMAT TErrorCode = 112
  TErrorCode_LZ4_COMPRESSION_INPUT_TOO_LARGE TErrorCode = 113
  TErrorCode_SASL_APP_NAME_MISMATCH TErrorCode = 114
)

func (p TErrorCode) String() string {
  switch p {
  case TErrorCode_OK: return "OK"
  case TErrorCode_UNUSED: return "UNUSED"
  case TErrorCode_GENERAL: return "GENERAL"
  case TErrorCode_CANCELLED: return "CANCELLED"
  case TErrorCode_ANALYSIS_ERROR: return "ANALYSIS_ERROR"
  case TErrorCode_NOT_IMPLEMENTED_ERROR: return "NOT_IMPLEMENTED_ERROR"
  case TErrorCode_RUNTIME_ERROR: return "RUNTIME_ERROR"
  case TErrorCode_MEM_LIMIT_EXCEEDED: return "MEM_LIMIT_EXCEEDED"
  case TErrorCode_INTERNAL_ERROR: return "INTERNAL_ERROR"
  case TErrorCode_RECOVERABLE_ERROR: return "RECOVERABLE_ERROR"
  case TErrorCode_PARQUET_MULTIPLE_BLOCKS: return "PARQUET_MULTIPLE_BLOCKS"
  case TErrorCode_PARQUET_COLUMN_METADATA_INVALID: return "PARQUET_COLUMN_METADATA_INVALID"
  case TErrorCode_PARQUET_HEADER_PAGE_SIZE_EXCEEDED: return "PARQUET_HEADER_PAGE_SIZE_EXCEEDED"
  case TErrorCode_PARQUET_HEADER_EOF: return "PARQUET_HEADER_EOF"
  case TErrorCode_PARQUET_GROUP_ROW_COUNT_ERROR: return "PARQUET_GROUP_ROW_COUNT_ERROR"
  case TErrorCode_PARQUET_GROUP_ROW_COUNT_OVERFLOW: return "PARQUET_GROUP_ROW_COUNT_OVERFLOW"
  case TErrorCode_PARQUET_MISSING_PRECISION: return "PARQUET_MISSING_PRECISION"
  case TErrorCode_PARQUET_WRONG_PRECISION: return "PARQUET_WRONG_PRECISION"
  case TErrorCode_PARQUET_BAD_CONVERTED_TYPE: return "PARQUET_BAD_CONVERTED_TYPE"
  case TErrorCode_PARQUET_INCOMPATIBLE_DECIMAL: return "PARQUET_INCOMPATIBLE_DECIMAL"
  case TErrorCode_SEQUENCE_SCANNER_PARSE_ERROR: return "SEQUENCE_SCANNER_PARSE_ERROR"
  case TErrorCode_SNAPPY_DECOMPRESS_INVALID_BLOCK_SIZE: return "SNAPPY_DECOMPRESS_INVALID_BLOCK_SIZE"
  case TErrorCode_SNAPPY_DECOMPRESS_INVALID_COMPRESSED_LENGTH: return "SNAPPY_DECOMPRESS_INVALID_COMPRESSED_LENGTH"
  case TErrorCode_SNAPPY_DECOMPRESS_UNCOMPRESSED_LENGTH_FAILED: return "SNAPPY_DECOMPRESS_UNCOMPRESSED_LENGTH_FAILED"
  case TErrorCode_SNAPPY_DECOMPRESS_RAW_UNCOMPRESS_FAILED: return "SNAPPY_DECOMPRESS_RAW_UNCOMPRESS_FAILED"
  case TErrorCode_SNAPPY_DECOMPRESS_DECOMPRESS_SIZE_INCORRECT: return "SNAPPY_DECOMPRESS_DECOMPRESS_SIZE_INCORRECT"
  case TErrorCode_FRAGMENT_EXECUTOR: return "FRAGMENT_EXECUTOR"
  case TErrorCode_PARTITIONED_HASH_JOIN_MAX_PARTITION_DEPTH: return "PARTITIONED_HASH_JOIN_MAX_PARTITION_DEPTH"
  case TErrorCode_PARTITIONED_AGG_MAX_PARTITION_DEPTH: return "PARTITIONED_AGG_MAX_PARTITION_DEPTH"
  case TErrorCode_MISSING_BUILTIN: return "MISSING_BUILTIN"
  case TErrorCode_RPC_GENERAL_ERROR: return "RPC_GENERAL_ERROR"
  case TErrorCode_RPC_RECV_TIMEOUT: return "RPC_RECV_TIMEOUT"
  case TErrorCode_UDF_VERIFY_FAILED: return "UDF_VERIFY_FAILED"
  case TErrorCode_PARQUET_CORRUPT_RLE_BYTES: return "PARQUET_CORRUPT_RLE_BYTES"
  case TErrorCode_AVRO_DECIMAL_RESOLUTION_ERROR: return "AVRO_DECIMAL_RESOLUTION_ERROR"
  case TErrorCode_AVRO_DECIMAL_METADATA_MISMATCH: return "AVRO_DECIMAL_METADATA_MISMATCH"
  case TErrorCode_AVRO_SCHEMA_RESOLUTION_ERROR: return "AVRO_SCHEMA_RESOLUTION_ERROR"
  case TErrorCode_AVRO_SCHEMA_METADATA_MISMATCH: return "AVRO_SCHEMA_METADATA_MISMATCH"
  case TErrorCode_AVRO_UNSUPPORTED_DEFAULT_VALUE: return "AVRO_UNSUPPORTED_DEFAULT_VALUE"
  case TErrorCode_AVRO_MISSING_FIELD: return "AVRO_MISSING_FIELD"
  case TErrorCode_AVRO_MISSING_DEFAULT: return "AVRO_MISSING_DEFAULT"
  case TErrorCode_AVRO_NULLABILITY_MISMATCH: return "AVRO_NULLABILITY_MISMATCH"
  case TErrorCode_AVRO_NOT_A_RECORD: return "AVRO_NOT_A_RECORD"
  case TErrorCode_PARQUET_DEF_LEVEL_ERROR: return "PARQUET_DEF_LEVEL_ERROR"
  case TErrorCode_PARQUET_NUM_COL_VALS_ERROR: return "PARQUET_NUM_COL_VALS_ERROR"
  case TErrorCode_PARQUET_DICT_DECODE_FAILURE: return "PARQUET_DICT_DECODE_FAILURE"
  case TErrorCode_SSL_PASSWORD_CMD_FAILED: return "SSL_PASSWORD_CMD_FAILED"
  case TErrorCode_SSL_CERTIFICATE_PATH_BLANK: return "SSL_CERTIFICATE_PATH_BLANK"
  case TErrorCode_SSL_PRIVATE_KEY_PATH_BLANK: return "SSL_PRIVATE_KEY_PATH_BLANK"
  case TErrorCode_SSL_CERTIFICATE_NOT_FOUND: return "SSL_CERTIFICATE_NOT_FOUND"
  case TErrorCode_SSL_PRIVATE_KEY_NOT_FOUND: return "SSL_PRIVATE_KEY_NOT_FOUND"
  case TErrorCode_SSL_SOCKET_CREATION_FAILED: return "SSL_SOCKET_CREATION_FAILED"
  case TErrorCode_MEM_ALLOC_FAILED: return "MEM_ALLOC_FAILED"
  case TErrorCode_PARQUET_REP_LEVEL_ERROR: return "PARQUET_REP_LEVEL_ERROR"
  case TErrorCode_PARQUET_UNRECOGNIZED_SCHEMA: return "PARQUET_UNRECOGNIZED_SCHEMA"
  case TErrorCode_COLLECTION_ALLOC_FAILED: return "COLLECTION_ALLOC_FAILED"
  case TErrorCode_TMP_DEVICE_BLACKLISTED: return "TMP_DEVICE_BLACKLISTED"
  case TErrorCode_TMP_FILE_BLACKLISTED: return "TMP_FILE_BLACKLISTED"
  case TErrorCode_RPC_CLIENT_CONNECT_FAILURE: return "RPC_CLIENT_CONNECT_FAILURE"
  case TErrorCode_STALE_METADATA_FILE_TOO_SHORT: return "STALE_METADATA_FILE_TOO_SHORT"
  case TErrorCode_PARQUET_BAD_VERSION_NUMBER: return "PARQUET_BAD_VERSION_NUMBER"
  case TErrorCode_SCANNER_INCOMPLETE_READ: return "SCANNER_INCOMPLETE_READ"
  case TErrorCode_SCANNER_INVALID_READ: return "SCANNER_INVALID_READ"
  case TErrorCode_AVRO_BAD_VERSION_HEADER: return "AVRO_BAD_VERSION_HEADER"
  case TErrorCode_UDF_MEM_LIMIT_EXCEEDED: return "UDF_MEM_LIMIT_EXCEEDED"
  case TErrorCode_UNUSED_65: return "UNUSED_65"
  case TErrorCode_COMPRESSED_FILE_MULTIPLE_BLOCKS: return "COMPRESSED_FILE_MULTIPLE_BLOCKS"
  case TErrorCode_COMPRESSED_FILE_BLOCK_CORRUPTED: return "COMPRESSED_FILE_BLOCK_CORRUPTED"
  case TErrorCode_COMPRESSED_FILE_DECOMPRESSOR_ERROR: return "COMPRESSED_FILE_DECOMPRESSOR_ERROR"
  case TErrorCode_COMPRESSED_FILE_DECOMPRESSOR_NO_PROGRESS: return "COMPRESSED_FILE_DECOMPRESSOR_NO_PROGRESS"
  case TErrorCode_COMPRESSED_FILE_TRUNCATED: return "COMPRESSED_FILE_TRUNCATED"
  case TErrorCode_DATASTREAM_SENDER_TIMEOUT: return "DATASTREAM_SENDER_TIMEOUT"
  case TErrorCode_KUDU_IMPALA_TYPE_MISSING: return "KUDU_IMPALA_TYPE_MISSING"
  case TErrorCode_IMPALA_KUDU_TYPE_MISSING: return "IMPALA_KUDU_TYPE_MISSING"
  case TErrorCode_KUDU_NOT_SUPPORTED_ON_OS: return "KUDU_NOT_SUPPORTED_ON_OS"
  case TErrorCode_KUDU_NOT_ENABLED: return "KUDU_NOT_ENABLED"
  case TErrorCode_PARTITIONED_HASH_JOIN_REPARTITION_FAILS: return "PARTITIONED_HASH_JOIN_REPARTITION_FAILS"
  case TErrorCode_UNUSED_77: return "UNUSED_77"
  case TErrorCode_AVRO_TRUNCATED_BLOCK: return "AVRO_TRUNCATED_BLOCK"
  case TErrorCode_AVRO_INVALID_UNION: return "AVRO_INVALID_UNION"
  case TErrorCode_AVRO_INVALID_BOOLEAN: return "AVRO_INVALID_BOOLEAN"
  case TErrorCode_AVRO_INVALID_LENGTH: return "AVRO_INVALID_LENGTH"
  case TErrorCode_SCANNER_INVALID_INT: return "SCANNER_INVALID_INT"
  case TErrorCode_AVRO_INVALID_RECORD_COUNT: return "AVRO_INVALID_RECORD_COUNT"
  case TErrorCode_AVRO_INVALID_COMPRESSED_SIZE: return "AVRO_INVALID_COMPRESSED_SIZE"
  case TErrorCode_AVRO_INVALID_METADATA_COUNT: return "AVRO_INVALID_METADATA_COUNT"
  case TErrorCode_SCANNER_STRING_LENGTH_OVERFLOW: return "SCANNER_STRING_LENGTH_OVERFLOW"
  case TErrorCode_PARQUET_CORRUPT_PLAIN_VALUE: return "PARQUET_CORRUPT_PLAIN_VALUE"
  case TErrorCode_PARQUET_CORRUPT_DICTIONARY: return "PARQUET_CORRUPT_DICTIONARY"
  case TErrorCode_TEXT_PARSER_TRUNCATED_COLUMN: return "TEXT_PARSER_TRUNCATED_COLUMN"
  case TErrorCode_SCRATCH_LIMIT_EXCEEDED: return "SCRATCH_LIMIT_EXCEEDED"
  case TErrorCode_BUFFER_ALLOCATION_FAILED: return "BUFFER_ALLOCATION_FAILED"
  case TErrorCode_PARQUET_ZERO_ROWS_IN_NON_EMPTY_FILE: return "PARQUET_ZERO_ROWS_IN_NON_EMPTY_FILE"
  case TErrorCode_NO_REGISTERED_BACKENDS: return "NO_REGISTERED_BACKENDS"
  case TErrorCode_KUDU_KEY_ALREADY_PRESENT: return "KUDU_KEY_ALREADY_PRESENT"
  case TErrorCode_KUDU_NOT_FOUND: return "KUDU_NOT_FOUND"
  case TErrorCode_KUDU_SESSION_ERROR: return "KUDU_SESSION_ERROR"
  case TErrorCode_AVRO_UNSUPPORTED_TYPE: return "AVRO_UNSUPPORTED_TYPE"
  case TErrorCode_AVRO_INVALID_DECIMAL: return "AVRO_INVALID_DECIMAL"
  case TErrorCode_KUDU_NULL_CONSTRAINT_VIOLATION: return "KUDU_NULL_CONSTRAINT_VIOLATION"
  case TErrorCode_PARQUET_TIMESTAMP_OUT_OF_RANGE: return "PARQUET_TIMESTAMP_OUT_OF_RANGE"
  case TErrorCode_SCRATCH_ALLOCATION_FAILED: return "SCRATCH_ALLOCATION_FAILED"
  case TErrorCode_SCRATCH_READ_TRUNCATED: return "SCRATCH_READ_TRUNCATED"
  case TErrorCode_KUDU_TIMESTAMP_OUT_OF_RANGE: return "KUDU_TIMESTAMP_OUT_OF_RANGE"
  case TErrorCode_MAX_ROW_SIZE: return "MAX_ROW_SIZE"
  case TErrorCode_IR_VERIFY_FAILED: return "IR_VERIFY_FAILED"
  case TErrorCode_MINIMUM_RESERVATION_UNAVAILABLE: return "MINIMUM_RESERVATION_UNAVAILABLE"
  case TErrorCode_ADMISSION_REJECTED: return "ADMISSION_REJECTED"
  case TErrorCode_ADMISSION_TIMED_OUT: return "ADMISSION_TIMED_OUT"
  case TErrorCode_THREAD_CREATION_FAILED: return "THREAD_CREATION_FAILED"
  case TErrorCode_DISK_IO_ERROR: return "DISK_IO_ERROR"
  case TErrorCode_DATASTREAM_RECVR_CLOSED: return "DATASTREAM_RECVR_CLOSED"
  case TErrorCode_BAD_PRINCIPAL_FORMAT: return "BAD_PRINCIPAL_FORMAT"
  case TErrorCode_LZ4_COMPRESSION_INPUT_TOO_LARGE: return "LZ4_COMPRESSION_INPUT_TOO_LARGE"
  case TErrorCode_SASL_APP_NAME_MISMATCH: return "SASL_APP_NAME_MISMATCH"
  }
  return "<UNSET>"
}

func TErrorCodeFromString(s string) (TErrorCode, error) {
  switch s {
  case "OK": return TErrorCode_OK, nil 
  case "UNUSED": return TErrorCode_UNUSED, nil 
  case "GENERAL": return TErrorCode_GENERAL, nil 
  case "CANCELLED": return TErrorCode_CANCELLED, nil 
  case "ANALYSIS_ERROR": return TErrorCode_ANALYSIS_ERROR, nil 
  case "NOT_IMPLEMENTED_ERROR": return TErrorCode_NOT_IMPLEMENTED_ERROR, nil 
  case "RUNTIME_ERROR": return TErrorCode_RUNTIME_ERROR, nil 
  case "MEM_LIMIT_EXCEEDED": return TErrorCode_MEM_LIMIT_EXCEEDED, nil 
  case "INTERNAL_ERROR": return TErrorCode_INTERNAL_ERROR, nil 
  case "RECOVERABLE_ERROR": return TErrorCode_RECOVERABLE_ERROR, nil 
  case "PARQUET_MULTIPLE_BLOCKS": return TErrorCode_PARQUET_MULTIPLE_BLOCKS, nil 
  case "PARQUET_COLUMN_METADATA_INVALID": return TErrorCode_PARQUET_COLUMN_METADATA_INVALID, nil 
  case "PARQUET_HEADER_PAGE_SIZE_EXCEEDED": return TErrorCode_PARQUET_HEADER_PAGE_SIZE_EXCEEDED, nil 
  case "PARQUET_HEADER_EOF": return TErrorCode_PARQUET_HEADER_EOF, nil 
  case "PARQUET_GROUP_ROW_COUNT_ERROR": return TErrorCode_PARQUET_GROUP_ROW_COUNT_ERROR, nil 
  case "PARQUET_GROUP_ROW_COUNT_OVERFLOW": return TErrorCode_PARQUET_GROUP_ROW_COUNT_OVERFLOW, nil 
  case "PARQUET_MISSING_PRECISION": return TErrorCode_PARQUET_MISSING_PRECISION, nil 
  case "PARQUET_WRONG_PRECISION": return TErrorCode_PARQUET_WRONG_PRECISION, nil 
  case "PARQUET_BAD_CONVERTED_TYPE": return TErrorCode_PARQUET_BAD_CONVERTED_TYPE, nil 
  case "PARQUET_INCOMPATIBLE_DECIMAL": return TErrorCode_PARQUET_INCOMPATIBLE_DECIMAL, nil 
  case "SEQUENCE_SCANNER_PARSE_ERROR": return TErrorCode_SEQUENCE_SCANNER_PARSE_ERROR, nil 
  case "SNAPPY_DECOMPRESS_INVALID_BLOCK_SIZE": return TErrorCode_SNAPPY_DECOMPRESS_INVALID_BLOCK_SIZE, nil 
  case "SNAPPY_DECOMPRESS_INVALID_COMPRESSED_LENGTH": return TErrorCode_SNAPPY_DECOMPRESS_INVALID_COMPRESSED_LENGTH, nil 
  case "SNAPPY_DECOMPRESS_UNCOMPRESSED_LENGTH_FAILED": return TErrorCode_SNAPPY_DECOMPRESS_UNCOMPRESSED_LENGTH_FAILED, nil 
  case "SNAPPY_DECOMPRESS_RAW_UNCOMPRESS_FAILED": return TErrorCode_SNAPPY_DECOMPRESS_RAW_UNCOMPRESS_FAILED, nil 
  case "SNAPPY_DECOMPRESS_DECOMPRESS_SIZE_INCORRECT": return TErrorCode_SNAPPY_DECOMPRESS_DECOMPRESS_SIZE_INCORRECT, nil 
  case "FRAGMENT_EXECUTOR": return TErrorCode_FRAGMENT_EXECUTOR, nil 
  case "PARTITIONED_HASH_JOIN_MAX_PARTITION_DEPTH": return TErrorCode_PARTITIONED_HASH_JOIN_MAX_PARTITION_DEPTH, nil 
  case "PARTITIONED_AGG_MAX_PARTITION_DEPTH": return TErrorCode_PARTITIONED_AGG_MAX_PARTITION_DEPTH, nil 
  case "MISSING_BUILTIN": return TErrorCode_MISSING_BUILTIN, nil 
  case "RPC_GENERAL_ERROR": return TErrorCode_RPC_GENERAL_ERROR, nil 
  case "RPC_RECV_TIMEOUT": return TErrorCode_RPC_RECV_TIMEOUT, nil 
  case "UDF_VERIFY_FAILED": return TErrorCode_UDF_VERIFY_FAILED, nil 
  case "PARQUET_CORRUPT_RLE_BYTES": return TErrorCode_PARQUET_CORRUPT_RLE_BYTES, nil 
  case "AVRO_DECIMAL_RESOLUTION_ERROR": return TErrorCode_AVRO_DECIMAL_RESOLUTION_ERROR, nil 
  case "AVRO_DECIMAL_METADATA_MISMATCH": return TErrorCode_AVRO_DECIMAL_METADATA_MISMATCH, nil 
  case "AVRO_SCHEMA_RESOLUTION_ERROR": return TErrorCode_AVRO_SCHEMA_RESOLUTION_ERROR, nil 
  case "AVRO_SCHEMA_METADATA_MISMATCH": return TErrorCode_AVRO_SCHEMA_METADATA_MISMATCH, nil 
  case "AVRO_UNSUPPORTED_DEFAULT_VALUE": return TErrorCode_AVRO_UNSUPPORTED_DEFAULT_VALUE, nil 
  case "AVRO_MISSING_FIELD": return TErrorCode_AVRO_MISSING_FIELD, nil 
  case "AVRO_MISSING_DEFAULT": return TErrorCode_AVRO_MISSING_DEFAULT, nil 
  case "AVRO_NULLABILITY_MISMATCH": return TErrorCode_AVRO_NULLABILITY_MISMATCH, nil 
  case "AVRO_NOT_A_RECORD": return TErrorCode_AVRO_NOT_A_RECORD, nil 
  case "PARQUET_DEF_LEVEL_ERROR": return TErrorCode_PARQUET_DEF_LEVEL_ERROR, nil 
  case "PARQUET_NUM_COL_VALS_ERROR": return TErrorCode_PARQUET_NUM_COL_VALS_ERROR, nil 
  case "PARQUET_DICT_DECODE_FAILURE": return TErrorCode_PARQUET_DICT_DECODE_FAILURE, nil 
  case "SSL_PASSWORD_CMD_FAILED": return TErrorCode_SSL_PASSWORD_CMD_FAILED, nil 
  case "SSL_CERTIFICATE_PATH_BLANK": return TErrorCode_SSL_CERTIFICATE_PATH_BLANK, nil 
  case "SSL_PRIVATE_KEY_PATH_BLANK": return TErrorCode_SSL_PRIVATE_KEY_PATH_BLANK, nil 
  case "SSL_CERTIFICATE_NOT_FOUND": return TErrorCode_SSL_CERTIFICATE_NOT_FOUND, nil 
  case "SSL_PRIVATE_KEY_NOT_FOUND": return TErrorCode_SSL_PRIVATE_KEY_NOT_FOUND, nil 
  case "SSL_SOCKET_CREATION_FAILED": return TErrorCode_SSL_SOCKET_CREATION_FAILED, nil 
  case "MEM_ALLOC_FAILED": return TErrorCode_MEM_ALLOC_FAILED, nil 
  case "PARQUET_REP_LEVEL_ERROR": return TErrorCode_PARQUET_REP_LEVEL_ERROR, nil 
  case "PARQUET_UNRECOGNIZED_SCHEMA": return TErrorCode_PARQUET_UNRECOGNIZED_SCHEMA, nil 
  case "COLLECTION_ALLOC_FAILED": return TErrorCode_COLLECTION_ALLOC_FAILED, nil 
  case "TMP_DEVICE_BLACKLISTED": return TErrorCode_TMP_DEVICE_BLACKLISTED, nil 
  case "TMP_FILE_BLACKLISTED": return TErrorCode_TMP_FILE_BLACKLISTED, nil 
  case "RPC_CLIENT_CONNECT_FAILURE": return TErrorCode_RPC_CLIENT_CONNECT_FAILURE, nil 
  case "STALE_METADATA_FILE_TOO_SHORT": return TErrorCode_STALE_METADATA_FILE_TOO_SHORT, nil 
  case "PARQUET_BAD_VERSION_NUMBER": return TErrorCode_PARQUET_BAD_VERSION_NUMBER, nil 
  case "SCANNER_INCOMPLETE_READ": return TErrorCode_SCANNER_INCOMPLETE_READ, nil 
  case "SCANNER_INVALID_READ": return TErrorCode_SCANNER_INVALID_READ, nil 
  case "AVRO_BAD_VERSION_HEADER": return TErrorCode_AVRO_BAD_VERSION_HEADER, nil 
  case "UDF_MEM_LIMIT_EXCEEDED": return TErrorCode_UDF_MEM_LIMIT_EXCEEDED, nil 
  case "UNUSED_65": return TErrorCode_UNUSED_65, nil 
  case "COMPRESSED_FILE_MULTIPLE_BLOCKS": return TErrorCode_COMPRESSED_FILE_MULTIPLE_BLOCKS, nil 
  case "COMPRESSED_FILE_BLOCK_CORRUPTED": return TErrorCode_COMPRESSED_FILE_BLOCK_CORRUPTED, nil 
  case "COMPRESSED_FILE_DECOMPRESSOR_ERROR": return TErrorCode_COMPRESSED_FILE_DECOMPRESSOR_ERROR, nil 
  case "COMPRESSED_FILE_DECOMPRESSOR_NO_PROGRESS": return TErrorCode_COMPRESSED_FILE_DECOMPRESSOR_NO_PROGRESS, nil 
  case "COMPRESSED_FILE_TRUNCATED": return TErrorCode_COMPRESSED_FILE_TRUNCATED, nil 
  case "DATASTREAM_SENDER_TIMEOUT": return TErrorCode_DATASTREAM_SENDER_TIMEOUT, nil 
  case "KUDU_IMPALA_TYPE_MISSING": return TErrorCode_KUDU_IMPALA_TYPE_MISSING, nil 
  case "IMPALA_KUDU_TYPE_MISSING": return TErrorCode_IMPALA_KUDU_TYPE_MISSING, nil 
  case "KUDU_NOT_SUPPORTED_ON_OS": return TErrorCode_KUDU_NOT_SUPPORTED_ON_OS, nil 
  case "KUDU_NOT_ENABLED": return TErrorCode_KUDU_NOT_ENABLED, nil 
  case "PARTITIONED_HASH_JOIN_REPARTITION_FAILS": return TErrorCode_PARTITIONED_HASH_JOIN_REPARTITION_FAILS, nil 
  case "UNUSED_77": return TErrorCode_UNUSED_77, nil 
  case "AVRO_TRUNCATED_BLOCK": return TErrorCode_AVRO_TRUNCATED_BLOCK, nil 
  case "AVRO_INVALID_UNION": return TErrorCode_AVRO_INVALID_UNION, nil 
  case "AVRO_INVALID_BOOLEAN": return TErrorCode_AVRO_INVALID_BOOLEAN, nil 
  case "AVRO_INVALID_LENGTH": return TErrorCode_AVRO_INVALID_LENGTH, nil 
  case "SCANNER_INVALID_INT": return TErrorCode_SCANNER_INVALID_INT, nil 
  case "AVRO_INVALID_RECORD_COUNT": return TErrorCode_AVRO_INVALID_RECORD_COUNT, nil 
  case "AVRO_INVALID_COMPRESSED_SIZE": return TErrorCode_AVRO_INVALID_COMPRESSED_SIZE, nil 
  case "AVRO_INVALID_METADATA_COUNT": return TErrorCode_AVRO_INVALID_METADATA_COUNT, nil 
  case "SCANNER_STRING_LENGTH_OVERFLOW": return TErrorCode_SCANNER_STRING_LENGTH_OVERFLOW, nil 
  case "PARQUET_CORRUPT_PLAIN_VALUE": return TErrorCode_PARQUET_CORRUPT_PLAIN_VALUE, nil 
  case "PARQUET_CORRUPT_DICTIONARY": return TErrorCode_PARQUET_CORRUPT_DICTIONARY, nil 
  case "TEXT_PARSER_TRUNCATED_COLUMN": return TErrorCode_TEXT_PARSER_TRUNCATED_COLUMN, nil 
  case "SCRATCH_LIMIT_EXCEEDED": return TErrorCode_SCRATCH_LIMIT_EXCEEDED, nil 
  case "BUFFER_ALLOCATION_FAILED": return TErrorCode_BUFFER_ALLOCATION_FAILED, nil 
  case "PARQUET_ZERO_ROWS_IN_NON_EMPTY_FILE": return TErrorCode_PARQUET_ZERO_ROWS_IN_NON_EMPTY_FILE, nil 
  case "NO_REGISTERED_BACKENDS": return TErrorCode_NO_REGISTERED_BACKENDS, nil 
  case "KUDU_KEY_ALREADY_PRESENT": return TErrorCode_KUDU_KEY_ALREADY_PRESENT, nil 
  case "KUDU_NOT_FOUND": return TErrorCode_KUDU_NOT_FOUND, nil 
  case "KUDU_SESSION_ERROR": return TErrorCode_KUDU_SESSION_ERROR, nil 
  case "AVRO_UNSUPPORTED_TYPE": return TErrorCode_AVRO_UNSUPPORTED_TYPE, nil 
  case "AVRO_INVALID_DECIMAL": return TErrorCode_AVRO_INVALID_DECIMAL, nil 
  case "KUDU_NULL_CONSTRAINT_VIOLATION": return TErrorCode_KUDU_NULL_CONSTRAINT_VIOLATION, nil 
  case "PARQUET_TIMESTAMP_OUT_OF_RANGE": return TErrorCode_PARQUET_TIMESTAMP_OUT_OF_RANGE, nil 
  case "SCRATCH_ALLOCATION_FAILED": return TErrorCode_SCRATCH_ALLOCATION_FAILED, nil 
  case "SCRATCH_READ_TRUNCATED": return TErrorCode_SCRATCH_READ_TRUNCATED, nil 
  case "KUDU_TIMESTAMP_OUT_OF_RANGE": return TErrorCode_KUDU_TIMESTAMP_OUT_OF_RANGE, nil 
  case "MAX_ROW_SIZE": return TErrorCode_MAX_ROW_SIZE, nil 
  case "IR_VERIFY_FAILED": return TErrorCode_IR_VERIFY_FAILED, nil 
  case "MINIMUM_RESERVATION_UNAVAILABLE": return TErrorCode_MINIMUM_RESERVATION_UNAVAILABLE, nil 
  case "ADMISSION_REJECTED": return TErrorCode_ADMISSION_REJECTED, nil 
  case "ADMISSION_TIMED_OUT": return TErrorCode_ADMISSION_TIMED_OUT, nil 
  case "THREAD_CREATION_FAILED": return TErrorCode_THREAD_CREATION_FAILED, nil 
  case "DISK_IO_ERROR": return TErrorCode_DISK_IO_ERROR, nil 
  case "DATASTREAM_RECVR_CLOSED": return TErrorCode_DATASTREAM_RECVR_CLOSED, nil 
  case "BAD_PRINCIPAL_FORMAT": return TErrorCode_BAD_PRINCIPAL_FORMAT, nil 
  case "LZ4_COMPRESSION_INPUT_TOO_LARGE": return TErrorCode_LZ4_COMPRESSION_INPUT_TOO_LARGE, nil 
  case "SASL_APP_NAME_MISMATCH": return TErrorCode_SASL_APP_NAME_MISMATCH, nil 
  }
  return TErrorCode(0), fmt.Errorf("not a valid TErrorCode string")
}


func TErrorCodePtr(v TErrorCode) *TErrorCode { return &v }

func (p TErrorCode) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *TErrorCode) UnmarshalText(text []byte) error {
q, err := TErrorCodeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *TErrorCode) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = TErrorCode(v)
return nil
}

func (p * TErrorCode) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
