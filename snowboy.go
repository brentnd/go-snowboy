package snowboy

import (
	"unsafe"
	"errors"

	"github.com/Kitt-AI/snowboy/swig/Go"
)

type detector struct {
	raw snowboydetect.SnowboyDetect
}

func NewSnowboyDetector(resource, keywords, sensitivity string, audioGain float32) detector {
	d := detector{}
	d.raw = snowboydetect.NewSnowboyDetect(resource, keywords)
	d.raw.SetSensitivity("0.5")
	d.raw.SetAudioGain(audioGain)
	return d
}

func (d *detector) Close() error {
	if d.raw != nil {
		snowboydetect.DeleteSnowboyDetect(d.raw)
		return nil
	} else {
		return errors.New("snowboy not initialize")
	}
}

func (d *detector) RunDetection(data []byte) int {
	ptr := snowboydetect.SwigcptrInt16_t(unsafe.Pointer(&data[0]))
	return d.raw.RunDetection(ptr, len(data) / 2 /* len of int16 */)
}
