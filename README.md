# go-waveform
<p align="center">
	<img src="image.svg" alt="go-waveform">
</p>

Golang project for the purpose of practicing, understanding golang and audio file format.

## Usage

## Requirements

Go 1.11 (Go Modules)

### Cli

Generate waveform image via CLI

```bash
GO111MODULE=on make
./bin/go-waveform -t dark -r 1000 -f png example/violin.wav

open example/violin-L.png
open example/violin-R.png
```

### Import as package

Import as package,

```go
import "github.com/cettoana/go-waveform"
```

and decode wav audio file

```go
w := waveform.DecodeWav(wavFileBytes)
fmt.Println(w.BitsPerSample)

data, _ := w.GetData()

if stereoData, ok := data.(*waveform.StereoData); ok {
  fmt.Println(stereoData.RSample)
  fmt.Println(stereoData.LSample)
} else if monoData, ok := data.(*waveform.MonoData); ok {
  fmt.Println(monoData.Sample)
}
```

### Todo

- More wav file details
- Converter
- Raw Waveform
- Waveform image customization
- Test
