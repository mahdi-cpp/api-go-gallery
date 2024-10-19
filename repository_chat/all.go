package repository_chat

import (
	"github.com/gin-gonic/gin"
	"math"
)

func RestChatV2() map[string]any {
	return gin.H{
		"instagramPostDTO1": instagramPostDTO1,
		"instagramPostDTO2": instagramPostDTO2,
		"instagramPostDTO3": instagramPostDTO3,
		"storyDTO":          storyDTO,
		"movieDTO":          movieDTO,
		"pdfDTO":            pdfDTO,
		"electronicDTO":     electronicDTO,
		"mapDTO":            mapDTO,
		"questionVoiceDTO":  questionVoiceDTO,
	}
}

func InitModels() {
	instagramPostDTO1 = GetInstagram("/var/cloud/id/girl/", "/var/cloud/call2/far.jpg")
	instagramPostDTO2 = GetInstagram("/var/cloud/id/cut/", "/var/cloud/call2/nar3.jpg")
	instagramPostDTO3 = GetInstagram("/var/cloud/fa/", "/var/cloud/call2/01.jpg")

	storyDTO = GetStory("/var/cloud/bb/", "/var/cloud/call2/01.jpg")
	movieDTO = GetMovies("/var/cloud/chat/movie/movie/")
	pdfDTO = GetPdfs("/var/cloud/chat/pdf/")
	electronicDTO = GetElectronic("/var/cloud/chat/electronic/")
	mapDTO = GetMaps("/var/cloud/chat/map/")
	questionVoiceDTO = GetQuestionVoices("/var/cloud/chat/voice/")
}

func dp(value float32) float32 {
	if value == 0 {
		return 0
	}
	return float32(math.Ceil(float64(2.625 * value)))
}
