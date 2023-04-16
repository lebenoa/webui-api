package img2img

const (
	JUST_RESIZE     = iota // Resize image to target resolution. Unless height and width match, you will get incorrect aspect ratio.
	CROP_AND_RESIZE        // Resize the image so that entirety of target resolution is filled with the image. Crop parts that stick out.
	RESIZE_AND_FILL        // Resize the image so that entirety of image is inside target resolution. Fill empty space with image's colors.
)
