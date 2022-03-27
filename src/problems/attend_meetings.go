package problems

import "sort"

type meetings [][]int

func (a meetings) Len() int           { return len(a) }
func (a meetings) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a meetings) Less(i, j int) bool { return a[i][0] < a[j][0] }

func AttendAllMeetings(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 1
	}

	sort.Sort(meetings(intervals))

	for i := 1; i < len(intervals); i++ {
		a := intervals[i-1]
		b := intervals[i]
		if a[1] > b[0] {
			return 0
		}
	}
	return 1
}
