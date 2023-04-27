package utils

func IsList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
