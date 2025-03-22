package test

import "testing"

func TestMain(t *testing.T) {
	t.Run("1.test", TestGetInfoLogger)
	println("Test Get Info Logger done")
	t.Run("2.test", TestGetWarningLogger)
	println("Test Get Warning Logger done")
	t.Run("3.test", TestGetErrorLogger)
	println("Test Get Error Logger done")
}
