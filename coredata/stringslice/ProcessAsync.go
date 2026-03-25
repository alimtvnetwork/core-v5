package stringslice

import "sync"

func ProcessAsync(
	processor func(index int, item any) string,
	items ...any,
) []string {
	if len(items) == 0 {
		return []string{}
	}

	list := make([]string, len(items))

	wg := sync.WaitGroup{}

	singleProcessorFunc := func(index int, item any) {
		list[index] = processor(index, item)

		wg.Done()
	}

	wg.Add(len(items))
	for i, item := range items {
		go singleProcessorFunc(i, item)
	}

	wg.Wait()

	return list
}
