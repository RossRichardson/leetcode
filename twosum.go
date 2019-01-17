/*
  TwoSum

  ref: https://leetcode.com/problems/two-sum/
  
  A naive implementation that beats 39% of submissions
  
  Complexity: 
    - Time: O(n^2) (even with x = i+1)
    - Space: O(1)
*/

import "fmt"

func main() {
    input := []int{2,7,11,15}
    result := twoSum(input, 9)
    fmt.Println("input:", input, "result:", result)
}

func twoSum(nums []int, target int) []int {
    l := len(nums)
	for i := 0; i < l-1; i++ {
		first := nums[i]
		for x := i + 1; x < l; x++ {
			if first+nums[x] == target {
				return []int{i, x}
			}
		}
	}
    return nil
}
