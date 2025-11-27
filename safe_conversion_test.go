// Package safeconversion_test provides tests for the safe conversion functions.
package safeconversion_test

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	safe "github.com/bsv-blockchain/go-safe-conversion"
)

const (
	zeroValueName         = "zero value"
	positiveValueName     = "positive value"
	maxUint32Name         = "max uint32"
	negativeValueName     = "negative value"
	valueTooLargeName     = "value too large"
	valueMuchTooLargeName = "value much too large"
	maxIntName            = "max int"
	maxInt32Name          = "max int32"
	minInt32Name          = "min int32"
	maxInt64Name          = "max int64"
	minInt64Name          = "min int64"
	valueTooSmallName     = "value too small"
	maxUint16Name         = "max uint16"
	maxUint64Name         = "max uint64"
)

// TestIntToUint32 tests the conversion from int to uint32.
func TestIntToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		expect  uint32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint32Name, math.MaxUint32, math.MaxUint32, false},
		{negativeValueName, -1, 0, true},
		{valueTooLargeName, math.MaxInt, 0, true}, // Assuming this is on a 64-bit system where MaxInt > MaxUint32
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.IntToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint64ToUint32 tests the conversion from uint64 to uint32.
func TestUint64ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		expect  uint32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint32Name, uint64(math.MaxUint32), math.MaxUint32, false},
		{valueTooLargeName, uint64(math.MaxUint32) + 1, 0, true},
		{valueMuchTooLargeName, math.MaxUint64, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint64ToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestInt64ToUint64 tests the conversion from int64 to uint64.
func TestInt64ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		expect  uint64
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt64Name, math.MaxInt64, uint64(math.MaxInt64), false},
		{negativeValueName, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Int64ToUint64(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestIntToUint64 tests the conversion from int to uint64.
func TestIntToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		expect  uint64
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxIntName, math.MaxInt, uint64(math.MaxInt), false},
		{negativeValueName, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.IntToUint64(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint64ToInt tests the conversion from uint64 to int.
func TestUint64ToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		expect  int
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxIntName, uint64(math.MaxInt), math.MaxInt, false},
		{valueTooLargeName, uint64(math.MaxInt) + 1, 0, true},
		{valueMuchTooLargeName, math.MaxUint64, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint64ToInt(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestInt64ToInt32 tests the conversion from int64 to int32.
func TestInt64ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		expect  int32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt32Name, int64(math.MaxInt32), math.MaxInt32, false},
		{minInt32Name, int64(math.MinInt32), math.MinInt32, false},
		{valueTooLargeName, int64(math.MaxInt32) + 1, 0, true},
		{valueTooSmallName, int64(math.MinInt32) - 1, 0, true},
		{maxInt64Name, math.MaxInt64, 0, true},
		{minInt64Name, math.MinInt64, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Int64ToInt32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestIntToInt32 tests the conversion from int64 to int32.
func TestIntToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		expect  int32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt32Name, math.MaxInt32, math.MaxInt32, false},
		{minInt32Name, math.MinInt32, math.MinInt32, false},
		// These tests below assume a 64-bit system where int can exceed int32 range
		{valueTooLargeName, int(math.MaxInt32) + 1, 0, true},
		{valueTooSmallName, int(math.MinInt32) - 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.IntToInt32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestInt32ToUint32 tests the conversion from int32 to uint32.
func TestInt32ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int32
		expect  uint32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt32Name, math.MaxInt32, uint32(math.MaxInt32), false},
		{negativeValueName, -1, 0, true},
		{minInt32Name, math.MinInt32, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Int32ToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestInt64ToUint32 tests the conversion from int64 to uint32.
func TestInt64ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		expect  uint32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint32Name, int64(math.MaxUint32), math.MaxUint32, false},
		{negativeValueName, -1, 0, true},
		{valueTooLargeName, int64(math.MaxUint32) + 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Int64ToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestBigWordToUint32 tests the conversion from big.Word to uint32.
func TestBigWordToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   big.Word
		expect  uint32
		wantErr bool
	}{
		{zeroValueName, big.Word(0), 0, false},
		{positiveValueName, big.Word(100), 100, false},
		{maxUint32Name, big.Word(math.MaxUint32), math.MaxUint32, false},
		// This test might not work on 32-bit systems where big.Word is uint32
		{valueTooLargeName, big.Word(uint64(math.MaxUint32) + 1), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.BigWordToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestIntToUint16 tests the conversion from int to uint16.
func TestIntToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		expect  uint16
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint16Name, math.MaxUint16, math.MaxUint16, false},
		{negativeValueName, -1, 0, true},
		{valueTooLargeName, int(math.MaxUint16) + 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.IntToUint16(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestIntToInt16 tests the conversion from int to int16.
func TestIntToInt16(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		expect  int16
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{"max int16", 32767, 32767, false},   // MaxInt16 is 32767
		{"min int16", -32768, -32768, false}, // MinInt16 is -32768
		{valueTooLargeName, 32768, 0, true},  // MaxInt16 + 1
		{valueTooSmallName, -32769, 0, true}, // MinInt16 - 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.IntToInt16(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUintToUint32 tests the conversion from uint to uint32.
func TestUintToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint
		expect  uint32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint32Name, uint(math.MaxUint32), math.MaxUint32, false},
		// This test might not work on 32-bit systems where uint is uint32
		{valueTooLargeName, uint(uint64(math.MaxUint32) + 1), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.UintToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestTimeToUint32 tests the conversion from time.Time to uint32.
func TestTimeToUint32(t *testing.T) {
	tests := []struct {
		name    string
		input   time.Time
		expect  uint32
		wantErr bool
	}{
		{"epoch start", time.Unix(0, 0), 0, false},
		{"positive timestamp", time.Unix(100, 0), 100, false},
		{"max uint32 timestamp", time.Unix(math.MaxUint32, 0), math.MaxUint32, false},
		{"negative timestamp", time.Unix(-1, 0), 0, true},
		{"timestamp too large", time.Unix(int64(math.MaxUint32)+1, 0), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.TimeToUint32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint32ToUint8 tests the conversion from uint32 to uint8.
func TestUint32ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		expect  uint8
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{"max uint8", uint32(math.MaxUint8), math.MaxUint8, false},
		{valueTooLargeName, uint32(math.MaxUint8) + 1, 0, true},
		{maxUint32Name, math.MaxUint32, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint32ToUint8(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUintptrToInt tests the conversion from uintptr to int.
func TestUintptrToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   uintptr
		expect  int
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxIntName, uintptr(math.MaxInt), math.MaxInt, false},
		{valueTooLargeName, uintptr(math.MaxInt) + 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.UintptrToInt(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint64ToInt64 tests the conversion from uint64 to int64.
func TestUint64ToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		expect  int64
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt64Name, uint64(math.MaxInt64), math.MaxInt64, false},
		{valueTooLargeName, uint64(math.MaxInt64) + 1, 0, true},
		{maxUint64Name, math.MaxUint64, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint64ToInt64(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint32ToInt32 tests the conversion from uint32 to int32.
func TestUint32ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		expect  int32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt32Name, uint32(math.MaxInt32), math.MaxInt32, false},
		{valueTooLargeName, uint32(math.MaxInt32) + 1, 0, true},
		{maxUint32Name, math.MaxUint32, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint32ToInt32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint64ToInt32 tests the conversion from uint64 to int32.
func TestUint64ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		expect  int32
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxInt32Name, uint64(math.MaxInt32), math.MaxInt32, false},
		{valueTooLargeName, uint64(math.MaxInt32) + 1, 0, true},
		{maxUint64Name, math.MaxUint64, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint64ToInt32(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint32ToInt64 tests the conversion from uint32 to int64.
func TestUint32ToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		expect  int64
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint32Name, math.MaxUint32, int64(math.MaxUint32), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint32ToInt64(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint32ToUint64 tests the conversion from uint32 to uint64.
func TestUint32ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   uint32
		expect  uint64
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint32Name, math.MaxUint32, uint64(math.MaxUint32), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint32ToUint64(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// TestUint64ToUint16 tests the conversion from uint64 to uint16.
func TestUint64ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		expect  uint16
		wantErr bool
	}{
		{zeroValueName, 0, 0, false},
		{positiveValueName, 100, 100, false},
		{maxUint16Name, uint64(math.MaxUint16), math.MaxUint16, false},
		{valueTooLargeName, uint64(math.MaxUint16) + 1, 0, true},
		{maxUint64Name, math.MaxUint64, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe.Uint64ToUint16(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expect, result)
		})
	}
}
