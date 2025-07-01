package safeconversion_test

import (
	"math/big"
	"testing"
	"time"

	safe "github.com/bsv-blockchain/go-safe-conversion"
)

var (
	benchUint8  uint8
	benchUint16 uint16
	benchUint32 uint32
	benchUint64 uint64
	benchInt    int
	benchInt32  int32
	benchInt64  int64
)

// BenchmarkBigWordToUint32 benchmarks the performance of BigWordToUint32.
func BenchmarkBigWordToUint32(b *testing.B) {
	var r uint32
	var err error
	v := big.Word(100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.BigWordToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}

// BenchmarkInt32ToUint32 benchmarks the performance of Int32ToUint32.
func BenchmarkInt32ToUint32(b *testing.B) {
	var r uint32
	var err error
	const v int32 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Int32ToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}

// BenchmarkInt64ToInt32 benchmarks the performance of Int64ToInt32.
func BenchmarkInt64ToInt32(b *testing.B) {
	var r int32
	var err error
	const v int64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Int64ToInt32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt32 = r
}

// BenchmarkInt64ToUint32 benchmarks the performance of Int64ToUint32.
func BenchmarkInt64ToUint32(b *testing.B) {
	var r uint32
	var err error
	const v int64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Int64ToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}

// BenchmarkInt64ToUint64 benchmarks the performance of Int64ToUint64.
func BenchmarkInt64ToUint64(b *testing.B) {
	var r uint64
	var err error
	const v int64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Int64ToUint64(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint64 = r
}

// BenchmarkIntToInt16 benchmarks the performance of IntToInt16.
func BenchmarkIntToInt16(b *testing.B) {
	var r int16
	var err error
	const v int = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.IntToInt16(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	// assign to global to avoid compiler optimization
	benchInt32 = int32(r)
}

// BenchmarkIntToInt32 benchmarks the performance of IntToInt32.
func BenchmarkIntToInt32(b *testing.B) {
	var r int32
	var err error
	const v int = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.IntToInt32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt32 = r
}

// BenchmarkIntToUint16 benchmarks the performance of IntToUint16.
func BenchmarkIntToUint16(b *testing.B) {
	var r uint16
	var err error
	const v int = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.IntToUint16(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint16 = r
}

// BenchmarkIntToUint32 benchmarks the performance of IntToUint32.
func BenchmarkIntToUint32(b *testing.B) {
	var r uint32
	var err error
	const v int = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.IntToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}

// BenchmarkIntToUint64 benchmarks the performance of IntToUint64.
func BenchmarkIntToUint64(b *testing.B) {
	var r uint64
	var err error
	const v int = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.IntToUint64(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint64 = r
}

// BenchmarkTimeToUint32 benchmarks the performance of TimeToUint32.
func BenchmarkTimeToUint32(b *testing.B) {
	var r uint32
	var err error
	v := time.Unix(100, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.TimeToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}

// BenchmarkUint32ToInt32 benchmarks the performance of Uint32ToInt32.
func BenchmarkUint32ToInt32(b *testing.B) {
	var r int32
	var err error
	const v uint32 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint32ToInt32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt32 = r
}

// BenchmarkUint32ToInt64 benchmarks the performance of Uint32ToInt64.
func BenchmarkUint32ToInt64(b *testing.B) {
	var r int64
	const v uint32 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, _ = safe.Uint32ToInt64(v)
	}
	benchInt64 = r
}

// BenchmarkUint32ToUint64 benchmarks the performance of Uint32ToUint64.
func BenchmarkUint32ToUint64(b *testing.B) {
	var r uint64
	const v uint32 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, _ = safe.Uint32ToUint64(v)
	}
	benchUint64 = r
}

// BenchmarkUint32ToUint8 benchmarks the performance of Uint32ToUint8.
func BenchmarkUint32ToUint8(b *testing.B) {
	var r uint8
	var err error
	const v uint32 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint32ToUint8(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint8 = r
}

// BenchmarkUint64ToInt benchmarks the performance of Uint64ToInt.
func BenchmarkUint64ToInt(b *testing.B) {
	var r int
	var err error
	const v uint64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint64ToInt(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt = r
}

// BenchmarkUint64ToInt32 benchmarks the performance of Uint64ToInt32.
func BenchmarkUint64ToInt32(b *testing.B) {
	var r int32
	var err error
	const v uint64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint64ToInt32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt32 = r
}

// BenchmarkUint64ToInt64 benchmarks the performance of Uint64ToInt64.
func BenchmarkUint64ToInt64(b *testing.B) {
	var r int64
	var err error
	const v uint64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint64ToInt64(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt64 = r
}

// BenchmarkUint64ToUint16 benchmarks the performance of Uint64ToUint16.
func BenchmarkUint64ToUint16(b *testing.B) {
	var r uint16
	var err error
	const v uint64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint64ToUint16(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint16 = r
}

// BenchmarkUint64ToUint32 benchmarks the performance of Uint64ToUint32.
func BenchmarkUint64ToUint32(b *testing.B) {
	var r uint32
	var err error
	const v uint64 = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.Uint64ToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}

// BenchmarkUintptrToInt benchmarks the performance of UintptrToInt.
func BenchmarkUintptrToInt(b *testing.B) {
	var r int
	var err error
	const v uintptr = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.UintptrToInt(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchInt = r
}

// BenchmarkUintToUint32 benchmarks the performance of UintToUint32.
func BenchmarkUintToUint32(b *testing.B) {
	var r uint32
	var err error
	const v uint = 100
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r, err = safe.UintToUint32(v)
		if err != nil {
			b.Fatal(err)
		}
	}
	benchUint32 = r
}
