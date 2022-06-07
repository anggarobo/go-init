package operator

func OnProcess(loading bool, val uint8) string {
	var isEmpty = !loading && val > 0

	if isEmpty {
		return "done"
	} else {
		return "Please wait"
	}
}