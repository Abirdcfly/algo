package _5_array

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Author: leo
 */

type Array struct {
	data   []int
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0}
}

func (this *Array) Len() uint {
	return this.length
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= this.length {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("越界")
	}
	return this.data[index], nil
}

//插入数值到索引index上
func (this *Array) Insert(index uint, v int) error {
	if int(index) > cap(this.data) {
		return errors.New("插入位置超过cap")
	}
	if this.length > 0 {
		for i := this.length - 1; i >= index; i-- {
			this.data[i+1] = this.data[i]
		}
	}

	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) InsertToTail(v int) error {
	return nil
}

//删除索引index上的值
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("越界")
	}
	num := this.data[index]
	for i := int(index); i < cap(this.data)-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return num, nil
}

//打印数列
func (this *Array) Print() {
	fmt.Printf("%v\n", this.data)
}
