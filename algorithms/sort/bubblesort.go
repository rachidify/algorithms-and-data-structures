package sort

// BubbleSort ...
func BubbleSort(list []int) {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
				swapped = true
			}
		}
	}
}
