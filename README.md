# AUTOMATIC1111's Webui API

AUTOMATIC1111's Webui **API Wrapper** for **GO**. So you don't have to do it yourself.   
Aim to be as easy to use as possible ***without*** performance in mind.  

## Currently Support (And also roadmap)

- [ ] Auth Related
- [x] Txt2Img
- [x] Img2Img
- [x] Extras (Single)
- [x] Extras (Batch)
- [x] PNG Info
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
        Txt2Img:     "/sdapi/v1/txt2img",
        Img2Img:     "/sdapi/v1/img2img",
        ExtraSingle: "/sdapi/v1/extra-single-image",
        ExtraBatch:  "/sdapi/v1/extra-batch-images",
        PNGInfo:     "/sdapi/v1/png-info",
        Progress:    "/sdapi/v1/progress",
        Interrogate: "/sdapi/v1/interrogate",
        Interrupt:   "/sdapi/v1/interrupt",
        Skip:        "/sdapi/v1/skip",
        SDModels:    "/sdapi/v1/sd-models",
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
