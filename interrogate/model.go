package interrogate

const (
	CLIP      = "clip"
	DEEPBOORU = "deepdanbooru"
)

type model struct {
	CLIP      string
	DEEPBOORU string
}

var Model = model{
	CLIP:      CLIP,
	DEEPBOORU: DEEPBOORU,
}
