package progress

// Api defines the common api between circular and linear progress components.
type Api interface {
	Determinate(bool)
	Open()
	Close()
	SetProgress(float64)
}
