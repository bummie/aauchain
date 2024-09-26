package blockchain

// https://stackoverflow.com/a/70168266
func uint64ToLenBytes(v uint64, l int) (b []byte) {
	b = make([]byte, l)

	for i := 0; i < l; i++ {
		f := 8 * i
		b[i] = byte(v >> f)
	}

	return
}

func int64ToLenBytes(v int64, l int) (b []byte) {
	return uint64ToLenBytes(uint64(v), l)
}

// https://stackoverflow.com/a/70168266
