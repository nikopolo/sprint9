package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		// log.Println("incorrect value")
		return nil
	}
	randomList := make([]int, 0, size)

	for i := 0; i < size; i++ {
		randomList = append(randomList, rand.Int())
	}

	return randomList
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	if len(data) < CHUNKS { // если данных меньше, чем CHUNKS
		return maximum(data)
	}

	maxList := make([]int, CHUNKS) // слайс максимумов

	if len(data)%CHUNKS != 0 {
		remains := len(data) % CHUNKS // остаток
		size := len(data) / CHUNKS
		curIndex := 0 // индекс начала отсчета
		for i := 0; i < CHUNKS; i++ {
			beginIndex := curIndex // начальный индекс
			chunkSize := size      // размер CHUNK с данными
			if i < remains {       // увеличиваем размер, чтобы распределить остаток
				chunkSize++
			}
			endIndex := beginIndex + chunkSize // конечный индекс
			list := data[beginIndex:endIndex]  // получаемый слайс
			// fmt.Println(list)
			curIndex = endIndex
			wg.Add(1)

			go func(sl []int, index int) {
				defer wg.Done()

				maxList[i] = maximum(list)
			}(list, i)
		}

	} else {
		size := len(data) / CHUNKS // размер слайса
		for i := 0; i < CHUNKS; i++ {
			beginIndex := i * size            // начальный индекс
			endIndex := beginIndex + size     // конечный индекс
			list := data[beginIndex:endIndex] // получаемый слайс

			wg.Add(1)

			go func(sl []int, index int) {
				defer wg.Done()

				maxList[i] = maximum(list)
			}(list, i)

		}
	}
	wg.Wait()

	return maximum(maxList)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	listRand := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	begin := time.Now()
	max := maximum(listRand)
	elapsed := time.Since(begin).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	begin = time.Now()
	max = maxChunks(listRand)
	elapsed = time.Since(begin).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
