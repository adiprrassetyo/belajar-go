package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// unit test dengan table test
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Adi",
			request: "Adi",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
		{
			name:    "AdiKurniawanPrasetyo",
			request: "Adi Kurniawan Prasetyo",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// unit test dengan sub test
func BenchmarkSub(b *testing.B) {
	b.Run("Adi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adi")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}

// unit test benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adi")
	}
}

// unit test benchmark
func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

// table test dengan iterasi menggunakan for loop (tidak perlu membuat file unit test baru)
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Adi",
			request:  "Adi",
			expected: "Hello Adi",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Prasetyo",
			request:  "Prasetyo",
			expected: "Hello Prasetyo",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// sub test untuk memisahkan unit test yang berbeda beda menjadi satu file unit test yang sama (tidak perlu membuat file unit test baru)
func TestSubTest(t *testing.T) {
	t.Run("Adi", func(t *testing.T) {
		result := HelloWorld("Adi")
		require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	})
	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	})
	t.Run("Prasetyo", func(t *testing.T) {
		result := HelloWorld("Prasetyo")
		require.Equal(t, "Hello Prasetyo", result, "Result must be 'Hello Prasetyo'")
	})
}

// test main untuk setup sebelum dan sesudah unit test bisa untuk setup database
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

// 3. testing with testify (recommended) -use skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS") // skip test
	}

	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
}

// 2. testing with testify (recommended) -use require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	fmt.Println("TestHelloWorld with Require Done") // will not run
}

// 2. testing with testify (recommended) -use assert
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Adi")
	assert.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	fmt.Println("TestHelloWorld with Assert Done") // will run
}

// 1. testing with if else (not recommended)
func TestHelloWorldAdi(t *testing.T) {
	result := HelloWorld("Adi")

	if result != "Hello Adi" {
		// error
		t.Error("Result must be 'Hello Adi'")
	}

	fmt.Println("TestHelloWorldAdi Done")
}

func TestHelloWorldPrasetyo(t *testing.T) {
	result := HelloWorld("Prasetyo")

	if result != "Hello Prasetyo" {
		// error
		t.Fatal("Result must be 'Hello Prasetyo'")
	}

	fmt.Println("TestHelloWorldPrasetyo Done")
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")

	if result != "Hello Kurniawan" {
		// error
		t.Fatal("Result must be 'Hello Kurniawan'")
	}

	fmt.Println("TestHelloWorldKurniawan Done")
}
