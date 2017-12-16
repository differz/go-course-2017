package shapes

func init() {
	Register("cone", NewCone)
	Register("cube", NewCube)
	Register("cylinder", NewCylinder)
	Register("prizm", NewPrizm)
	Register("pyramid", NewPyramid)
	Register("sphere", NewSphere)
}
