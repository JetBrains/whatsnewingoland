package cleanup_before_commit

import "testing"

func TestBeforeCommit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"OK"},
		{"FAIL"},
	}
	var message = "test failed!"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "FAIL" {
				t.Fatal(message)
			}
		})
	}
}
