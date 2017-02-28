package snowboy

import (
	"unsafe"
	"errors"

	"github.com/Kitt-AI/snowboy/swig/Go"
	"strconv"
)

type Keyword string
const (
	KeywordSilence = Keyword("silence")
	KeywordAlexa   = Keyword("alexa")
	KeywordSnowboy = Keyword("snowboy")
)

type Hotword struct {
	Model       string
	Sensitivity float32
	Keyword     Keyword
}

type Detector struct {
	raw snowboydetect.SnowboyDetect
	keywords          []Keyword
	sensitivity       string
	modelStr          string
	initialized       bool
	ResourceFilename  string
	AudioGain         float32
}

func (d *Detector) AddHotword(h Hotword) error {
	if d.initialized {
		return errors.New("detector already initialized")
	}
	if len(d.keywords) > 0 {
		d.modelStr += ","
		d.sensitivity += ","
	}
	d.modelStr += h.Model
	d.sensitivity += strconv.FormatFloat(float64(h.Sensitivity), 'f', 2, 64)
	d.keywords = append(d.keywords, h.Keyword)
	return nil
}

func (d *Detector) initialize() {
	if d.initialized {
		return
	}
	d.raw = snowboydetect.NewSnowboyDetect(d.ResourceFilename, d.modelStr)
	d.raw.SetSensitivity(d.sensitivity)
	d.raw.SetAudioGain(d.AudioGain)
	d.initialized = true
}

func (d *Detector) Close() error {
	if d.initialized {
		snowboydetect.DeleteSnowboyDetect(d.raw)
		return nil
	} else {
		return errors.New("snowboy not initialize")
	}
}

func (d *Detector) RunDetection(data []byte) (*Keyword, error) {
	d.initialize()
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
