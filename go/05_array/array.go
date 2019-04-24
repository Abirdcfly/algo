package _5_array

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

}

func (this *Array) Len() uint {
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (int, error) {
}

//插入数值到索引index上
func (this *Array) Insert(index uint, v int) error {
}

func (this *Array) InsertToTail(v int) error {
}

//删除索引index上的值
func (this *Array) Delete(index uint) (int, error) {
}

//打印数列
func (this *Array) Print() {
}
