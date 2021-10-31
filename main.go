package main

import (
	"fmt"
	"go-slice-helper/slice_helper"
)

const addIndex int = 1
const addText string = "testX"

const reduceIndex int = 1

func main() {
	slice := []string{"test1", "test2", "test3", "test4"}

	// スライスのaddIndex番目に要素を追加
	addSlice, err1 := slice_helper.Insert(slice, addText, addIndex)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	addSlice = addSlice.([]string)
	fmt.Println("AddSliceMethod :: ", addSlice)

	// スライスのreduceIndex番目の要素を削除
	reduceSlice, err2 := slice_helper.Reduce(slice, reduceIndex)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Println("ReduceSliceMethod :: ", reduceSlice)
}
