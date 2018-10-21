package random

const (
	alphaNumerics = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	numberString      = "0123456789"
	numberStringBytes = []byte(numberString)

	letterString      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterStringBytes = []byte(letterString)

	poundTag       = []byte("#")[0]
	quetionMarkTag = []byte("?")[0]
)
