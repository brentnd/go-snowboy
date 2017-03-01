package snowboy

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"unsafe"

	"github.com/Kitt-AI/snowboy/swig/Go"
)

// Detector is holds the context and base impl for snowboy audio detection
type Detector struct {
	raw            snowboydetect.SnowboyDetect
	initialized    bool
	handlers       map[int]handlerKeyword
	modelStr       string
	sensitivityStr string
	ResourceFile   string
	AudioGain      float32
}

// Creates a standard Detector from a resources file
// Gives a default gain of 1.0
func NewDetector(resourceFile string) Detector {
	return Detector{
		ResourceFile: resourceFile,
		AudioGain: 1.0,
	}
}

// Close handles cleanup required by snowboy library
//
// Clients must call Close on detectors after doing any detection
// Returns error if Detector was never used
func (d *Detector) Close() error {
	if d.initialized {
		d.initialized = false
		snowboydetect.DeleteSnowboyDetect(d.raw)
		return nil
	} else {
		return errors.New("snowboy not initialize")
	}
}

// Install a handler for the given hotword
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

// Installs a handle for the given hotword based on the func argument
// instead of the Handler interface
func (d *Detector) HandleFunc(hotword Hotword, handler func(string)) {
	d.Handle(hotword, handlerFunc(handler))
}

// Reads from data and calls previously installed handlers when detection occurs
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

func (d *Detector) initialize() {
	d.raw = snowboydetect.NewSnowboyDetect(d.ResourceFile, d.modelStr)
	d.raw.SetSensitivity(d.sensitivityStr)
	d.raw.SetAudioGain(d.AudioGain)
	d.initialized = true
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

func (d *Detector) runDetection(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	ptr := snowboydetect.SwigcptrInt16_t(unsafe.Pointer(&data[0]))
	return d.raw.RunDetection(ptr, len(data) / 2 /* len of int16 */)
}

// A Handler is used to handle when keywords are detected
//
// Detected will be call with the keyword string
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

type handlerFunc func(string)

func (f handlerFunc) Detected(keyword string) {
	f(keyword)
}

// A Hotword represents a model filename and sensitivity for a snowboy detectable word
//
// Model is the filename for the .umdl file
//
// Sensitivity is the sensitivity of this specific hotword
//
// Name is what will be used in calls to Handler.Detected(string)
type Hotword struct {
	Model       string
	Sensitivity float32
	Name        string
}

// Creates a hotword from model and sensitivity only, parsing
// the hotward name from the model filename
func NewHotword(model string, sensitivity float32) Hotword {
	h := Hotword{
		Model: model,
		Sensitivity: sensitivity,
	}
	name := strings.TrimRight(model, ".umdl")
	nameParts := strings.Split(name, "/")
	h.Name = nameParts[len(nameParts) - 1]
	return h
}
