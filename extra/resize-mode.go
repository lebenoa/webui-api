package extra

type resizeMode struct {
	SCALE_BY int // Multiplier to width and height of the original image
	SCALE_TO int // Scale to specify width and height
}

var ResizeMode = resizeMode{
	SCALE_BY: ScaleBy,
	SCALE_TO: ScaleTo,
}
