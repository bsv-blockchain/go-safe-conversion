package safeconversion

import (
	"fmt"
	"math/big"
	"time"
)

// ExampleIntToUint32 demonstrates converting an int to a uint32.
func ExampleIntToUint32() {
	v, err := IntToUint32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint64ToUint32 demonstrates converting a uint64 to a uint32.
func ExampleUint64ToUint32() {
	v, err := Uint64ToUint32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleInt64ToUint64 demonstrates converting an int64 to a uint64.
func ExampleInt64ToUint64() {
	v, err := Int64ToUint64(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleIntToUint64 demonstrates converting an int to a uint64.
func ExampleIntToUint64() {
	v, err := IntToUint64(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint64ToInt demonstrates converting a uint64 to an int.
func ExampleUint64ToInt() {
	v, err := Uint64ToInt(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleInt64ToInt32 demonstrates converting an int64 to an int32.
func ExampleInt64ToInt32() {
	v, err := Int64ToInt32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleIntToInt32 demonstrates converting an int to an int32.
func ExampleIntToInt32() {
	v, err := IntToInt32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleInt32ToUint32 demonstrates converting an int32 to a uint32.
func ExampleInt32ToUint32() {
	v, err := Int32ToUint32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleInt64ToUint32 demonstrates converting an int64 to a uint32.
func ExampleInt64ToUint32() {
	v, err := Int64ToUint32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleBigWordToUint32 demonstrates converting a big.Word to a uint32.
func ExampleBigWordToUint32() {
	v, err := BigWordToUint32(big.Word(42))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleIntToUint16 demonstrates converting an int to a uint16.
func ExampleIntToUint16() {
	v, err := IntToUint16(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleIntToInt16 demonstrates converting an int to an int16.
func ExampleIntToInt16() {
	v, err := IntToInt16(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUintToUint32 demonstrates converting a uint to a uint32.
func ExampleUintToUint32() {
	v, err := UintToUint32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleTimeToUint32 demonstrates converting a time.Time to a uint32.
func ExampleTimeToUint32() {
	v, err := TimeToUint32(time.Unix(42, 0))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint32ToUint8 demonstrates converting a uint32 to a uint8.
func ExampleUint32ToUint8() {
	v, err := Uint32ToUint8(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUintptrToInt demonstrates converting a uintptr to an int.
func ExampleUintptrToInt() {
	v, err := UintptrToInt(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint64ToInt64 demonstrates converting a uint64 to an int64.
func ExampleUint64ToInt64() {
	v, err := Uint64ToInt64(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint32ToInt32 demonstrates converting a uint32 to an int32.
func ExampleUint32ToInt32() {
	v, err := Uint32ToInt32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint64ToInt32 demonstrates converting a uint64 to an int32.
func ExampleUint64ToInt32() {
	v, err := Uint64ToInt32(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}

// ExampleUint32ToInt64 demonstrates converting a uint32 to an int64.
func ExampleUint32ToInt64() {
	v, _ := Uint32ToInt64(42)
	fmt.Println(v)
	// Output: 42
}

// ExampleUint32ToUint64 demonstrates converting a uint32 to a uint64.
func ExampleUint32ToUint64() {
	v, _ := Uint32ToUint64(42)
	fmt.Println(v)
	// Output: 42
}

// ExampleUint64ToUint16 demonstrates converting a uint64 to a uint16.
func ExampleUint64ToUint16() {
	v, err := Uint64ToUint16(42)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)
	// Output: 42
}
