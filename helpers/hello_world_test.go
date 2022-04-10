package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dira")

	if result != "Hello Dira" {
		t.Error("Expected Hello Dira, got ", result)
	}
}

func TestHellWolrdAssert(t *testing.T) {
	result := HelloWorld("Dira")
	assert.Equal(t, "Hello Dira", result, "Result must be Hello Dira")
}

func TestHellWolrdRequire(t *testing.T) {
	result := HelloWorld("Dira")
	require.Equal(t, "Hello Dira", result, "Result must be Hello Dira")
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dira")
	}
}

func BenchmarkHelloWorldDira(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Apsyadira")
	}
}