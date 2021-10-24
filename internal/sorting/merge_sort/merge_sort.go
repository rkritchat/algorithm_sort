package merge

func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	var middleIndex = len(data) / 2
	var lefArr = data[:int(middleIndex)]
	var righArr = data[int(middleIndex):]
	return merge(MergeSort(lefArr), MergeSort(righArr))
}

func merge(leftArr, rightArr []int) []int {
	var r []int
	var leftIndex = 0
	var rightIndex = 0

	for leftIndex < len(leftArr) && rightIndex < len(rightArr) { //valid until both of left and righ contain data in that index
		if leftArr[leftIndex] < rightArr[rightIndex] {
			r = append(r, leftArr[leftIndex])
			leftIndex += 1
		} else {
			r = append(r, rightArr[rightIndex])
			rightIndex += 1
		}
	}
	return append(append(r, leftArr[leftIndex:]...), rightArr[rightIndex:]...) //merge all remaingin on any left or right array
}
