package img2img

type inpaintMode struct {
	INPAINT_MASKED_AREA         int // Inpaint the masked area.
	INPAINT_OUTSIDE_MASKED_AREA int // Inpaint outside the masked area/everything except for masked area.
}

var INPAINT_MODE = inpaintMode{
	INPAINT_MASKED_AREA:         0,
	INPAINT_OUTSIDE_MASKED_AREA: 1,
}
