package querybuilder

import "strconv"

type limit struct {
	limit  int
	offset int
}

func (l *limit) format() string {
	if l.limit == 0 {
		return ""
	}

	str := " LIMIT " + strconv.Itoa(l.limit)

	if l.offset != 0 {
		str += " OFFSET " + strconv.Itoa(l.offset)
	}

	return str
}
