package cp857_test

import (
	"testing"

	"github.com/gunim/cp857"
)

func TestDecodeByte(t *testing.T) {
	tests := []struct {
		name string
		x    byte
		want rune
	}{
		{"Ğ", 166, 0x011E},
		{"ğ", 167, 0x011F},
		{"╦", 203, 0x2566},
		{"¶", 244, 0x00B6},
		{"d", 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cp857.DecodeByte(tt.x); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func EncodeRune(t *testing.T) {
	tests := []struct {
		name   string
		r      rune
		wantB  byte
		wantOk bool
	}{
		{"Ğ", 0x011E, 166, true},
		{"ğ", 0x011F, 167, true},
		{"╦", 0x2566, 203, true},
		{"¶", 0x00B6, 244, true},
		{"d", 100, 100, true},
		{"⌘", '\u2318', '\x1A', false},
		{"😀", '\U0001F600', '\x1A', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB, gotOk := cp857.EncodeRune(tt.r); gotB != tt.wantB || gotOk != tt.wantOk {
				t.Errorf("Encode() = %v, %v; want %v, %v", gotB, gotOk, tt.wantB, tt.wantOk)
			}
		})
	}
}
