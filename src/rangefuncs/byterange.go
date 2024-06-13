package rangefuncs

import "fmt"

type Range struct {
	Start int64
	End   int64
}

func (r *Range) Next() (int64, bool) {
	if r.Start < r.End {
		value := r.Start
		r.Start++
		return value, true
	}
	return 0, false
}

func (r Range) Len() int64 {
	return r.End - r.Start
}

func calculateRanges(filesize, rangesize int64) []Range {
	var ranges []Range
	partSize := filesize / rangesize

	for i := int64(0); i < rangesize; i++ {

		start := i * partSize
		end := (i + 1) * partSize
		if i == rangesize-1 {
			end = filesize
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges
}

func printRanges(ranges []Range) {
	r := Range{}
	for i := range ranges {
		fmt.Printf("Range %d: Start: %d End: %d, Length: %d\n", i+1, r.Start, r.End-1, r.Len())
	}
}
