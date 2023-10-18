package gt

import "strconv"

func FmtMembers(array []string) []string {
	var Member []string
	if len(array) != 1 {
		for i := 0; i < len(array); i++ {

			Membermiddles := "(" + strconv.Itoa(i+1) + ") " + array[i] + "  "
			Member = append(Member, Membermiddles)

		}
		array = Member
	}
	return array
}
