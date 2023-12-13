package physics

import "fmt"

type Rect struct {
	Body *Body
	// use proper const
	shape                                   string
	width, height                           int
	originalVertices, vertices, normalFaces Matrix2d
	normalFaceToVertices                    map[int]int
	edgePoint                               V
}

func (r *Rect) Sides() int {
	return 4
}

func NewRectangleFromCorners(topLeft, bottomRight V, boundingRadius int, os ...Option) *Rect {
	width := bottomRight.X() - topLeft.X()
	height := bottomRight.Y() - topLeft.Y()
	center := New(width/2, height/2).Add(topLeft)

	if width <= 0 || height <= 0 {
		return nil
	}

	return NewRectangle(center, int(width), int(height), boundingRadius, os...)
}

func NewRectangle(center V, width, height, boundingRadius int, os ...Option) *Rect {
	// todo set bounding to half the diagonal: Math.sqrt(width*width + height*height)/2;
	// Make boundingRadius an option
	// make angle an option,
	// use defaults
	r := &Rect{
		Body:                 NewRigidBody(center, boundingRadius, os...),
		shape:                ShapeRectangle,
		width:                width,
		height:               height,
		normalFaceToVertices: map[int]int{},
	}

	r.calculateShapeEdges()

	return r
}

func (r *Rect) Radiuses() []*LineSegment {
	radiues := []*LineSegment{}
	for _, vertex := range r.vertices {
		radiues = append(radiues, NewLineSegment(r.Location(), vertex))
	}

	return radiues
}

func (r *Rect) Edges() []*LineSegment {
	vertices := r.vertices
	lastIndex := len(vertices) - 1
	edges := []*LineSegment{}
	for i := 0; i < len(vertices); i++ {
		nextI := i + 1
		if i == lastIndex {
			nextI = 0
		}

		edges = append(edges, NewLineSegment(vertices[i], vertices[nextI]))
	}

	return edges
}

func (r *Rect) Accelerate(incrementBy float64) {
	r.Body.Accelerate(incrementBy)
}

func (r *Rect) Drag(by float64) {
	r.Body.Drag(by)
}

func (r *Rect) RotatedEdgePoint() V {
	edgePoint := r.Body.center.Add(New(0, -float64(r.height/2)))
	edgePoint = edgePoint.Rotate(r.Body.center, r.Body.angle)

	return edgePoint
}

func (r *Rect) Acceleration() float64 {
	return r.Body.Acceleration()
}

func (r *Rect) Faces() []V {
	return r.normalFaces
}

func (r *Rect) Width() int {
	return r.width
}

func (r *Rect) Height() int {
	return r.height
}

func (r *Rect) TopLeft() V {
	topLeftNoRotation, _, _, _ := r.calculateCorners()

	return topLeftNoRotation.Rotate(r.Body.center, r.Body.angle)
}

func (r *Rect) TopRight() V {
	_, rightNoRotation, _, _ := r.calculateCorners()

	return rightNoRotation.Rotate(r.Body.center, r.Body.angle)
}

func (r *Rect) BottomLeft() V {
	_, _, _, bottomLeft := r.calculateCorners()

	return bottomLeft.Rotate(r.Body.center, r.Body.angle)
}

func (r *Rect) BottomRight() V {
	_, _, bottomRight, _ := r.calculateCorners()

	return bottomRight.Rotate(r.Body.center, r.Body.angle)
}

func (r *Rect) calculateCorners() (V, V, V, V) {
	center := r.Body.center

	halfWidth := float64(r.width) / 2
	halfHeight := float64(r.height) / 2

	centerXHalf := center.X() - halfWidth
	centerYHalf := center.Y() - halfHeight
	centerXAddHalf := center.X() + halfWidth
	centerYAddHalf := center.Y() + halfHeight

	topLeft := New(centerXHalf, centerYHalf)
	topRight := New(centerXAddHalf, centerYHalf)
	bottomRight := New(centerXAddHalf, centerYAddHalf)
	bottomLeft := New(centerXHalf, centerYAddHalf)

	return topLeft, topRight, bottomRight, bottomLeft
}

func (r *Rect) setFaces() {
	if len(r.vertices) == 0 {
		fmt.Printf("Error - vertices not set")
		return
	}
	topLeft, topRight, bottomRight, bottomLeft := r.vertices[0], r.vertices[1], r.vertices[2], r.vertices[3]

	faces := make([]V, 4)
	faces[0] = topRight.Sub(bottomRight).Normalize()
	r.normalFaceToVertices[0] = 0
	faces[1] = bottomRight.Sub(bottomLeft).Normalize()
	r.normalFaceToVertices[1] = 1
	faces[2] = bottomLeft.Sub(topLeft).Normalize()
	r.normalFaceToVertices[2] = 2
	faces[3] = topLeft.Sub(topRight).Normalize()
	r.normalFaceToVertices[3] = 3

	r.normalFaces = faces
}

func (r *Rect) GetVerticesFromNormalFace(v V) *V {
	for i, face := range r.normalFaces {
		if face.Equal(v) {
			return &r.vertices[r.normalFaceToVertices[i]]
		}
	}

	return nil
}

// todo: major Apply and rigid body need to
// add a moveby slice which is applied in the update to rigid body first then vertices, reclaculate the vertices
func (r *Rect) Move(by ...V) {
	r.Body.Move(by...)
}

func (r *Rect) Vertices() []V {
	return r.vertices
}

func (r *Rect) Location() V {
	return r.Body.center
}

func (r *Rect) Update() (bool, V) {
	if r.Body.fixed {
		return false, ZeroVector
	}

	hasMoved, movedBy := r.Body.Update()
	r.originalVertices = r.originalVertices.Add(movedBy)
	r.vertices = r.vertices.Add(movedBy)

	// transpose need
	r.vertices = r.originalVertices.Rotate(r.Body.center, r.Body.angle)
	r.setFaces()

	return hasMoved, movedBy
}

// Scale This was not from the book I have gone rogue here but it works
// This could be used for animations and explosiions
func (r *Rect) Scale(x float64) {
	oldCenter := r.Body.center
	origin := V{0, 0}
	toOrigin := origin.Sub(oldCenter)

	r.Move(toOrigin)
	// todo this doesn't do anything
	r.Body.center.Scale(x)
	r.vertices = r.vertices.Map(func(v V) V {
		return v.Scale(x)
	})
	r.originalVertices = r.originalVertices.Map(func(v V) V {
		return v.Scale(x)
	})
	//
	r.Move(oldCenter)
}

func (r *Rect) Transpose(to V) {
	r.Body.Transpose(to)
	r.calculateShapeEdges()
}

func (r *Rect) calculateShapeEdges() {
	topLeft, topRight, bottomRight, bottomLeft := r.calculateCorners()
	originalVertices := NewMatrix2d(topLeft, topRight, bottomRight, bottomLeft)
	r.originalVertices = originalVertices
	r.vertices = originalVertices
	r.setFaces()
}
