package interrogate

const (
	CLIP      = "clip"         // Human style
	DEEPBOORU = "deepdanbooru" // Booru tags style
)

type model struct {
	CLIP      string // Human style
	DEEPBOORU string // Booru tags style
}

var Model = model{
	CLIP:      CLIP,
	DEEPBOORU: DEEPBOORU,
}
