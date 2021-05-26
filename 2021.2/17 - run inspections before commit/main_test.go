package inspections_before_commit

import "testing"

func TestBeforeCommit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"OK"},
		{"FAIL"},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "FAIL" {
				t.Fatal("test failed!")
			}
		})
	}
}

