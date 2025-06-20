package Handlers

import "strconv"

func parseUint(s string) uint {
	id, _ := strconv.ParseUint(s, 10, 32)
	return uint(id)
}
