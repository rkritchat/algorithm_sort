package bubble

//return O(n^2) Osquare
func BubbleSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := len(data) - 1; j >= i+1; j-- {
			if data[j] < data[j-1] {
				tmp := data[j]
				data[j] = data[j-1]
				data[j-1] = tmp
			}
		}
	}
	return data
}
