package earth

// Position est the geometric characteristic of an object on the earth
type Position struct {
	Loc VectorA
	Cap VectorA
}

// MoveObject move an object on the earth
func MoveObject(position Position, speed VectorL, deltaTime float64) (newPosition Position) {
	lon := position.Loc.Lon
	length := VectorL{
		U: speed.U * deltaTime,
		V: speed.V * deltaTime,
	}
	angle := Length2Angle(length, lon)
	newLoc := VectorA{
		Lon: position.Loc.Lon + angle.Lon,
		Lat: position.Loc.Lat + angle.Lat,
	}
	newPosition = Position{
		newLoc,
		position.Cap,
	}
	return
}
