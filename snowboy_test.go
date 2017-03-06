package snowboy

import (
	"os"
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	resourceFile = "build/common.res"
	alexaKeywordFile = "build/alexa.umdl"
)

func TestNewDetector(t *testing.T) {
	d := NewDetector(resourceFile)
	defer d.Close()

	assert.Equal(t, resourceFile, d.ResourceFile, "NewDetect sets ResourceFile")
	assert.Equal(t, float32(1.0), d.AudioGain, "Gain should default to 1.0")
}

func TestDetector_AudioFormat(t *testing.T) {
	d := NewDetector(resourceFile)
	defer d.Close()

	d.Handle(NewDefaultHotword(alexaKeywordFile), nil)

	a, b, c := d.AudioFormat()
	assert.Equal(t, 16000, a)
	assert.Equal(t, 1, b)
	assert.Equal(t, 16, c)
}

func TestDetector_HandleSilenceFunc(t *testing.T) {
	d := NewDetector(resourceFile)
	defer d.Close()

	// Not under test but underlying lib requires a keyword
	d.Handle(NewDefaultHotword(alexaKeywordFile), nil)

	handled := false
	d.HandleSilenceFunc(1 * time.Second, func(keyword string) {
		handled = true
	})

	silence, err := os.Open("audio/silence.wav")
	if err != nil {
		t.Error("can't open silence wav")
	}
	d.ReadAndDetect(silence)
	if !handled {
		t.Error("handler in HandleSilenceFunc never called")
	}
}
