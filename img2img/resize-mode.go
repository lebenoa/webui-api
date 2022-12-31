package img2img

const (
	JustResize    = iota // Resize image to target resolution. Unless height and width match, you will get incorrect aspect ratio.
	CropAndResize        // Resize the image so that entirety of target resolution is filled with the image. Crop parts that stick out.
	ResizeAndFill        // Resize the image so that entirety of image is inside target resolution. Fill empty space with image's colors.

	JUST_RESIZE     = JustResize    // Resize image to target resolution. Unless height and width match, you will get incorrect aspect ratio.
	CROP_AND_RESIZE = CropAndResize // Resize the image so that entirety of target resolution is filled with the image. Crop parts that stick out.
	RESIZE_AND_FILL = ResizeAndFill // Resize the image so that entirety of image is inside target resolution. Fill empty space with image's colors.
)

type resizeMode struct {
	JustResize    int // Resize image to target resolution. Unless height and width match, you will get incorrect aspect ratio.
	CropAndResize int // Resize the image so that entirety of target resolution is filled with the image. Crop parts that stick out.
	ResizeAndFill int // Resize the image so that entirety of image is inside target resolution. Fill empty space with image's colors.

	JUST_RESIZE     int // Resize image to target resolution. Unless height and width match, you will get incorrect aspect ratio.
	CROP_AND_RESIZE int // Resize the image so that entirety of target resolution is filled with the image. Crop parts that stick out.
	RESIZE_AND_FILL int // Resize the image so that entirety of image is inside target resolution. Fill empty space with image's colors.
}

var RESIZE_MODE = resizeMode{
	JustResize:    JustResize,
	CropAndResize: CropAndResize,
	ResizeAndFill: ResizeAndFill,

	JUST_RESIZE:     JustResize,
	CROP_AND_RESIZE: CropAndResize,
	RESIZE_AND_FILL: ResizeAndFill,
}
