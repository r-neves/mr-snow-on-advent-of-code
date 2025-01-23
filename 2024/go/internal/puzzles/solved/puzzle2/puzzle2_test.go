package puzzle2

import "testing"

func TestIsValidReportWithError(t *testing.T) {
	testCases := []struct {
		Name    string
		Report  []int
		IsValid bool
	}{
		{
			Name:    "Baseline ascending valid report",
			Report:  []int{1, 4, 5, 6, 7},
			IsValid: true,
		},
		{
			Name:    "Baseline descending valid report",
			Report:  []int{7, 6, 5, 4, 1},
			IsValid: true,
		},
		{
			Name:    "Valid: First index diff is out of range",
			Report:  []int{1, 6, 7, 8, 9},
			IsValid: true,
		},
		{
			Name:    "Valid: Last index diff is out of range",
			Report:  []int{5, 6, 7, 8, 20},
			IsValid: true,
		},
		{
			Name:    "Valid: One value is out of order",
			Report:  []int{1, 4, 3, 6, 7},
			IsValid: true,
		},
		{
			Name:    "Valid: Same value is out of order and big diff",
			Report:  []int{1, 8, 3, 4, 5},
			IsValid: true,
		},
		{
			Name:    "Valid: One repeated value",
			Report:  []int{1, 4, 5, 5, 6},
			IsValid: true,
		},
		{
			Name:    "Invalid: Two repeated values",
			Report:  []int{1, 1, 5, 5, 6},
			IsValid: false,
		},
		{
			Name:    "Invalid: Middle diff separates sequence",
			Report:  []int{1, 4, 5, 10, 11},
			IsValid: false,
		},
		{
			Name:    "Invalid: Two values are out of order",
			Report:  []int{1, 5, 3, 6, 5},
			IsValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if got := isValidReportWithError(tc.Report); got != tc.IsValid {
				t.Errorf("Got: %v, want: %v", got, tc.IsValid)
			}
		})
	}
}
