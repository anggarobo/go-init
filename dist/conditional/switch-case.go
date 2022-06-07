package conditional

func MsgGrade(val uint8) string {
	switch {
	case val >= 8:
		return "Perfect!"
	case ( val < 8 ) && ( val >= 6 ):
		return "Awesome!"
	default:
		return "You need to learn more ya!"
	}
}