package puzzle17

import (
	"testing"
)

func TestRunProgram1(t *testing.T) {
	testCases := []struct {
		name                                     string
		registerA, registerB, registerC          int
		program                                  []int
		expectedRegA, expectedRegB, expectedRegC *int
		expectedOutput                           *string
	}{
		{
			name:         "Test 1",
			registerC:    9,
			program:      []int{2, 6},
			expectedRegB: toPtr(1),
		},
		{
			name:           "Test 2",
			registerA:      10,
			program:        []int{5, 0, 5, 1, 5, 4},
			expectedOutput: toPtr("0,1,2"),
		},
		{
			name:           "Test 3",
			registerA:      2024,
			program:        []int{0, 1, 5, 4, 3, 0},
			expectedRegA:   toPtr(0),
			expectedOutput: toPtr("4,2,5,6,7,7,7,7,3,1,0"),
		},
		{
			name:         "Test 4",
			registerB:    29,
			program:      []int{1, 7},
			expectedRegB: toPtr(26),
		},
		{
			name:         "Test 5",
			registerB:    2024,
			registerC:    43690,
			program:      []int{4, 0},
			expectedRegB: toPtr(44354),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			regA, regB, regC, output := runProgram(tc.registerA, tc.registerB, tc.registerC, tc.program)
			if tc.expectedRegA != nil && regA != *tc.expectedRegA {
				t.Errorf("Expected register A to be %v, got %v", *tc.expectedRegA, regA)
			}
			if tc.expectedRegB != nil && regB != *tc.expectedRegB {
				t.Errorf("Expected register B to be %v, got %v", *tc.expectedRegB, regB)
			}
			if tc.expectedRegC != nil && regC != *tc.expectedRegC {
				t.Errorf("Expected register C to be %v, got %v", *tc.expectedRegC, regC)
			}
			if tc.expectedOutput != nil && output != *tc.expectedOutput {
				t.Errorf("Expected output to be %v, got %v", *tc.expectedOutput, output)
			}
		})
	}
}

func toPtr[T int | string](value T) *T {
	return &value
}
