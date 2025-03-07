package hello

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	result := add(5, 2)
	expected := 7

	if result != expected {
		t.Errorf("Add(5,2)= %d; want %d", result, expected)
	} else {
		fmt.Println("test passed")
	}

}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
