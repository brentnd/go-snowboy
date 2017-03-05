package snowboy

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	resourceFile = "build/common.res"
	alexaKeywordFile = "build/alexa.umdl"
)

func TestNewDetector(t *testing.T) {
	d := NewDetector(resourceFile)
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
