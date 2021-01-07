package main

import (
	"os"
	"strings"

	"github.com/cettoana/go-waveform"
	"github.com/cettoana/go-waveform/image"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "23:59:59"})

	app := &cli.App{
		Name:    "go-waveform",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "svg",
				Usage:   "Output waveform format",
			},
			&cli.StringFlag{
				Name:    "resolution",
				Aliases: []string{"r"},
				Value:   "1000",
				Usage:   "Output waveform resolution",
			},
			&cli.StringFlag{
				Name:    "width",
				Aliases: []string{"w"},
				Value:   "0",
				Usage:   "Output image file width",
			},
			&cli.StringFlag{
				Name:    "theme",
				Aliases: []string{"t"},
				Value:   "dark",
				Usage:   "Output image theme (light/dark)",
			},
		},
		Action: func(c *cli.Context) error {
			pwd, err := os.Getwd()
			if err != nil {
				log.Error().Msg(err.Error())
				return err
			}

			fileName := c.Args().First()
			format := c.String("format")
			theme := c.String("theme")
			resolution := c.Int("resolution")
			width := c.Int("width")

			fi, err := os.Open(pwd + "/" + fileName)
			if err != nil {
				log.Error().Msg(err.Error())
				return err
			}

			stat, err := fi.Stat()
			if err != nil {
				log.Error().Msg(err.Error())
				return err
			}

			log.Info().Msg("read data...")

			b := make([]byte, stat.Size())

			if _, err = fi.Read(b); err != nil {
				log.Error().Msg(err.Error())
				return err
			}

			log.Info().Msg("read complete")

			w := waveform.DecodeWav(b)

			log.Info().Str("WaveFormat", w.WaveFormat.String()).Msg("fmt Chunk")
			log.Info().Uint16("NumChannels", w.NumChannels).Msg("fmt Chunk")
			log.Info().Uint32("SampleRate", w.SampleRate).Msg("fmt Chunk")
			log.Info().Uint16("BitsPerSample", w.BitsPerSample).Msg("fmt Chunk")
			log.Info().Uint32("BitsPerSample", w.DataChuckSize).Msg("data Chunk")

			data, err := w.GetData()
			if err != nil {
				log.Error().Msg(err.Error())
				return err
			}

			if err := image.OutputWaveformImage(data, &image.Option{
				FileName:   strings.Replace(fileName, ".wav", "", 1),
				FileType:   format,
				Width:      width,
				Style:      "default",
				Theme:      theme,
				Fast:       false,
				Resolution: resolution,
			}); err != nil {
				log.Error().Msg(err.Error())
				return err
			}

			log.Info().Msg("complete")

			return nil
		},
	}

	app.Run(os.Args)
}
