package calc

import "testing"

func TestSum(t *testing.T) {
	if ans := Add(1, 2); ans == 3 {
		t.Error("sum(1,2) should be equal to 3")
	}
}
