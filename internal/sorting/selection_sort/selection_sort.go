package selection

func SelctionSort(data []int) []int {
	if len(data) == 0 {
		return data
	}

	for i := range data {
		for j := range data {
			if data[i] < data[j] {
				tmp := data[i]
				data[i] = data[j]
				data[j] = tmp
			}
		}
	}
	return data
}
