# AUTOMATIC1111's Webui API

High level API written in GO.

## Currently Support (And also roadmap)

- [ ] Auth Related
- [x] Txt2Img
- [ ] Img2Img
- [ ] Extras (Single)
- [ ] Extras (Batch)
- [ ] PNG Info
- [x] Progress
- [ ] Interrogate
- [ ] Interrupt
- [ ] Skip
- [ ] Options
- [ ] Get Available Model(s)
 
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

```go
package main

import (
	"fmt"
	"image/png"
	"os"
	"github.com/Meonako/webui-api"
	"text/tabwriter"
)

func main() {
	API := api.New(api.Config{
		UseDefault: true,
	})

	resp, err := API.Text2Image(api.Txt2Image{
		Prompt: api.BuildPrompt(
			"solo", "cute", "adorable", "innocent",
			"blush", "girly", "narrow waist", "absurdly long hair",
			"(pink hair)", "hair ornament", "pink eyes", "hair flower",
		),
		//
		// Or you can pass a long string
		//
		// Prompt: "solo, cute, adorable, innocent, blush, girly, narrow waist, absurdly long hair, (pink hair), hair ornament, pink eyes, hair flower",
		//
		// Or you can do this
		//
		// Prompt: "solo, cute, adorable, innocent" +
		// "blush, girly, narrow waist, absurdly long hair" +
		// "(pink hair), hair ornament, pink eyes, hair flower",
		//
		NegativePrompt: api.BuildPrompt(
			"sketch by bad-artist", "lowres", "bad anatomy",
			"bad hands", "text", "error", "missing fingers",
			"extra digit", "fewer digits", "cropped",
			"worst quality", "low quality", "normal quality",
			"jpeg artifacts", "signature", "watermark",
			"username", "blurry", "simple background",
			"(track suit)", "(jacket)", "(name tag)",
			"(sleeveless), (shoes), (socks), hat",
			"covered navel, clothes lift, clothes pull, (head out of frame)",
		),
	})
	if err != nil {
		panic(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	fmt.Fprintf(writer, "Steps\t|\t%v\n", resp.Parameters.Steps)
	fmt.Fprintf(writer, "Sampler Name\t|\t%v\n", resp.Parameters.SamplerName)
	fmt.Fprintf(writer, "Sampler Index\t|\t%v\n", resp.Parameters.SamplerIndex)
	fmt.Fprintf(writer, "Width\t|\t%v\n", resp.Parameters.Width)
	fmt.Fprintf(writer, "Height\t|\t%v\n", resp.Parameters.Height)
	writer.Flush()

	readers, err := resp.MakeBytesReader()
	if err != nil {
		panic(err)
	}

	for index, reader := range readers {
		file, err := os.OpenFile(fmt.Sprintf("Image (%v).png", index), os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		img, err := png.Decode(reader)
		if err != nil {
			panic(err)
		}

		err = png.Encode(file, img)
		if err != nil {
			panic(err)
		}
	}
}

```

The output should be like this

![image](https://user-images.githubusercontent.com/76484203/207892808-dbe685d6-5933-4cf1-a925-b5dd2797b407.png)

and you'll have file named "Image (0).png" in your root folder

## Credits
- [go-json](https://github.com/goccy/go-json) / [goccy](https://github.com/goccy) | for fast JSON encode/decode
