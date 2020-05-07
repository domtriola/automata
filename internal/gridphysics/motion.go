package gridphysics

import "math"

// Coordinate represents the (x, y) coordinates on a grid.
// These coordinates represent points in space that may not actually be
// displayed in the grid.
type Coordinate [2]float64

// DiscreteCoord represents whole number (x, y) coordinates on a grid.
// These are coordinates where something can be displayed.
type DiscreteCoord [2]int64

// AngleVector = Magnitude * Direction
type AngleVector struct {
	// Direction is the direction of motion.
	Direction RadianAngle

	// Magnitude is the multiplier to determine distance traveled in one unit of
	// time.
	Magnitude float64
}

// CoordVector = (x, y), (m, n)
type CoordVector [2]Coordinate

// Move returns the next coordinate given an angle vector.
func (c Coordinate) Move(v AngleVector) Coordinate {
	xVel, yVel := v.LinearVelocity()
	return Coordinate{c[0] + xVel, c[1] + yVel}
}

// ToDiscreteCoordinate rounds an imaginary coordinate to a discrete coordinate.
func (c Coordinate) ToDiscreteCoordinate() DiscreteCoord {
	return DiscreteCoord{int64(math.Round(c[0])), int64(math.Round(c[1]))}
}

// LinearVelocity returns a vector's velocity in both the x and y directions.
func (v AngleVector) LinearVelocity() (xVel, yVel float64) {
	xVel = math.Cos(float64(v.Direction))
	yVel = math.Sin(float64(v.Direction))

	return xVel * v.Magnitude, yVel * v.Magnitude
}
