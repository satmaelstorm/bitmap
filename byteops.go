package bitmap

func Int64ToBytes(in int64) [8]byte {
	var result [8]byte
	result[0] = byte(in & 0xFF)
	result[1] = byte((in >> 8) & 0xFF)
	result[2] = byte((in >> 16) & 0xFF)
	result[3] = byte((in >> 24) & 0xFF)
	result[4] = byte((in >> 32) & 0xFF)
	result[5] = byte((in >> 40) & 0xFF)
	result[6] = byte((in >> 48) & 0xFF)
	result[7] = byte((in >> 56) & 0xFF)
	return result
}

func BytesToInt64(in [8]byte) int64 {
	var result int64
	for i, v := range in {
		add := int64(v) << (i * 8)
		result |= add
	}
	return result
}

func Int32ToBytes(in int32) [4]byte {
	var result [4]byte
	result[0] = byte(in & 0xFF)
	result[1] = byte((in >> 8) & 0xFF)
	result[2] = byte((in >> 16) & 0xFF)
	result[3] = byte((in >> 24) & 0xFF)
	return result
}

func BytesToInt32(in [4]byte) int32 {
	var result int32
	for i, v := range in {
		add := int32(v) << (i * 8)
		result |= add
	}
	return result
}

func Int16ToBytes(in int16) [2]byte {
	var result [2]byte
	result[0] = byte(in & 0xFF)
	result[1] = byte((in >> 8) & 0xFF)
	return result
}

func BytesToInt16(in [2]byte) int16 {
	var result int16
	for i, v := range in {
		add := int16(v) << (i * 8)
		result |= add
	}
	return result
}
