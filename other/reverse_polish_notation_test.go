package other

import "testing"

func Test_evaluateRPN(t *testing.T) {
	expression := "5 3 + 4 *"

	rpn, err := evaluateRPN(expression)
	if err != nil {
		t.Errorf("Didn't expect error: %s", err)
	}

	rslt := (5 + 3) * 4

	if rpn != rslt {
		t.Errorf("expected to get %d, but got %d", rslt, rpn)
	}
}
