package manhattan

import (
	"math"
	"strconv"
)

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// )

// sample grid
//
// 17  16  15  14  13
// 18   5   4   3  12
// 19   6   1   2  11
// 20   7   8   9  10
// 21  22  23---> ...

// Distance calculates the distance from the passed location to location 1
func Distance(location int) int {
	distance := 0
	if location < 0 {
		distance = -1
	} else if location == 1 {
		distance = 0
	} else if location == 2 {
		distance = 1
	} else if location == 3 {
		distance = 2
	} else if location == 4 {
		distance = 1
	} else {
		squareSize := math.Floor(math.Sqrt(float64(location-1)) + 1)
		originX, originY := getCoordsOf1(squareSize)
		locationX, locationY := getCoordsOfLocation(squareSize, float64(location))
		distance = int(math.Abs(originX-locationX) + math.Abs(originY-locationY))
	}
	return distance
}

func getCoordsOf1(squareSize float64) (x, y float64) {
	x = math.Floor(squareSize/2) - (1 - math.Mod(squareSize, 2))
	y = math.Floor(squareSize / 2)
	return
}

func getCoordsOfLocation(squareSize float64, location float64) (x, y float64) {
	distanceAround := location - (squareSize-1)*(squareSize-1)
	if math.Mod(squareSize, 2) > 0 {
		if distanceAround >= 1 && distanceAround <= squareSize {
			x = 0
			y = distanceAround - 1
		} else if distanceAround > squareSize && distanceAround <= (squareSize+squareSize-1) {
			x = distanceAround - squareSize
			y = squareSize - 1
		}
	} else {
		if distanceAround >= 1 && distanceAround <= squareSize {
			x = squareSize - 1
			y = squareSize - distanceAround
		} else if distanceAround > squareSize && distanceAround <= (squareSize+squareSize-1) {
			x = squareSize - (distanceAround - squareSize) - 1
			y = 0
		}

	}
	return
}

type coord struct {
	x, y int
}

var m map[string]int

func init() {
	m = make(map[string]int)
}

func getKey(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func setter(x, y, value int) {
	m[getKey(x, y)] = value
}

func getter(x, y int) (value int) {
	value = m[getKey(x, y)]
	return
}

func isSet(x, y int) bool {
	return getter(x, y) != 0
}

var offsets = [4]coord{
	{0, -1},
	{-1, 0},
	{0, 1},
	{1, 0},
}

func leftDir(dir int) (newdir int) {
	return (dir + 1) % 4
}

func leftCoords(x, y, dir int) (newx, newy int) {
	newx = x + offsets[leftDir(dir)].x
	newy = y + offsets[leftDir(dir)].y
	return
}

func straightCoords(x, y, dir int) (newx, newy int) {
	newx = x + offsets[dir].x
	newy = y + offsets[dir].y
	return
}

func coordsAround(x, y int) (c []coord) {
	c = make([]coord, 8, 8)
	c[0].x = x
	c[0].y = y - 1
	c[1].x = x - 1
	c[1].y = y - 1
	c[2].x = x - 1
	c[2].y = y
	c[3].x = x - 1
	c[3].y = y + 1
	c[4].x = x
	c[4].y = y + 1
	c[5].x = x + 1
	c[5].y = y + 1
	c[6].x = x + 1
	c[6].y = y
	c[7].x = x + 1
	c[7].y = y - 1
	return
}

func valuesAround(x, y int) (v []int) {
	coords := coordsAround(x, y)
	v = make([]int, 8, 8)
	for i, coord := range coords {
		v[i] = getter(coord.x, coord.y)
	}
	return
}

func sumValuesAround(x, y int) (value int) {
	value = 0
	for _, v := range valuesAround(x, y) {
		value += v
	}
	return
}

func nextMove(x, y, dir int) (newX, newY, newDir int) {
	tryLeftX, tryLeftY := leftCoords(x, y, dir)
	if !isSet(tryLeftX, tryLeftY) {
		newX = tryLeftX
		newY = tryLeftY
		newDir = leftDir(dir)
	} else {
		newX, newY = straightCoords(x, y, dir)
		newDir = dir
	}
	return
}

func GoUntilGreaterValue(v int) (value, x, y, dir int) {
	x, y, dir = 0, 0, 2
	value = 1
	setter(x, y, value)
	for value <= v {
		x, y, dir = nextMove(x, y, dir)
		value = sumValuesAround(x, y)
		setter(x, y, value)
	}
	return
}
