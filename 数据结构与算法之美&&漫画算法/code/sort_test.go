package code

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println()
	fmt.Println("TestBubbleSort------")
	a := []int{9, 11, 4, 15, 28, 10, 2}
	fmt.Println("init", a)
	fmt.Println(BubbleSort(a))
	fmt.Println()
}

func TestInsertSort(t *testing.T) {
	fmt.Println()
	fmt.Println("InsertSort------")
	a := []int{9, 11, 4, 15, 28, 10, 2}
	fmt.Println("init", a)
	fmt.Println(InsertSort(a))
}

func TestSelectSort(t *testing.T) {
	fmt.Println()
	fmt.Println("SelectSort------")
	a := []int{9, 11, 4, 15, 28, 10, 2}
	fmt.Println("init", a)
	fmt.Println(SelectSort(a))
}
