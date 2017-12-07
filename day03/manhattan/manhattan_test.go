package manhattan

import "testing"

// sample grid
//
// 17  16  15  14  13
// 18   5   4   3  12
// 19   6   1   2  11
// 20   7   8   9  10
// 21  22  23---> ...

func TestDistance(t *testing.T) {
	testValues := []struct {
		location int
		distance int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{12, 3},
		{23, 2},
		{25, 4},
		{48, 5},
		{49, 6},

		{1024, 31},
	}

	for i, testValue := range testValues {
		result := Distance(testValue.location)
		if result != testValue.distance {
			t.Errorf("Distance from location[%d]: \"%d\" was incorrect, expected: %d, actual: %d.", i, testValue.location, testValue.distance, result)
		}
	}
}

// sample grid
// 37 36  35  34  35  32 31
// 38 17  16  15  14  13 30
// 39 18   5   4   3  12 29
// 40 19   6   1   2  11 28
// 41 20   7   8   9  10 27
// 42 21  22  23  24  25 26
// 43 44  45  46  47  48 49
func TestCoordsOf1(t *testing.T) {
	testValues := []struct {
		size float64
		x    float64
		y    float64
	}{
		{1, 0, 0},
		{2, 0, 1},
		{3, 1, 1},
		{4, 1, 2},
		{5, 2, 2},
		{6, 2, 3},
		{7, 3, 3},
	}

	for i, testValue := range testValues {
		result_x, result_y := getCoordsOf1(testValue.size)
		if (result_x != testValue.x) || (result_y != testValue.y) {
			t.Errorf("coords of 1 [%d] in square of size: \"%f\" was incorrect, expected: (%f, %f), actual: (%f, %f).", i, testValue.size, testValue.x, testValue.y, result_x, result_y)
		}
	}
}

func TestCoordsOfLocation(t *testing.T) {
	// sample grid size 7             // sample grid size 6
	// 37 36  35  34  35  32 31       // 36  35  34  35  32 31
	// 38 17  16  15  14  13 30       // 17  16  15  14  13 30
	// 39 18   5   4   3  12 29       // 18   5   4   3  12 29
	// 40 19   6   1   2  11 28       // 19   6   1   2  11 28
	// 41 20   7   8   9  10 27       // 20   7   8   9  10 27
	// 42 21  22  23  24  25 26       // 21  22  23  24  25 26
	// 43 44  45  46  47  48 49
	testValues := []struct {
		size     float64
		location float64
		x        float64
		y        float64
	}{
		{6, 26, 5, 5},
		{6, 27, 5, 4},
		{6, 28, 5, 3},
		{6, 29, 5, 2},
		{6, 30, 5, 1},
		{6, 31, 5, 0},
		{6, 32, 4, 0},
		{6, 33, 3, 0},
		{6, 34, 2, 0},
		{6, 35, 1, 0},
		{6, 36, 0, 0},
		{7, 37, 0, 0},
		{7, 38, 0, 1},
		{7, 39, 0, 2},
		{7, 40, 0, 3},
		{7, 41, 0, 4},
		{7, 42, 0, 5},
		{7, 43, 0, 6},
		{7, 44, 1, 6},
		{7, 45, 2, 6},
		{7, 46, 3, 6},
		{7, 47, 4, 6},
		{7, 48, 5, 6},
		{7, 49, 6, 6},
	}

	for i, testValue := range testValues {
		result_x, result_y := getCoordsOfLocation(testValue.size, testValue.location)
		if (result_x != testValue.x) || (result_y != testValue.y) {
			t.Errorf("coords of location [%d] : \"%f\" was incorrect, expected: (%f, %f), actual: (%f, %f).", i, testValue.location, testValue.x, testValue.y, result_x, result_y)
		}
	}
}
