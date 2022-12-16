# AUTOMATIC1111's Webui API

High level API written in GO.

## Currently Support (And also roadmap)

- [ ] Auth Related
- [x] Txt2Img
- [x] Img2Img
- [ ] Extras (Single)
- [ ] Extras (Batch)
- [ ] PNG Info
- [x] Progress
- [ ] Interrogate
- [ ] Interrupt
- [ ] Skip
- [ ] Options
- [x] Get Available Model(s)
 
## Getting Started

```
go get github.com/Meonako/webui-api
```
Then Import

---

_OR_

---

Simply add package to import like this
```go
import (
    ...
    "github.com/Meonako/webui-api"
)
```

Then run `go mod tidy`

## Default Value
```go
var DefaultConfig = Config{
    BaseURL: "http://127.0.0.1:7860",
    Path: &APIPath{
        // Don't change any of these unless you know what you're doing. 
        // I purposely exported this as I don't know If I'll still maintain this pkg in the future
        Txt2Img: "/sdapi/v1/txt2img",
        Progress: "/sdapi/v1/progress",
    },
    DefaultSampler:  sampler.EULER_A,
    DefaultSteps:    28,
    DefaultCFGScale: 7,
    DefaultWidth:    512,
    DefaultHeight:   512,
}
```

## Example

Move [HERE](https://github.com/Meonako/webui-api/wiki/Example)

## Credits
- [go-json](https://github.com/goccy/go-json) / [goccy](https://github.com/goccy) | for fast JSON encode/decode
