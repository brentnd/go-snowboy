package snowboy

const(
	EndpointBase = "https://snowboy.kitt.ai/api/"
	EndpointVersion = "v1"
	EndpointTrain = EndpointBase + EndpointVersion + "/train"
	EndpointImprove = EndpointBase + EndpointVersion + "/improve"

)

type Language string
const (
	LanguageArabic Language = "ar"
	LanguageChinese         = "zh"
	LanguageDutch           = "nl"
	LanguageEnglish         = "en"
	LanguageFrench          = "fr"
	LanguageGerman          = "dt"
	LanguageHindi           = "hi"
	LanguageItalian         = "it"
	LanguageJapanese        = "jp"
	LanguageKorean          = "ko"
	LanguagePersian         = "fa"
	LanguagePolish          = "pl"
	LanguagePortuguese      = "pt"
	LanguageRussian         = "ru"
	LanguageSpanish         = "es"
	LanguageOther           = "ot"
)

type AgeGroup string
const (
	AgeGroup0s AgeGroup = "0_9"
	AgeGroup10s         = "10_19"
	AgeGroup20s         = "20_29"
	AgeGroup30s         = "30_39"
	AgeGroup40s         = "40_49"
	AgeGroup50s         = "50_59"
	AgeGroup60plus      = "60+"
)

type Gender string
const (
	GenderMale Gender = "M"
	GenderFemale      = "F"
)

type VoiceSample struct {
	Wave string `json:"wave"`
}

type TrainRequest struct {
	VoiceSamples []VoiceSample `json:"voice_samples"`
	Token          string      `json:"token"`
	Name           string      `json:"name"`
	Language       Language    `json:"language"`
	AgeGroup       AgeGroup    `json:"age_group"`
	Gender         Gender      `json:"gender"`
	Microphone     string      `json:"microphone"`
}