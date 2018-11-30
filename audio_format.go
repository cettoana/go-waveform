package wav

// AudioFormat type
type AudioFormat uint16

// Audio formats from the format chunk, byte[20:22]
const (
	PCMInteger           AudioFormat = 1
	ADPCM                            = 2
	PCMFloating                      = 3
	ULaw                             = 7
	WaveFormatExtensible             = 65534
)

func (f AudioFormat) String() string {
	switch f {
	case 1:
		return "Integer PCM"
	case 2:
		return "ADPCM"
	case 3:
		return "Floating-point PCM"
	case 7:
		return "μ-law"
	case 65534:
		return "WaveFormatExtensible"
	default:
		return "Unknown"
	}
}
