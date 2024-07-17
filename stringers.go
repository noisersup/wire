// Code generated by "stringer -linecomment -output stringers.go -type OpCode,OpMsgFlagBit,OpQueryFlagBit,OpReplyFlagBit"; DO NOT EDIT.

package wire

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpCodeReply-1]
	_ = x[OpCodeUpdate-2001]
	_ = x[OpCodeInsert-2002]
	_ = x[OpCodeGetByOID-2003]
	_ = x[OpCodeQuery-2004]
	_ = x[OpCodeGetMore-2005]
	_ = x[OpCodeDelete-2006]
	_ = x[OpCodeKillCursors-2007]
	_ = x[OpCodeCompressed-2012]
	_ = x[OpCodeMsg-2013]
}

const (
	_OpCode_name_0 = "OP_REPLY"
	_OpCode_name_1 = "OP_UPDATEOP_INSERTOP_GET_BY_OIDOP_QUERYOP_GET_MOREOP_DELETEOP_KILL_CURSORS"
	_OpCode_name_2 = "OP_COMPRESSEDOP_MSG"
)

var (
	_OpCode_index_1 = [...]uint8{0, 9, 18, 31, 39, 50, 59, 74}
	_OpCode_index_2 = [...]uint8{0, 13, 19}
)

func (i OpCode) String() string {
	switch {
	case i == 1:
		return _OpCode_name_0
	case 2001 <= i && i <= 2007:
		i -= 2001
		return _OpCode_name_1[_OpCode_index_1[i]:_OpCode_index_1[i+1]]
	case 2012 <= i && i <= 2013:
		i -= 2012
		return _OpCode_name_2[_OpCode_index_2[i]:_OpCode_index_2[i+1]]
	default:
		return "OpCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpMsgChecksumPresent-1]
	_ = x[OpMsgMoreToCome-2]
	_ = x[OpMsgExhaustAllowed-65536]
}

const (
	_OpMsgFlagBit_name_0 = "checksumPresentmoreToCome"
	_OpMsgFlagBit_name_1 = "exhaustAllowed"
)

var (
	_OpMsgFlagBit_index_0 = [...]uint8{0, 15, 25}
)

func (i OpMsgFlagBit) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _OpMsgFlagBit_name_0[_OpMsgFlagBit_index_0[i]:_OpMsgFlagBit_index_0[i+1]]
	case i == 65536:
		return _OpMsgFlagBit_name_1
	default:
		return "OpMsgFlagBit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpQueryTailableCursor-2]
	_ = x[OpQuerySlaveOk-4]
	_ = x[OpQueryOplogReplay-8]
	_ = x[OpQueryNoCursorTimeout-16]
	_ = x[OpQueryAwaitData-32]
	_ = x[OpQueryExhaust-64]
	_ = x[OpQueryPartial-128]
}

const (
	_OpQueryFlagBit_name_0 = "TailableCursor"
	_OpQueryFlagBit_name_1 = "SlaveOk"
	_OpQueryFlagBit_name_2 = "OplogReplay"
	_OpQueryFlagBit_name_3 = "NoCursorTimeout"
	_OpQueryFlagBit_name_4 = "AwaitData"
	_OpQueryFlagBit_name_5 = "Exhaust"
	_OpQueryFlagBit_name_6 = "Partial"
)

func (i OpQueryFlagBit) String() string {
	switch {
	case i == 2:
		return _OpQueryFlagBit_name_0
	case i == 4:
		return _OpQueryFlagBit_name_1
	case i == 8:
		return _OpQueryFlagBit_name_2
	case i == 16:
		return _OpQueryFlagBit_name_3
	case i == 32:
		return _OpQueryFlagBit_name_4
	case i == 64:
		return _OpQueryFlagBit_name_5
	case i == 128:
		return _OpQueryFlagBit_name_6
	default:
		return "OpQueryFlagBit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpReplyCursorNotFound-1]
	_ = x[OpReplyQueryFailure-2]
	_ = x[OpReplyShardConfigStale-4]
	_ = x[OpReplyAwaitCapable-8]
}

const (
	_OpReplyFlagBit_name_0 = "CursorNotFoundQueryFailure"
	_OpReplyFlagBit_name_1 = "ShardConfigStale"
	_OpReplyFlagBit_name_2 = "AwaitCapable"
)

var (
	_OpReplyFlagBit_index_0 = [...]uint8{0, 14, 26}
)

func (i OpReplyFlagBit) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _OpReplyFlagBit_name_0[_OpReplyFlagBit_index_0[i]:_OpReplyFlagBit_index_0[i+1]]
	case i == 4:
		return _OpReplyFlagBit_name_1
	case i == 8:
		return _OpReplyFlagBit_name_2
	default:
		return "OpReplyFlagBit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}