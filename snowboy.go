package snowboy

import (
	"unsafe"
	"errors"

	"github.com/Kitt-AI/snowboy/swig/Go"
	"io"
	"strconv"
	"fmt"
	"io/ioutil"
)

type Handler interface {
	Detected(string)
}

type handlerKeyword struct {
	Handler
	keyword string
}

func (h handlerKeyword) call() {
	h.Handler.Detected(h.keyword)
}

type HandlerFunc func(string)

func (f HandlerFunc) Detected(keyword string) {
	f(keyword)
}

type Hotword struct {
	Model       string
	Sensitivity float32
	Name        string
}

type Detector struct {
	raw            snowboydetect.SnowboyDetect
	initialized    bool
	handlers       map[int]handlerKeyword
	modelStr       string
	sensitivityStr string
	ResourceFile   string
	AudioGain      float32
}

func (d *Detector) initialize() {
	d.raw = snowboydetect.NewSnowboyDetect(d.ResourceFile, d.modelStr)
	d.raw.SetSensitivity(d.sensitivityStr)
	d.raw.SetAudioGain(d.AudioGain)
	d.initialized = true
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

func (d *Detector) runDetection(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	ptr := snowboydetect.SwigcptrInt16_t(unsafe.Pointer(&data[0]))
	return d.raw.RunDetection(ptr, len(data) / 2 /* len of int16 */)
}

func (d *Detector) route(result int) {
	if result == -2 {
		// TODO: silence handler
		fmt.Println("silence")
	} else if result == -1 {
		fmt.Println("error in snowboy")
	} else if result == 0 {
		fmt.Println("nothing detected")
	} else if result > 0 {
		handlerKeyword, ok := d.handlers[result]
		if ok {
			handlerKeyword.call()
		} else {
			fmt.Println("no handler for result", result)
		}
	}
}

func (d *Detector) Handle(hotword Hotword, handler Handler) {
	if len(d.handlers) > 0 {
		d.modelStr += ","
		d.sensitivityStr += ","
	}
	d.modelStr += hotword.Model
	d.sensitivityStr += strconv.FormatFloat(float64(hotword.Sensitivity), 'f', 2, 64)
	if d.handlers == nil {
		d.handlers = make(map[int]handlerKeyword)
	}
	d.handlers[len(d.handlers) + 1] = handlerKeyword{
		Handler: handler,
		keyword: hotword.Name,
	}
}

func (d *Detector) HandleFunc(hotword Hotword, handler func(string)) {
	d.Handle(hotword, HandlerFunc(handler))
}

func (d *Detector) ReadAndDetect(data io.Reader) (err error) {
	d.initialize()
	// TODO: buffer data into chunks
	bytes, err := ioutil.ReadAll(data)
	if err != nil {
		return
	}
	d.route(d.runDetection(bytes))
	return
}
