package waveform

// WaveFormat type
type WaveFormat uint16

// WaveFormat wave format tag
const (
	WaveFormatPCM        WaveFormat = 0x0001
	WaveFormatIEEEFloat             = 0x0003
	WaveFormatALaw                  = 0x0006
	WaveFormatMULaw                 = 0x0007
	WaveFormatExtensible            = 0xFFFE
)

func (f WaveFormat) String() string {
	switch f {
	case 1:
		return "PCM"
	case 3:
		return "IEEE Float"
	case 6:
		return "A-lAW"
	case 7:
		return "μ-law"
	case 65534:
		return "SubFormat"
	default:
		return "Unknown"
	}
}
