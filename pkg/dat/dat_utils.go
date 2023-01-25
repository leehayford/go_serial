package dat

import (
	"encoding/binary"
	// "log"
)

func BytesToUInt16(bytes []byte) uint16 {

	x := []byte{}
	n := len(bytes)
	// log.Println("Received bytes:\t", n)

	if n < 2 {
		// Pad the high end
		cnt := 2 - n
		// log.Println("Padding high end with bytes:\t", cnt)
		for cnt > 0 {
			x = append(x, 0x00)
			cnt--
		}
		x = append(x, bytes[0:n+1]...)
	} else if n > 2 {
		// Grab the low end
		// log.Println("Grabbing the two low bytes")
		x = bytes[n-2 : n+1]
	} else {
		x = bytes
	}

	return binary.BigEndian.Uint16(x)
}

func BytesToUInt32(bytes []byte) uint32 {

	x := []byte{}
	n := len(bytes)
	// log.Println("Received bytes:\t", n)
	if n < 4 {
		// Pad the high end
		cnt := 4 - n
		// log.Println("Padding high end with bytes:\t", cnt)
		for cnt > 0 {
			x = append(x, 0x00)
			cnt--
		}
		x = append(x, bytes[0:n+1]...)
	} else if n > 4 {
		// Grab the low end
		// log.Println("Grabbing the two low bytes")
		x = bytes[n-4 : n+1]
	} else {
		x = bytes
	}

	return binary.BigEndian.Uint32(x)
}

func MaxSliceUInt32(slice []uint32) uint32 {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func MinSliceUInt32(slice []uint32) uint32 {
	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min
}

func MinMaxUInt32(slice []uint32) (uint32, uint32) {

	min := slice[0]
	max := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}

func StripFFs(arr []byte) []byte {
	for i, v := range arr {
		if v == 0xFF {
			arr = append(arr[0:i])
			break
		}
	}
	return arr
}
