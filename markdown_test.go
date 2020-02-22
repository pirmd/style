package style

import (
	"testing"
)

func TestMkdTextEscaping(t *testing.T) {
	st := NewMarkdown()

	testCases := []struct {
		in  string
		out string
	}{
		{"*Test* is _fun_!", "\\*Test\\* is \\_fun\\_\\!"},
		{"**Test** is [really] _fun_!", "\\*\\*Test\\*\\* is \\[really\\] \\_fun\\_\\!"},
		{"Test a backslash \\", "Test a backslash \\\\"},
	}

	for _, tc := range testCases {
		got := st.Escape(tc.in)
		if got != tc.out {
			t.Errorf("escape markdown: '%s' failed. Want: %s. Got : %s", tc.in, tc.out, got)
		}

		gotgot := st.Escape(tc.in)
		if gotgot != tc.out {
			t.Errorf("escape markdown (second pass): '%s' failed. Want: %s. Got : %s", tc.in, tc.out, gotgot)
		}
	}
}
