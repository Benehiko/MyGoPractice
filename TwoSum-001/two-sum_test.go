package TwoSum_001

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func TestRecursionTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(recursionTwoSum(nums, target))
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{2, 7, 11, 15}
		target := 9
		fmt.Println(twoSum(nums, target))
	}
}

func BenchmarkRecursionTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{2, 7, 11, 15}
		target := 9
		fmt.Println(recursionTwoSum(nums, target))
	}
}
