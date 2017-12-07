package day01a

import "testing"

func TestEvalCaptcha(t *testing.T) {
	testValues := []struct {
		captcha string
		result  int
	}{
		{"123", 0},
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
		{"1A23", 0},
		{"11b22", 3},
		{"1111c", 3},
		{"1234c", 0},
		{"91212c129", 9},
	}

	for i, testValue := range testValues {
		result := EvalCaptcha(testValue.captcha)
		if result != testValue.result {
			t.Errorf("Eval of captcha[%d]: \"%s\" was incorrect, expected: %d, actual: %d.", i, testValue.captcha, testValue.result, result)
		}
	}
}

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
