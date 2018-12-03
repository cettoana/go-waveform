package waveform

import (
	"encoding/binary"
	"errors"
	"math"
)

// Parser type
type Parser func([]byte) float64

// DecodeWav decode wav file information from bytes
func DecodeWav(bytes []byte) *Wav {
	waveFormat := WaveFormat(binary.LittleEndian.Uint16(bytes[20:22]))

	numChannels := binary.LittleEndian.Uint16(bytes[22:24])

	sampleRate := binary.LittleEndian.Uint32(bytes[24:28])

	bitsPerSample := binary.LittleEndian.Uint16(bytes[34:36])

	subchunk1Size := binary.LittleEndian.Uint32(bytes[16:20])

	subchunk2Start := 20 + subchunk1Size
	subchunk2ID := string(bytes[subchunk2Start : subchunk2Start+4])
	subchunk2Size := binary.LittleEndian.Uint32(bytes[subchunk2Start+4 : subchunk2Start+8])

	dataStart := subchunk2Start + 8
	dataSize := subchunk2Size

	if subchunk2ID == "fact" {
		subchunk3Start := subchunk2Start + 8 + subchunk2Size
		subchunk3Size := binary.LittleEndian.Uint32(bytes[subchunk3Start+4 : subchunk3Start+8])

		dataStart = subchunk3Start + 8
		dataSize = subchunk3Size
	}

	data := bytes[dataStart : dataStart+dataSize]

	return &Wav{
		WaveFormat:    waveFormat,
		NumChannels:   numChannels,
		SampleRate:    sampleRate,
		BitsPerSample: bitsPerSample,
		DataChuckSize: dataSize,
		Data:          data,
	}
}

// 0 to 255
func int8BitsParser(b []byte) float64 {
	return float64(b[0])
}

// -32768 to 32767
func int16BitsParser(b []byte) float64 {
	value := int16(binary.LittleEndian.Uint16(b))
	return float64(value)
}

func int32BitsParser(b []byte) float64 {
	value := int32(binary.LittleEndian.Uint32(b))
	return float64(value)
}

func float32BitsParser(b []byte) float64 {
	bits := binary.LittleEndian.Uint32(b)
	value := math.Float32frombits(bits)
	return float64(value)
}

func float64BitsParser(b []byte) float64 {
	bits := binary.LittleEndian.Uint64(b)
	return math.Float64frombits(bits)
}

// GetSampleParser get sample parser
func GetSampleParser(bitsPerSample uint16, waveFormat WaveFormat) (func([]byte) float64, error) {
	if waveFormat == WaveFormatPCM {
		if bitsPerSample == 8 {
			return int8BitsParser, nil
		}

		if bitsPerSample == 16 {
			return int16BitsParser, nil
		}

		if bitsPerSample == 32 {
			return int32BitsParser, nil
		}
	}

	if waveFormat == WaveFormatIEEEFloat {
		if bitsPerSample == 32 {
			return float32BitsParser, nil
		}

		if bitsPerSample == 64 {
			return float64BitsParser, nil
		}
	}

	return nil, errors.New("format not support")
}
