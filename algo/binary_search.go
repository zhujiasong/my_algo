package algo

func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchLeft(nums []int, target int) int {
	low, high := 0, len(nums)

	for low < high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		}
	}

	return low
}
