package snowboy

import (
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

type Hotwords struct {
	modelStr       string
	sensitivityStr string
	keywords       []Keyword
}

func (h *Hotwords) Add(word Hotword) {
	if len(h.keywords) > 0 {
		h.modelStr += ","
		h.sensitivityStr += ","
	}
	h.modelStr += word.Model
	h.sensitivityStr += strconv.FormatFloat(float64(word.Sensitivity), 'f', 2, 64)
	h.keywords = append(h.keywords, word.Keyword)
}
