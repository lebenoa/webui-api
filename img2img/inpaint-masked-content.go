package img2img

const (
	FILL           = iota // fill it with colors of the image.
	ORIGINAL              // keep whatever was there originally.
	LATENT_NOISE          // fill it with latent space noise.
	LATNET_NOTHING        // fill it with latent space zeroes.
)
