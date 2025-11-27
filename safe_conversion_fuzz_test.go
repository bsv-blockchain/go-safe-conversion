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

// FuzzIntToUint32 validates IntToUint32 with random inputs.
func FuzzIntToUint32(f *testing.F) {
	f.Add(0)
	f.Add(math.MaxUint32)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, v int) {
		r, err := safe.IntToUint32(v)
		if v < 0 || v > math.MaxUint32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(v), r)
	})
}

// FuzzUint64ToUint32 validates Uint64ToUint32 with random inputs.
func FuzzUint64ToUint32(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint32))
	f.Add(uint64(math.MaxUint32 + 1))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.Uint64ToUint32(v)
		if v > uint64(math.MaxUint32) {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(v), r)
	})
}

// FuzzInt64ToUint64 validates Int64ToUint64 with random inputs.
func FuzzInt64ToUint64(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(-1))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v int64) {
		r, err := safe.Int64ToUint64(v)
		if v < 0 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint64(v), r)
	})
}

// FuzzIntToUint64 validates IntToUint64 with random inputs.
func FuzzIntToUint64(f *testing.F) {
	f.Add(0)
	f.Add(-1)
	f.Add(math.MaxInt)
	f.Fuzz(func(t *testing.T, v int) {
		r, err := safe.IntToUint64(v)
		if v < 0 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint64(v), r)
	})
}

// FuzzUint64ToInt validates Uint64ToInt with random inputs.
func FuzzUint64ToInt(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.Uint64ToInt(v)
		if v > math.MaxInt {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int(v), r)
	})
}

// FuzzInt64ToInt32 validates Int64ToInt32 with random inputs.
func FuzzInt64ToInt32(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MaxInt32))
	f.Add(int64(math.MinInt32))
	f.Fuzz(func(t *testing.T, v int64) {
		r, err := safe.Int64ToInt32(v)
		if v < math.MinInt32 || v > math.MaxInt32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int32(v), r)
	})
}

// FuzzIntToInt32 validates IntToInt32 with random inputs.
func FuzzIntToInt32(f *testing.F) {
	f.Add(0)
	f.Add(math.MaxInt32)
	f.Add(math.MinInt32)
	f.Fuzz(func(t *testing.T, v int) {
		r, err := safe.IntToInt32(v)
		if v < math.MinInt32 || v > math.MaxInt32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int32(v), r)
	})
}

// FuzzInt32ToUint32 validates Int32ToUint32 with random inputs.
func FuzzInt32ToUint32(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(-1))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		r, err := safe.Int32ToUint32(v)
		if v < 0 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(v), r)
	})
}

// FuzzInt64ToUint32 validates Int64ToUint32 with random inputs.
func FuzzInt64ToUint32(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(-1))
	f.Add(int64(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v int64) {
		r, err := safe.Int64ToUint32(v)
		if v < 0 || v > math.MaxUint32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(v), r)
	})
}

// FuzzBigWordToUint32 validates BigWordToUint32 with random inputs.
func FuzzBigWordToUint32(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.BigWordToUint32(big.Word(v))
		if v > math.MaxUint32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(v), r)
	})
}

// FuzzIntToUint16 validates IntToUint16 with random inputs.
func FuzzIntToUint16(f *testing.F) {
	f.Add(0)
	f.Add(math.MaxUint16)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, v int) {
		r, err := safe.IntToUint16(v)
		if v < 0 || v > math.MaxUint16 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint16(v), r)
	})
}

// FuzzIntToInt16 validates IntToInt16 with random inputs.
func FuzzIntToInt16(f *testing.F) {
	f.Add(0)
	f.Add(safe.MinInt16)
	f.Add(safe.MaxInt16)
	f.Fuzz(func(t *testing.T, v int) {
		r, err := safe.IntToInt16(v)
		if v < safe.MinInt16 || v > safe.MaxInt16 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int16(v), r)
	})
}

// FuzzUintToUint32 validates UintToUint32 with random inputs.
func FuzzUintToUint32(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint) {
		r, err := safe.UintToUint32(v)
		if v > math.MaxUint32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(v), r)
	})
}

// FuzzTimeToUint32 validates TimeToUint32 with random inputs.
func FuzzTimeToUint32(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1))
	f.Fuzz(func(t *testing.T, ts int64) {
		r, err := safe.TimeToUint32(time.Unix(ts, 0))
		if ts < 0 || ts > math.MaxUint32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint32(ts), r)
	})
}

// FuzzUint32ToUint8 validates Uint32ToUint8 with random inputs.
func FuzzUint32ToUint8(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint8))
	f.Fuzz(func(t *testing.T, v uint32) {
		r, err := safe.Uint32ToUint8(v)
		if v > math.MaxUint8 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint8(v), r)
	})
}

// FuzzUintptrToInt validates UintptrToInt with random inputs.
func FuzzUintptrToInt(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.UintptrToInt(uintptr(v))
		if v > uint64(math.MaxInt) {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int(v), r)
	})
}

// FuzzConvertUint64ToInt64 validates Uint64ToInt64 with random inputs.
func FuzzConvertUint64ToInt64(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.Uint64ToInt64(v)
		if v > math.MaxInt64 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int64(v), r)
	})
}

// FuzzConvertUint32ToInt32 validates Uint32ToInt32 with random inputs.
func FuzzConvertUint32ToInt32(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v uint32) {
		r, err := safe.Uint32ToInt32(v)
		if v > math.MaxInt32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int32(v), r)
	})
}

// FuzzConvertUint64ToInt32 validates Uint64ToInt32 with random inputs.
func FuzzConvertUint64ToInt32(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.Uint64ToInt32(v)
		if v > math.MaxInt32 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int32(v), r)
	})
}

// FuzzUint32ToInt64 validates Uint32ToInt64 with random inputs.
func FuzzUint32ToInt64(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		r, err := safe.Uint32ToInt64(v)
		require.NoError(t, err)
		assert.Equal(t, int64(v), r)
	})
}

// FuzzUint32ToUint64 validates Uint32ToUint64 with random inputs.
func FuzzUint32ToUint64(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		r, err := safe.Uint32ToUint64(v)
		require.NoError(t, err)
		assert.Equal(t, uint64(v), r)
	})
}

// FuzzUint64ToUint16 validates Uint64ToUint16 with random inputs.
func FuzzUint64ToUint16(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint16))
	f.Fuzz(func(t *testing.T, v uint64) {
		r, err := safe.Uint64ToUint16(v)
		if v > math.MaxUint16 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, uint16(v), r)
	})
}
