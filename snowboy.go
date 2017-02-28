package snowboy

import (
	"unsafe"
	"errors"

	"github.com/Kitt-AI/snowboy/swig/Go"
)

type Detector struct {
	raw snowboydetect.SnowboyDetect
	keywords          []Keyword
	initialized       bool
}

func NewDetector(resourceFilename string, words Hotwords) Detector {
	d := Detector{}
	d.raw = snowboydetect.NewSnowboyDetect(resourceFilename, words.modelStr)
	d.raw.SetSensitivity(words.sensitivityStr)
	d.raw.SetAudioGain(1.0)
	d.keywords = words.keywords
	d.initialized = true
	return d
}

func (d *Detector) SetAudioGain(gain float32) {
	d.raw.SetAudioGain(gain)
}

func (d *Detector) Close() error {
	if d.initialized {
		d.initialized = false
		snowboydetect.DeleteSnowboyDetect(d.raw)
		return nil
	} else {
		return errors.New("snowboy not initialize")
	}
}

func (d *Detector) RunDetection(data []byte) (*Keyword, error) {
	if len(data) == 0 {
		return nil, nil
	}
	ptr := snowboydetect.SwigcptrInt16_t(unsafe.Pointer(&data[0]))
	res := d.raw.RunDetection(ptr, len(data) / 2 /* len of int16 */)
	if res == -2 {
		k := KeywordSilence
		return &k, nil
	} else if res == -1 {
		// TODO: extract real error if possible
		return nil, errors.New("snowboy error")
	} else if res == 0 {
		return nil, nil
	} else {
		return &d.keywords[res - 1], nil
	}
}
