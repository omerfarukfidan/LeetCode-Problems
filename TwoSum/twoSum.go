package TwoSum

func TwoSum(nums []int, target int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {

		for k := len(nums) - 1; k > -1; k-- {
			if i != k {
				sum = nums[i] + nums[k]
			}
			if sum == target {
				arr := []int{i, k}
				return arr
			}
		}
	}
	return nil
}
