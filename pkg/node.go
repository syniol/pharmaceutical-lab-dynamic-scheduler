package pkg

type Node[T any] struct {
	PreviousNode *T
	NextNode     *T
}
