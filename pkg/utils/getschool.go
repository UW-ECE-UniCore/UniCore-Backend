package utils

import (
	"strings"
	"unicore/pkg/constants"
)

var (
	schoolMap = map[string]int{
		"uwaterloo.ca":   constants.UW,
		"utoronto.ca":    constants.UofT,
		"student.ubc.ca": constants.UBC,
		"mcgill.ca":      constants.McGill,
		"mail.mcgill.ca": constants.McGill,
		"ualberta.ca":    constants.UAlberta,
		"uwo.ca":         constants.Western,
		"umontreal.ca":   constants.Montreal,
	}
)

func GetSchool(email string) int {
	postfix := strings.Split(email, "@")[1]
	if school, ok := schoolMap[postfix]; ok {
		return school
	}
	return -1
}
