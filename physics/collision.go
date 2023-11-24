package physics

type Collision struct {
	//beginning of overlap
	start V
	// end of the overlap
	end V
	// direction of depth
	normal V
	// smallest amount objects have crossed over
	depth float64 // maybe int
}

func NewCollision(start, normal V, depth float64) *Collision {
	return &Collision{
		start:  start,
		end:    start.Add(normal.Scale(depth)),
		normal: normal,
		depth:  depth,
	}
}

func NewCollisionFromTouch(touch Touch) *Collision {
	if touch.Empty() {
		return nil
	}

	return NewCollision(touch.Vector(), touch.LineAEndToOffSet().Direction().Invert().Normalize(), touch.LineAEndToOffSet().Magnitude())
}

func (c *Collision) Start() V {
	return c.start
}

func (c *Collision) End() V {
	return c.end
}

func (c *Collision) Normal() V {
	return c.normal
}

func (c *Collision) Depth() float64 {
	return c.depth
}

func (c *Collision) Resolve() V {
	return c.Normal().Invert().Scale(c.Depth())
}

func (c *Collision) ReverseDirection() *Collision {
	start := c.start
	return &Collision{
		normal: c.normal.Scale(-1),
		start:  c.end,
		end:    start,
		depth:  c.depth,
	}
}
