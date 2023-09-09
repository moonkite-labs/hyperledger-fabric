package utils

// Stops execution if an error is found
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
