package utils

import "strconv"

func StringToUint(id string) uint {
	u64ID, _ := strconv.ParseUint(id, 10, 32)
	return uint(u64ID)
}
