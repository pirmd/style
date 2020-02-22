package style

import (
	"testing"
)

func TestMandocEscaping(t *testing.T) {
	st := NewMan()

	testCases := []struct {
		in  string
		out string
	}{
		{"Test -- test is fun", "Test \\-\\- test is fun"},
		{"Test \\-\\- test is fun", "Test \\-\\- test is fun"},
		{"Test -- test is \\fBfun\\fP", "Test \\-\\- test is \\fBfun\\fP"},
		{"Test a backslash \\\\", "Test a backslash \\\\"},
	}

	for _, tc := range testCases {
		got := st.Escape(tc.in)
		if got != tc.out {
			t.Errorf("escape man: '%s' failed. Want: %s. Got : %s", tc.in, tc.out, got)
		}

		gotgot := st.Escape(got)
		if gotgot != tc.out {
			t.Errorf("escape man (second pass): '%s' failed. Want: %s. Got : %s", tc.in, tc.out, gotgot)
		}
	}
}
