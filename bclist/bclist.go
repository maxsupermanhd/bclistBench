package bclist

import "sync"

type bel[T any] struct {
	val  T
	sig  *sync.WaitGroup
	next *bel[T]
}

type BroadcastListener[T any] struct {
	current *bel[T]
}

func (b *BroadcastListener[T]) Wait() T {
	b.current.sig.Wait()
	v := b.current.next.val
	b.current = b.current.next
	return v
}

type BroadcastList[T any] struct {
	root *bel[T]
	head *bel[T]
}

func (b *BroadcastList[T]) Broadcast(val T) {
	b.head.next = &bel[T]{
		val:  val,
		sig:  &sync.WaitGroup{},
		next: nil,
	}
	b.head.next.sig.Add(1)
	release := b.head.sig
	b.head = b.head.next
	release.Done()
}

func (b *BroadcastList[T]) GetListener() *BroadcastListener[T] {
	return &BroadcastListener[T]{
		current: b.root,
	}
}

func NewBroadcastList[T any]() *BroadcastList[T] {
	root := &bel[T]{
		sig:  &sync.WaitGroup{},
		next: nil,
	}
	root.sig.Add(1)
	return &BroadcastList[T]{
		root: root,
		head: root,
	}
}
