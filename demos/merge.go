package demos

type Config struct {
	No int
}

func SplitAndMerge(nums []int) []int {
	lt10 := make([]int, 0)
	ge10 := make([]int, 0)
	for _, num := range nums {
		if num < 10 {
			lt10 = append(lt10, num)
		} else {
			ge10 = append(ge10, num)
		}
	}

	out := make([]int, 0, len(nums))
	out = append(out, lt10...)
	out = append(out, ge10...)
	return out
}

func ForLoopAndMerge(nums []int) []int {
	out := make([]int, 0, len(nums))
	for _, num := range nums {
		if num < 10 {
			out = append(out, num)
		}
	}

	for _, num := range nums {
		if num >= 10 {
			out = append(out, num)
		}
	}
	return out
}
