package day06

func FindStartOfPacket(datastream string) int {
	for i := 0; i < len(datastream)-4; i++ {
		if isStartOfPacket(datastream[i : i+4]) {
			return i + 4
		}
	}

	panic("cannot find start of packet")
}

func isStartOfPacket(window string) bool {
	return window[0] != window[1] &&
		window[0] != window[2] &&
		window[0] != window[3] &&

		window[1] != window[2] &&
		window[1] != window[3] &&

		window[2] != window[3]
}
