package extra

type resizeMode struct {
	SCALE_BY int // Multiplier to width and height of the original image
	SCALE_TO int //
}

var ResizeMode = resizeMode{
	SCALE_BY: ScaleBy,
	SCALE_TO: ScaleTo,
}
