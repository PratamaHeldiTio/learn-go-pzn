package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type benchmarkData struct {
	name    string
	request string
}

func BenchmarkTableTest(b *testing.B) {
	dataBench := []benchmarkData{
		{
			name:    "heldi",
			request: "heldi",
		},
		{
			name:    "tio",
			request: "tio",
		},
		{
			name:    "pratama",
			request: "pratama",
		},
		{
			name:    "heldi tio pratama",
			request: "heldi tio pratama",
		},
	}

	for _, value := range dataBench {
		b.Run(value.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(value.request)
			}
		})
	}
}

// func BenchmarkHelloWorld(b *testing.B) {
// 	// * sub benchmark
// 	b.Run("heldi", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("heldi")
// 		}
// 	})
// 	b.Run("tio", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("tio")
// 		}
// 	})
// 	b.Run("pratama", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("pratama")
// 		}
// 	})

// }

// func BenchmarkHelloWorldTio(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		HelloWorld("tio")
// 	}
// }

func TestMain(m *testing.M) {
	// *Before
	// fmt.Println("ini before")

	m.Run()

	// * After
	// fmt.Println("ini after")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("heldi")

	if result != "hello heldi" {
		t.Error("return must be hello heldi")
	}

	fmt.Println("ini Fail()")
}

func TestHelloWorldHeldi(t *testing.T) {
	result := HelloWorld("heldi")

	if result != "hello heldi" {
		t.Fatal("return must be hello heldi")
	}

	fmt.Println("ini Fail()")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("heldi")

	assert.Equal(t, "hello heldi", result, "return must be hello heldi")

	fmt.Println("ini pakai asssert")

}

func TestSubTest(t *testing.T) {
	t.Run("heldi", func(t *testing.T) {
		result := HelloWorld("heldi")
		assert.Equal(t, "hello heldi", result, "return must be hello heldi")
	})

	t.Run("tio", func(t *testing.T) {
		result := HelloWorld("tio")
		assert.Equal(t, "hello tio", result, "return must be hello heldi")
	})
}

func TestTableTest(t *testing.T) {
	dataTests := []DataTest{
		{
			name:     "heldi",
			request:  "heldi",
			expected: "hello heldi",
		},
		{
			name:     "tio",
			request:  "tio",
			expected: "hello tio",
		},
		{
			name:     "pratama",
			request:  "pratama",
			expected: "hello pratama",
		},
	}

	for _, test := range dataTests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "return must be hello "+test.request)
		})
	}
}

type DataTest struct {
	name     string
	request  string
	expected string
}
