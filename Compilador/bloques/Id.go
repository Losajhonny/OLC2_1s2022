package bloques

import "strconv"

var id int = 0

func GetId() string {
	id = id + 1
	return strconv.Itoa(id)
}
