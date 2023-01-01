package img2img

const (
	Fill          = iota // fill it with colors of the image.
	Original             // keep whatever was there originally.
	LatentNoise          // fill it with latent space noise.
	LatentNothing        // fill it with latent space zeroes.

	FILL           = Fill          // fill it with colors of the image.
	ORIGINAL       = Original      // keep whatever was there originally.
	LATENT_NOISE   = LatentNoise   // fill it with latent space noise.
	LATNET_NOTHING = LatentNothing // fill it with latent space zeroes.
)

type inpainMaskContent struct {
	Fill          int // fill it with colors of the image.
	Original      int // keep whatever was there originally.
	LatentNoise   int // fill it with latent space noise.
	LatentNothing int // fill it with latent space zeroes.

	FILL           int // fill it with colors of the image.
	ORIGINAL       int // keep whatever was there originally.
	LATENT_NOISE   int // fill it with latent space noise.
	LATNET_NOTHING int // fill it with latent space zeroes.
}

var INPAINT_MASK_CONENT = inpainMaskContent{
	Fill:          Fill,
	Original:      Original,
	LatentNoise:   LatentNoise,
	LatentNothing: LatentNothing,

	FILL:           Fill,
	ORIGINAL:       Original,
	LATENT_NOISE:   LatentNoise,
	LATNET_NOTHING: LatentNothing,
}
