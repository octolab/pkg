package sequence

// Range returns a specific slice for range iteration on another slice.
//
//  batch, entities := 100, entity.Fetch() // []entity.Model
//  for step, end := range sequence.Range(len(entities), batch) {
//  	service.Process(entities[batch*step:end])
//  }
//
func Range(size, batch int) []int {
	count := size / batch
	if size%batch != 0 {
		count++
	}
	batches := make([]int, count)
	for i := 0; i < count; i++ {
		border := (i + 1) * batch
		if border > size {
			border = size
		}
		batches[i] = border
	}
	return batches
}
