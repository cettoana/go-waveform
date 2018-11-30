package wav

import "errors"

// Sample sample
type Sample []float64

// MonoData mono data
type MonoData struct {
	Sample Sample
	Bound  *Bound
}

// StereoData stereo data
type StereoData struct {
	LSample Sample
	RSample Sample
	Bound   *Bound
}

// Bound sample value upper and lower boundary
type Bound struct {
	Upper float64
	Lower float64
}

// GetBound get sample value bound
func GetBound(bitsPerSample uint16, audioFormat AudioFormat) (*Bound, error) {
	if audioFormat == PCMFloating {
		return &Bound{
			Upper: 1.0,
			Lower: -1.0,
		}, nil
	}

	if audioFormat == 1 {
		if bitsPerSample == 8 {
			return &Bound{
				Upper: 255,
				Lower: 0,
			}, nil
		}

		if bitsPerSample == 16 {
			return &Bound{
				Upper: 32767,
				Lower: -32768,
			}, nil
		}

		if bitsPerSample == 32 {
			return &Bound{
				Upper: 2147483647,
				Lower: -2147483648,
			}, nil
		}
	}

	return nil, errors.New("audio format not support")
}
