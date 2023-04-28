package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringBuilderConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for x := 0; x < 1000; x++ {
			sb.WriteString("b")
		}
	}

}

func BenchmarkSringAddConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb string
		for x := 0; x < 1000; x++ {
			sb += "b"
		}

	}
}

func BenchmarkSprintfConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb string
		for x := 0; x < 1000; x++ {
			sb = fmt.Sprintf("%v%v", sb, "b")
		}

	}

}
