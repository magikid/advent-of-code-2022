package comms

func numberUniqueCharacters(numUnique int, buffer string) int {
	for i := numUnique; i < len(buffer); i++ {
		if unique(buffer[i-numUnique : i]) {
			return i
		}
	}

	return 0
}

func (d *Datastream) StartOfPacketMarkerLocation() int {
	if d.startOfPacketMarkerLocation != 0 {
		return d.startOfPacketMarkerLocation
	}

	return numberUniqueCharacters(4, d.DatastreamBuffer)
}

func (d *Datastream) StartOfMessageMarkerLocation() int {
	if d.startOfMessageMarkerLocation != 0 {
		return d.startOfMessageMarkerLocation
	}

	return numberUniqueCharacters(14, d.DatastreamBuffer)
}

// Found here: https://codereview.stackexchange.com/a/251210/268200
func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
