package checksum

import "testing"

func TestEvalChecksum(t *testing.T) {
	testValues := []struct {
		spreadsheet string
		checksum    int
	}{
		{"1\t1\t1\n2\t2\t2\n3\t3\t3\n", 0},
		{
			`1	2	3
4	5	6
7	8	9`, 6},
		{"1\t4\t7\n2\t5\t8\n3\t6\t9\n", 18},
		{"5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n", 18},
	}

	for i, testValue := range testValues {
		result := EvalChecksum(testValue.spreadsheet)
		if result != testValue.checksum {
			t.Errorf("Eval of captcha[%d]: \"%s\" was incorrect, expected: %d, actual: %d.", i, testValue.spreadsheet, testValue.checksum, result)
		}
	}
}

func TestEvalChecksum2(t *testing.T) {
	testValues := []struct {
		spreadsheet string
		checksum    int
	}{
		{`5	9	2	8
9	4	7	3
3	8	6	5`, 9},
	}

	for i, testValue := range testValues {
		result := EvalChecksum2(testValue.spreadsheet)
		if result != testValue.checksum {
			t.Errorf("Eval of captcha[%d]: \"%s\" was incorrect, expected: %d, actual: %d.", i, testValue.spreadsheet, testValue.checksum, result)
		}
	}
}
