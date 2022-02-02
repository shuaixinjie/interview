package designpattren

import "fmt"

const (
	football = iota
	swimming
)

type PlayStrategy interface {
	play(a, b int) error
}

type playFunc func(a, b int) error

func (f playFunc) play(a, b int) error {
	return f(a, b)
}

func New3(t int) (PlayStrategy, error) {
	if s, ok := strategy[t]; ok {
		return s, nil
	}
	return nil, nil
}

var strategy = map[int]PlayStrategy{
	football: playFunc(playFootball),
	swimming: playFunc(goSwimming),
}

func goSwimming(a, b int) error {
	fmt.Println("swimming...")
	return nil
}

func playFootball(a, b int) error {
	fmt.Println("play football...")
	return nil
}
