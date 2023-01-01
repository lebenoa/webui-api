package img2img

const (
	InpaintMaskedArea        = iota // Inpaint the masked area.
	InpaintOutsideMaskedArea        // Inpaint outside the masked area/everything except for masked area.

	INPAINT_MASKED_AREA         = InpaintMaskedArea        // Inpaint the masked area.
	INPAINT_OUTSIDE_MASKED_AREA = InpaintOutsideMaskedArea // Inpaint outside the masked area/everything except for masked area.
)

type inpaintMode struct {
	InpaintMaskedArea        int // Inpaint the masked area.
	InpaintOutsideMaskedArea int // Inpaint outside the masked area/everything except for masked area.

	INPAINT_MASKED_AREA         int // Inpaint the masked area.
	INPAINT_OUTSIDE_MASKED_AREA int // Inpaint outside the masked area/everything except for masked area.
}

var INPAINT_MODE = inpaintMode{
	InpaintMaskedArea:        InpaintMaskedArea,
	InpaintOutsideMaskedArea: InpaintOutsideMaskedArea,

	INPAINT_MASKED_AREA:         InpaintMaskedArea,
	INPAINT_OUTSIDE_MASKED_AREA: InpaintOutsideMaskedArea,
}
