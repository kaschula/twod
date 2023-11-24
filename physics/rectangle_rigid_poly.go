package physics

func (r *Rect) Rotate(byDegrees Degree) {
	r.Body.Rotate(byDegrees)
}

func (r *Rect) SetRotation(degrees Degree) {
	r.Body.SetRotation(degrees)
}

func (r *Rect) RotateTo(point V) {
	r.Body.RotateTo(point)
}

func (r *Rect) PreviousAngle() V {
	return r.Body.PreviousAngle()
}

func (r *Rect) Direction() Degree {
	return r.Body.Direction()
}
