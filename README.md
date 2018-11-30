# go-waveform
<p align="center">
	<img src="image.svg" alt="go-waveform">
</p>

Golang project for the purpose of practicing, understanding golang and audio file format.

## Usage

### Cli

Generate waveform image via CLI

```bash
make
./bin/go-waveform -r 1000 -f png example/violin.wav

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
w := wav.DecodeWav(wavFileBytes)
fmt.Println(w.BitsPerSample)

data, _ := w.GetData()

if _, ok := data.(*wav.StereoData); ok {
  fmt.Println(stereoData.RSample)
  fmt.Println(stereoData.LSample)
} else if _, ok := data.(*wav.MonoData); ok {
  fmt.Println(monoData.Sample)
}
```

### Todo

- More wav file details
- Converter
- Raw Waveform
- Waveform image customization
- Test
