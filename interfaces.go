package gotg

import "time"

type Searcher interface {
	Search(pattern string, from time.Time, limit, offset uint64) ([]Message, error)
}

type Messager interface {
	SendMessage(msg string) error
}
