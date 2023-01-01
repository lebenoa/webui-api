package interrogate

const (
	CLIP      = "clip"
	DEEPBOORU = "deepbooru"
)

type model struct {
	CLIP      string
	DEEPBOORU string
}

var Model = model{
	CLIP:      CLIP,
	DEEPBOORU: DEEPBOORU,
}
