package sudoku

import (
	"testing"
)

func TestCreateBitMask(t *testing.T) {
	tests := []struct {
		name string
		args []uint8
		want uint16
	}{
		{
			name: "Test 1",
			args: []uint8{1, 2, 3, 5},
			want: 0b0000010111,
		},
		{
			name: "Test 2",
			args: []uint8{},
			want: 0b0000000000,
		},
		{
			name: "Test 3",
			args: []uint8{1, 3, 5},
			want: 0b0000010101,
		},
		{
			name: "Test 4",
			args: []uint8{1, 2, 3, 4, 5},
			want: 0b0000011111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBitMask(tt.args); got != tt.want {
				t.Errorf("createBitMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCandidates(t *testing.T) {
	tests := []struct {
		name string
		args []uint16
		want uint16
	}{
		{
			name: "Test 1",
			args: []uint16{},
			want: 0b001111111111,
		},
		{
			name: "Test 2",
			args: []uint16{0b001000000000},
			want: 0b000111111111,
		},
		{
			name: "Test 3",
			args: []uint16{0b001000000001, 0b000000000001},
			want: 0b000111111110,
		},
		{
			name: "Test 4",
			args: []uint16{0b000100000001, 0b000000000010, 0b000000000000},
			want: 0b001011111100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createCandidate(0, tt.args); got != tt.want {
				t.Errorf("createCandidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
