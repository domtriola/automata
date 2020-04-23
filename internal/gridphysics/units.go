package gridphysics

import "math"

// DegreeAngle is an angle in degrees
type DegreeAngle float64

// RadianAngle is an angle in radians
type RadianAngle float64

// ToRadians converts degrees to radians
func (d DegreeAngle) ToRadians() (rad RadianAngle) {
	return RadianAngle(d * math.Pi / 180)
}
