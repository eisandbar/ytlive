package getter

func matchEnds(startString, endString, leftPad, rightPad string, data []byte) [][]byte {
	res := make([][]byte, 0)
	startByte, endByte := []byte(startString), []byte(endString)
	start, end, length := 0, 0, 0
	for i := range data {
		if end == len(endByte) {
			temp := makeString(leftPad, rightPad, data, length, end, i)
			res = append(res, temp)
			start, end, length = 0, 0, 0
			// Here we don't continue as we already started the next block
		}
		if start < len(startByte) {
			if data[i] == startByte[start] {
				start++
			} else {
				start = 0
				if data[i] == startByte[start] {
					start++
				}
			}
			continue
		}

		// increases only when we matched the startString
		length++

		if data[i] == endByte[end] {
			end++
		} else {
			end = 0
			if data[i] == endByte[end] {
				end++
			}
		}

	}
	return res
}

func makeString(leftPad, rightPad string, data []byte, length, end, current int) []byte {
	temp := make([]byte, len(leftPad)+len(rightPad)+length-end)
	ind := 0

	for i := range leftPad {
		temp[ind] = leftPad[i]
		ind++
	}
	for i := current - length; i < current-end; i++ {
		temp[ind] = data[i]
		ind++
	}
	for i := range rightPad {
		temp[ind] = rightPad[i]
		ind++
	}
	return temp
}
