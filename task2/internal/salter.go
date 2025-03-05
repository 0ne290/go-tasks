package internal

import (
	"time"
	"strconv"
)

func Salt(source string, label string, id int, createdAt time.Time) string {
	return "part1OfStaticSalt_" + label + "_" + createdAt.String() + "_" + source + "_part2OfStaticSalt_" + strconv.Itoa(id)
}
