package ptr

func String(s string) (*string) {
	return &s
}

func Uint32(i uint32) (*uint32) {
	return &i
}

func Int32(i int32) (*int32) {
	return &i
}

func Bool(b bool) (*bool) {
	return &b
}

func Float64(f float64) (*float64) {
	return &f
}