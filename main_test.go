package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		expectedNil bool
	}{
		{
			name:        "positive size",
			size:        1000,
			expectedNil: false,
		},
		{
			name:        "zero size",
			size:        0,
			expectedNil: true,
		},
		{
			name:        "negative size",
			size:        -1,
			expectedNil: true,
		},
	}

	for _, v := range tests {
		result := generateRandomElements(v.size)
		if v.expectedNil {
			assert.Nil(t, result)
		} else {
			require.NotNil(t, result)
		}

	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "empty slice",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			data:     []int{100},
			expected: 100,
		},
		{
			name:     "multiple elements",
			data:     []int{1, 5, 10},
			expected: 10,
		},
		{
			name:     "all same elements",
			data:     []int{5, 5, 5, 5, 5, 5, 5, 5},
			expected: 5,
		},
	}

	for _, v := range tests {
		result := maximum(v.data)
		assert.Equal(t, v.expected, result)
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "empty slice",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			data:     []int{100},
			expected: 100,
		},
		{
			name:     "multiple elements",
			data:     []int{1, 5, 10},
			expected: 10,
		},
		{
			name:     "all same elements",
			data:     []int{5, 5, 5, 5, 5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "with remainder",
			data:     []int{5, 5, 5, 5, 5, 5, 5, 5, 7},
			expected: 7,
		},
	}

	for _, v := range tests {
		result := maxChunks(v.data)
		assert.Equal(t, v.expected, result)
	}
}
