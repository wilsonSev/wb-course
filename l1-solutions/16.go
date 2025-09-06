package main

import "fmt"

func quickSort(arr []int, low int, high int) {
	if low >= high { // базовое условие
		return
	}
	pivot := partition(arr, low, high) // разделяем
	// рекурсивно сортируем обе части
	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}

func partition(arr []int, low int, high int) int {
	i := low + 1
	j := high // high берем включительно
	// за pivot далее принимаем arr[low]
	for {
		// ищем первый элемент < pivot
		for i <= j && arr[low] > arr[i] {
			i++
		}
		// ищем первый элемент > pivot
		for i <= j && arr[low] < arr[j] {
			j--
		}
		// условие выхода
		if i >= j {
			break
		}
		// меняем местами элементы и идем дальше
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	// j указывает на элемент < pivot и стоит на его позиции
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

func main() {
	arr := []int{4, 5, 2, 0, 1, 35, 3, 1, 1, 2, 4, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
