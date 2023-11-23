package sort

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// low and high is in range [low, high], not [low, high)
func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	index := Partition(nums, low, high)
	quickSort(nums, low, index-1)
	quickSort(nums, index+1, high)
}

// low and high is in range [low, high], not [low, high)
func Partition(nums []int, low, high int) int {
	p := nums[low]
	for low < high {
		for low < high && nums[high] >= p {
			high--
		}
		if low == high {
			break
		} else {
			nums[low] = nums[high]
			low++
		}

		for low < high && nums[low] <= p {
			low++
		}
		if low == high {
			break
		} else {
			nums[high] = nums[low]
			high--
		}
	}

	nums[low] = p
	return low
}
