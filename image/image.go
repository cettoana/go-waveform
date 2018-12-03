package image

import (
	"errors"
	"fmt"
	"image/color"
	"math"

	"github.com/cettoana/go-waveform"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// XY struct
type XY struct {
	X float64
	Y float64
}

// Option image option
type Option struct {
	Fast       bool
	FileName   string
	FileType   string
	Resolution int
	Style      string
	Width      int
	Theme      string
}

// OutputWaveformImage output waveform image
func OutputWaveformImage(data interface{}, option *Option) error {
	var outputFn func(waveform.Sample, *waveform.Bound, *Option, string) error

	if option.Style == "original" {
		outputFn = outputOriginalWavefromImage
	} else {
		outputFn = outputWaveformImage
	}

	switch data.(type) {
	case *waveform.MonoData:
		mono := data.(*waveform.MonoData)

		return outputFn(mono.Sample, mono.Bound, option, "")
	case *waveform.StereoData:
		stereo := data.(*waveform.StereoData)

		if err := outputFn(stereo.LSample, stereo.Bound, option, "-L"); err != nil {
			return err
		}

		return outputFn(stereo.RSample, stereo.Bound, option, "-R")
	default:
		return errors.New("sound system unsupport")
	}
}

func outputOriginalWavefromImage(sample waveform.Sample, bound *waveform.Bound, option *Option, postfix string) error {
	var points plotter.XYs

	p, err := plot.New()
	if err != nil {
		return err
	}

	if option.Fast {
		m := int(math.Ceil(float64(len(sample)) / 100000))

		for i := 0; i < len(sample); i += m {
			points = append(points, XY{X: float64(i), Y: sample[i]})
		}
	} else {
		for i, v := range sample {
			points = append(points, XY{X: float64(i), Y: v})
		}
	}

	l, err := plotter.NewLine(points)
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{G: 225, B: 255, A: 255}

	p.Add(l)

	p.HideX()
	p.HideY()
	p.X.Min = 0
	p.X.Max = float64(len(sample))
	p.Y.Min = bound.Lower
	p.Y.Max = bound.Upper
	p.BackgroundColor = color.Black

	fileName := fmt.Sprintf("%s%s.%s", option.FileName, postfix, option.FileType)

	return p.Save(vg.Points(5000), vg.Points(540), fileName)
}

func outputWaveformImage(sample waveform.Sample, bound *waveform.Bound, option *Option, postfix string) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	floor := (bound.Upper + bound.Lower) / 2

	n := option.Resolution
	m := int(float64(len(sample))/float64(n) + 0.5)

	if n > len(sample) {
		n = len(sample)
		m = 1
	}

	stroke := float64(2)
	width := float64(n * 5)

	if option.Width != 0 {
		width = float64(option.Width)
		stroke = width / float64(n) * 0.4
	}

	i := 0
	d := 1
	g := 155

	for i = 0; i+m < len(sample); i += m {
		xys := getXYs(i, sample[i:i+m], floor)

		l, err := plotter.NewLine(xys)
		if err != nil {
			return err
		}

		l.LineStyle.Width = vg.Points(stroke)
		l.Color = &color.RGBA{R: 50, G: uint8(g), B: 240, A: 255}

		p.Add(l)

		g += d

		if g > 225 {
			g = 225 - 1
			d = -1
		} else if g < 155 {
			g = 156
			d = 1
		}
	}

	xys := getXYs(i, sample[i:len(sample)-1], floor)

	l, err := plotter.NewLine(xys)
	if err != nil {
		return err
	}

	l.LineStyle.Width = vg.Points(stroke)
	l.Color = &color.RGBA{R: 50, G: uint8(g), B: 240, A: 255}

	p.Add(l)

	p.HideX()
	p.HideY()
	p.X.Min = 0
	p.X.Max = float64(len(sample))
	p.Y.Min = bound.Lower
	p.Y.Max = bound.Upper
	p.BackgroundColor = getBackgroundColor(option.Theme)

	fileName := fmt.Sprintf("%s%s.%s", option.FileName, postfix, option.FileType)

	return p.Save(vg.Points(width), vg.Points(540), fileName)
}

func getXYs(x int, s []float64, floor float64) *plotter.XYs {
	max := floor
	min := floor

	for _, y := range s {
		if y > floor {
			if y > max {
				max = y
			}
		} else {
			if y < min {
				min = y
			}
		}
	}

	return &plotter.XYs{
		XY{X: float64(x), Y: min},
		XY{X: float64(x), Y: max},
	}
}

func getBackgroundColor(theme string) color.Color {
	if theme == "dark" {
		return color.Black
	}

	return color.White
}
