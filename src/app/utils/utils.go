package utils

func VerifyRequiredFields(requiredFields []string, input map[string]any) (missing map[string]string) {
	missing = make(map[string]string)

	for _, field := range requiredFields {
		if _, ok := input[field]; !ok {
			missing[field] = "required"
		}
	}
	return missing
}
