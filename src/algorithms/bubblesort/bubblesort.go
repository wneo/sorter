package bubblesort

func BubbleSort(values []int) {
	flag := true
	lenth := len(values)
	for i := 0; i < lenth-1; i++ {
		for j := 0; j < lenth-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			} // end if
		} // end for j = ...
		if flag == true {
			break
		}
	} // end for i = ...
}
