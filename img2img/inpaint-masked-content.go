package img2img

type inpainMaskContent struct {
	FILL           int // fill it with colors of the image.
	ORIGINAL       int // keep whatever was there originally.
	LATENT_NOISE   int // fill it with latent space noise.
	LATNET_NOTHING int // fill it with latent space zeroes.
}

var INPAINT_MASK_CONENT = inpainMaskContent{
	FILL:           0,
	ORIGINAL:       1,
	LATENT_NOISE:   2,
	LATNET_NOTHING: 3,
}
