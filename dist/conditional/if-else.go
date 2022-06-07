package conditional

import "strconv"

func CheckGrade(point float32) string {
	if grade := point / 100; grade >= 100 {
		return "perfect point! "+ strconv.FormatFloat(float64(grade), 'f', 1, 32) + "%"
		} else if grade >= 70 {
		return "good point! "+ strconv.FormatFloat(float64(grade), 'f', 1, 32) + "%"
		} else {
			return "not bad point! "+ strconv.FormatFloat(float64(grade), 'f', 1, 32) + "%"
	}
}