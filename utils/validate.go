package utils

func CheckSubject(subject string) bool {
	available_subjects := []string{"amdl", "eiot", "wm", "nlp"}

	for _, val := range available_subjects {
		if val == subject {
			return true
		}
	}

	return false
}
