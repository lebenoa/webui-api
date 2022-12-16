package img2img

type resizeMode struct {
	JUST_RESIZE     int // Resize image to target resolution. Unless height and width match, you will get incorrect aspect ratio.
	CROP_AND_RESIZE int // Resize the image so that entirety of target resolution is filled with the image. Crop parts that stick out.
	RESIZE_AND_FILL int // Resize the image so that entirety of image is inside target resolution. Fill empty space with image's colors.
}

var RESIZE_MODE = resizeMode{
	JUST_RESIZE:     0,
	CROP_AND_RESIZE: 1,
	RESIZE_AND_FILL: 2,
}
