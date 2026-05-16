package twosum

// O(n²) - brute force, check every pair
func twoSum_1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// O(n²) - map stores index→value, still requires range scan to find value
func twoSum_2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		for k, v := range m {
			if v == value && k != i {
				return []int{i, k}
			}
		}
	}

	return nil
}

// O(n) - map stores value→index, look up complement while iterating
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if v, ok := m[value]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
