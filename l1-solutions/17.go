import "fmt"

func binarySearch(arr []int, x int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 9))
}
