package utils

import (
	"reflect"
	"testing"
)

func TestCountChar(t *testing.T) {
	testCases := []struct {
		s        string
		c        rune
		expected int
	}{
		{"hello", 'l', 2},
		{"world", 'o', 1},
		{"", 'a', 0},
	}

	for _, tc := range testCases {
		result := CountChar(tc.s, tc.c)
		if result != tc.expected {
			t.Errorf("CountChar(%q, %q) = %v; want %v", tc.s, tc.c, result, tc.expected)
		}
	}
}

func TestStrToInt(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{"123", 123},
		{"-456", -456},
		{"0", 0},
	}

	for _, tc := range testCases {
		result := StrToInt(tc.s)
		if result != tc.expected {
			t.Errorf("StrToInt(%q) = %v; want %v", tc.s, result, tc.expected)
		}
	}
}

func TestGroupChar(t *testing.T) {
	testCases := []struct {
		s        string
		expected []string
	}{
		{"hello", []string{"h", "e", "ll", "o"}},
		{"world", []string{"w", "o", "r", "l", "d"}},
		{"", []string{}},
		{"111222333", []string{"111", "222", "333"}},
	}

	for _, tc := range testCases {
		result := GroupChar(tc.s)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("GroupChar(%q) = %v; want %v", tc.s, result, tc.expected)
		}
	}
}
