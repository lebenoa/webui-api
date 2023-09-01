# This project is unmaintained. Looking for a new owner or you should use fork instead.

I'm not using GO anymore. While GO is definitely fun to use, it's "a bit **verbose**" and   
I think Rust is more enjoyable with something like `Option<T>` enum.

GO is a great language and maybe one of the first "***Errors are values***" and even with garbage collector, it is pretty fast.  
but for me, Rust is just more powerful with all these macro stuff  
(I can't write/create one obviously but I can't express enough how easy it is to use library like Tauri, Serde, poise-rs)

# AUTOMATIC1111's Webui API

AUTOMATIC1111's Webui **API** for **GO**. So you don't have to do it yourself.   
Aim to be as easy to use as possible ***without*** performance in mind.  

## Currently Support (And also roadmap)

- [x] Auth Related ( **DIDNT TEST** | Please [open an issue](https://github.com/Meonako/webui-api/issues/new) if you have encounter any problem )
- [x] Txt2Img
- [x] Img2Img
- [x] Extras (Single)
- [x] Extras (Batch)
- [x] PNG Info
- [x] Progress
- [x] Interrogate
- [x] Interrupt
- [x] Skip
- [x] Options
- [x] Get Available Model(s)
 
## Getting Started

***Required [Stable Diffusion Web UI](https://github.com/AUTOMATIC1111/stable-diffusion-webui) running with `--api` argument***

```
go get github.com/Meonako/webui-api
```
Then Import
```go
import (
    ...
    "github.com/Meonako/webui-api"
)
```

_OR_

Simply add package to import like this
```go
import (
    ...
    "github.com/Meonako/webui-api"
)
```

Then run `go mod tidy`

---

Initialize it like this

```go
API := api.New()
```

Without passing anything it'll return `http://127.0.0.1:7860` and all `V1 API path` as default.  
If you wanna change it, just pass in a new config like this

```go
API := api.New(api.Config{
    BaseURL: "colab.google.link",
    Path: &api.APIPath{
        Txt2Img: "/new/api/path/txt2img",
    },
})
```
**Be aware, if you change `Path` field, you'll have to manually add all other path.**  
> Say for above example, it'll be only `Txt2Img` path in there. When you call `Img2Img`, you'll get an `unexpected response`/`error` or worse like `panic`

---

Now that finished, we can start using it now. Let's say we'll do `TXT2IMG`
```go
resp, err := API.Text2Image(&api.Txt2Image{
    Prompt: "masterpiece, best quality, solo, cute, blue hair, purple eyes",
    NegativePrompt: "lowres, bad anatomy, low quality, normal quality, worst quality",
})
```
> **Keep in mind that this will block your app until API done generating image(s)**

---

When it's done, check for the `error` and then we can do

```go
imageList, err := resp.DecodeAllImages()
```
Check for the `error` and then we can save it to disk

```go
for index, image := range imageList {
    file, err := os.OpenFile(
        fmt.Sprintf("txt2img-result %v.png", index), 
        os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
        0777,
    )
    if err != nil {
        panic(err)
    }
    
    file.Write(image)
    file.Close()
}
```

**Hol'up, one tip tho. If you really care about performance, you can decode it yourself like this**  
> Before called `resp.DecodeAllImages()`

```go
for index := range resp.Images {
    decoded, err := resp.DecodeImage(index)
    if err != nil {
        panic(err)
    }
    
    file, err := os.OpenFile(
        fmt.Sprintf("txt2img-result %v.png", index), 
        os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
        0777,
    )
    if err != nil {
        panic(err)
    }
    
    file.Write(image)
    file.Close()
}
```

## Example

Move [HERE](https://github.com/Meonako/webui-api/wiki/Example)

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
        Options:     "/sdapi/v1/options",
        SDModels:    "/sdapi/v1/sd-models",
    },
}
```

## Credits
- [go-json](https://github.com/goccy/go-json) / [goccy](https://github.com/goccy) | for fast JSON encode/decode
