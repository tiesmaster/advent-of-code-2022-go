package day06

func FindFirstStartOfPacket(datastream string) int {
	const windowSize = 4
	return findSequenceStart(datastream, windowSize) + windowSize
}

func FindFirstStartOfMessage(datastream string) int {
	const windowSize = 14
	return findSequenceStart(datastream, windowSize) + windowSize
}

func findSequenceStart(datastream string, windowSize int) int {
	for i := 0; i < len(datastream)-windowSize; i++ {
		if isStartOfPacket(datastream[i : i+windowSize]) {
			return i
		}
	}

	panic("cannot find start of packet")
}

func isStartOfPacket(window string) bool {
	index := make(map[rune]int)
	for _, r := range window {
		index[r]++
	}

	for _, i := range index {
		if i > 1 {
			return false
		}
	}

	return true
}
