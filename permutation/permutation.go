package permutation

func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)
	return ret
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}
