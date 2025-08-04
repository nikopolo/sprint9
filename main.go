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
	src := rand.Int()
	for i := 0; i < size; i++ {
		randomList = append(randomList, src)
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

	size := len(data) / CHUNKS     // размер слайса
	maxList := make([]int, CHUNKS) // слайс максимумов

	for i := 0; i < CHUNKS; i++ {
		beginIndex := i * size            // начальный индекс
		endIndex := beginIndex + size     // конечный индекс
		list := data[beginIndex:endIndex] // получаемый слайс

		wg.Add(1)

		go func(sl []int, index int) {
			defer wg.Done()

			max := maximum(list)
			maxList[i] = max
		}(list, i)

	}

	wg.Wait()

	maxResult := maximum(maxList)

	return maxResult
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
