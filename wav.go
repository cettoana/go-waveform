package wav

import "errors"

// Wav struct
type Wav struct {
	AudioFormat   AudioFormat
	NumChannels   uint16
	SampleRate    uint32
	BitsPerSample uint16
	DataChuckSize uint32

	Data []byte
}

// GetData get wav audio data
func (w *Wav) GetData() (interface{}, error) {
	bytePerSample := int(w.BitsPerSample / 8)
	sampleParser, err := GetSampleParser(w.BitsPerSample, w.AudioFormat)
	if err != nil {
		return nil, err
	}

	if w.NumChannels == 1 {
		sample := parseMonoSample(w.Data, bytePerSample, sampleParser)
		bound, err := GetBound(w.BitsPerSample, w.AudioFormat)
		if err != nil {
			return nil, err
		}

		sample.Bound = bound
		return sample, nil
	}

	if w.NumChannels == 2 {
		sample := parseStereoSample(w.Data, bytePerSample, sampleParser)
		bound, err := GetBound(w.BitsPerSample, w.AudioFormat)
		if err != nil {
			return nil, err
		}

		sample.Bound = bound
		return sample, nil
	}

	return nil, errors.New("failed to sampled data from wav file")
}

func parseMonoSample(data []byte, bytePerSample int, parser Parser) *MonoData {
	end := len(data) / bytePerSample

	sample := make([]float64, 0)

	for i := 0; i < end; i += bytePerSample {
		s := parser(data[i : i+bytePerSample])

		sample = append(sample, s)
	}

	return &MonoData{
		Sample: sample,
	}
}

func parseStereoSample(data []byte, bytePerSample int, parser Parser) *StereoData {
	offset := bytePerSample * 2
	end := len(data)

	lSample := make([]float64, 0)
	rSample := make([]float64, 0)

	for i := 0; i < end; i += offset {
		l := parser(data[i : i+bytePerSample])
		r := parser(data[i+bytePerSample : i+offset])

		lSample = append(lSample, l)
		rSample = append(rSample, r)
	}

	return &StereoData{
		LSample: lSample,
		RSample: rSample,
	}
}
