package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type SidesNumType int 
// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

var SidesTriangle SidesNumType = 3
var SidesSquare SidesNumType = 4
var SidesCircle SidesNumType = 0
var area float
//const Pi float32 = 3.14

func CalcSquare(sideLen float64, sidesNum SidesNumType) float64 {
	if sidesNum == 3 {
		area = Sqrt(3)/4*sideLen*sideLen
	} else if sidesNum == 4 {
		area = sideLen*sideLen
	} else if sidesNum == 0 {
		area = Pi*sideLen*sideLen
	} else {
		area = 0
	}
	return area
}
