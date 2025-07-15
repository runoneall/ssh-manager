package sshShell

func TokenAt(slice []string, index int) string {
	if index < len(slice) {
		return slice[index]
	}
	return ""
}
