package main

func errorChecks(args []string) (bool, string) {
	errors := ""
	if len(args) != 2 {
		errors = "Incorrect length of arguments. Should be 2"
		return true, errors
	}
	return false, errors
}
