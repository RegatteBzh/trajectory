package earth

// Position est the geometric characteristic of an object on the earth
type Position struct {
	Loc VectorA
	Cap VectorA
}

// MoveObject move an object on the earth
func MoveObject(position Position, speed VectorL, deltaTime float64) (newPosition Position) {
	lambda := position.Loc.Lambda
	length := VectorL{
		speed.U * deltaTime,
		speed.V * deltaTime,
	}
	angle := Length2Angle(length, lambda)
	newLoc := VectorA{
		position.Loc.Lambda + angle.Lambda,
		position.Loc.Phi + angle.Phi,
	}
	newPosition = Position{
		newLoc,
		position.Cap,
	}
	return
}
