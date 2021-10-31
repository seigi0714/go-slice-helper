package slice_helper

import (
	"errors"
	"fmt"
	"reflect"
)

// 任意のインデックスに値を挿入する便利関数
func Insert(slice interface{}, v interface{}, index int) (interface{}, error) {
	if !isSlice(slice) {
		return nil, errors.New("1st Aug is must be slice")
	}
	err := typeCheck(slice, v)
	if err != nil {
		return nil, err
	}

	sliceRv := reflect.ValueOf(slice)
	valueRv := reflect.ValueOf(v)

	// スライスの長さより大きいインデックスを指定した場合,スライスの最後に要素追加
	if sliceRv.Len() < index {
		index = sliceRv.Len()
	}
	// sliceで言う　s[:index]
	firstSlice := sliceRv.Slice(0, index)
	latterSlice := sliceRv.Slice(index, sliceRv.Len())
	latterSlice = reflect.AppendSlice(toSlice(valueRv), latterSlice)

	returnSlice := reflect.AppendSlice(firstSlice, latterSlice)
	return returnSlice.Interface(), nil
}

// 任意のインデックスに値を挿入する便利関数
func Reduce(slice interface{}, index int) (interface{}, error) {
	if !isSlice(slice) {
		return nil, errors.New("1st Aug is must be slice")
	}

	sliceRv := reflect.ValueOf(slice)
	// index == 0 の場合先頭のみ削除
	if index == 0 {
		returnSlice := sliceRv.Slice(1, sliceRv.Len())
		return returnSlice.Interface(), nil
	}
	// index == スライス数 の場合末尾のみ削除
	if index == sliceRv.Len()-1 {
		returnSlice := sliceRv.Slice(0, sliceRv.Len()-1)
		return returnSlice.Interface(), nil
	}
	// index >= スライス数 エラー
	if index >= sliceRv.Len()-1 {
		return nil, fmt.Errorf("not index out of range [%d] with length %d", index, sliceRv.Len())
	}

	firstSlice := sliceRv.Slice(0, index)
	latterSlice := sliceRv.Slice(index+1, sliceRv.Len())
	returnSlice := reflect.AppendSlice(firstSlice, latterSlice)
	return returnSlice.Interface(), nil
}

func toSlice(rv reflect.Value) reflect.Value {
	sliceType := reflect.SliceOf(rv.Type())
	s := reflect.MakeSlice(sliceType, 0, 0)
	return reflect.Append(s, rv)
}

// 渡されたスライスが本当にスライスであるかのチェック
func isSlice(slice interface{}) bool {
	arrType := reflect.TypeOf(slice)
	kind := arrType.Kind()
	return kind == reflect.Slice
}

// お互いの方が一致するかのチェック
func typeCheck(slice interface{}, v interface{}) error {
	vSliceType := reflect.SliceOf(reflect.TypeOf(v))
	vSlice := reflect.MakeSlice(vSliceType, 1, 1)
	if reflect.TypeOf(slice) != vSlice.Type() {
		return errors.New("Did not match insert value type")
	}
	return nil
}
