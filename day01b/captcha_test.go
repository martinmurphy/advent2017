package day01b

import "testing"

func TestEvalCaptchaMid(t *testing.T) {
	testValues := []struct {
		captcha string
		result  int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for i, testValue := range testValues {
		result := EvalCaptchaMid(testValue.captcha)
		if result != testValue.result {
			t.Errorf("Eval of captcha[%d]: \"%s\" was incorrect, expected: %d, actual: %d.", i, testValue.captcha, testValue.result, result)
		}
	}
}
