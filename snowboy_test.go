package snowboy

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewDetector(t *testing.T) {
	filename := "testFile"
	d := NewDetector(filename)
	assert.Equal(t, filename, d.ResourceFile, "NewDetect sets ResourceFile")
	assert.Equal(t, float32(1.0), d.AudioGain, "Gain should default to 1.0")
}
