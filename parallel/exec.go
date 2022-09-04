package parallel

import "sync"

type Worker func()
type ChunkWorker func(first, last int)
type MirrorChunkWorker func(leftFirst, leftLast, rightFirst, rightLast int)

const DefaultChunkSize = 512

type chunkMetrics struct {
	nChunks        int
	chunkSize      int
	firstChunkSize int
}

func chunkPartitioner(first, last, requestedChunkSize int) chunkMetrics {
	/*
	 * This algorithm improves distribution of elements in chunks by avoiding
	 * small tail chunks. The leftover elements that do not fit neatly into
	 * the chunk size are redistributed to early chunks. This improves
	 * utilization of the processor's prefetch and reduces the number of
	 * tasks needed by 1.
	 */

	n := last - first
	nChunks := 0
	chunkSize := 0
	firstChunkSize := 0
	if n < requestedChunkSize {
		chunkSize = n
		firstChunkSize = n
		nChunks = 1
		return chunkMetrics{nChunks, chunkSize, firstChunkSize}
	}

	nChunks = (n / requestedChunkSize) + 1
	chunkSize = n / nChunks
	firstChunkSize = chunkSize
	nLeftoverItems := n - (nChunks * chunkSize)

	if nLeftoverItems == chunkSize {
		nChunks += 1
		return chunkMetrics{nChunks, chunkSize, firstChunkSize}
	} else if nLeftoverItems == 0 {
		firstChunkSize = chunkSize
		return chunkMetrics{nChunks, chunkSize, firstChunkSize}
	}

	nExtraItemsPerChunk := nLeftoverItems / nChunks
	nFinalLeftoverItems := nLeftoverItems - (nExtraItemsPerChunk * nChunks)

	chunkSize += nExtraItemsPerChunk
	firstChunkSize = chunkSize + nFinalLeftoverItems

	return chunkMetrics{nChunks, chunkSize, firstChunkSize}
}

func processChunk(metrics chunkMetrics, base, chunkIndex int, f ChunkWorker) {
	var thisChunkSize int
	if chunkIndex == 0 {
		thisChunkSize = metrics.firstChunkSize
	} else {
		thisChunkSize = metrics.chunkSize
	}
	var index int
	if chunkIndex == 0 {
		index = 0
	} else {
		index = (chunkIndex * metrics.chunkSize) + (metrics.firstChunkSize - metrics.chunkSize)
	}
	first := base + index
	last := first + thisChunkSize
	f(first, last)
}

func processMirrorChunk(metrics chunkMetrics, outerFirst, outerLast, chunkIndex int, f MirrorChunkWorker) {
	var thisChunkSize int
	if chunkIndex == 0 {
		thisChunkSize = metrics.firstChunkSize
	} else {
		thisChunkSize = metrics.chunkSize
	}
	var index int
	if chunkIndex == 0 {
		index = 0
	} else {
		index = (chunkIndex * metrics.chunkSize) + (metrics.firstChunkSize - metrics.chunkSize)
	}

	leftFirst := outerFirst + index
	leftLast := leftFirst + thisChunkSize

	rightLast := outerLast - index
	rightFirst := rightLast - thisChunkSize

	f(leftFirst, leftLast, rightFirst, rightLast)
}

func Fork(w1, w2 Worker) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		w1()
		wg.Done()
	}()

	go func() {
		w2()
		wg.Done()
	}()

	wg.Wait()
}

func Process(ws []Worker) {
	var wg sync.WaitGroup

	wg.Add(len(ws))
	for _, w := range ws {
		go func(wk Worker) {
			wk()
			wg.Done()
		}(w)
	}

	wg.Wait()
}

// ProcessChunks slices the range [first,last) up into chunks of
// DefaultChunkSize. Each of those chunks is processed by `f` in parallel.
// The amount of parallelism is determined by the Go runtime. This function
// will not return until all chunks have been processed.
func ProcessChunks(first, last int, f ChunkWorker) {
	var wg sync.WaitGroup
	metrics := chunkPartitioner(first, last, DefaultChunkSize)

	wg.Add(metrics.nChunks)
	for i := 0; i < metrics.nChunks; i++ {
		go func(chunk int) {
			processChunk(metrics, first, chunk, f)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// ProcessMirrorChunks slices the range [first, last-((last-first)/2)) into chunks of
// DefaultChunkSize/2 items. The chunks are taken from the beginning _and_ the
// end of the collection. The mirror chunk worker `f` is provided a range at the
// beginning and its mirror at the end of the range. This function will not
// return until all chunks have been processed.
func ProcessMirrorChunks(first, last int, f MirrorChunkWorker) {
	var wg sync.WaitGroup
	half := (last - first) / 2
	metrics := chunkPartitioner(first, first+half, DefaultChunkSize/2)

	wg.Add(metrics.nChunks)
	for i := 0; i < metrics.nChunks; i++ {
		go func(chunk int) {
			processMirrorChunk(metrics, first, last, chunk, f)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
