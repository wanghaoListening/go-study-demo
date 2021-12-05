package main

import "fmt"
import "sort"


//for range方式迭代的性能可能会更好一些，因为这种迭代可以保证不会出现数组越界的情形
func arrayExe()  {

	var s = [...]int{1,2,3,4}

	for i := range s{

		fmt.Println(i)
	}

	for i,v := range s{
		fmt.Println(i,v)
	}

	for i := 0;i<len(s);i++{

		fmt.Println(s[i])
	}

	//长度为0的数组在内存并不占用空间
	var d [0]int
	fmt.Println(d)

	c := make(chan struct{})

	go func() {
		fmt.Println("c")
		c <- struct{}{}
	}()
}

/**
 *和数组一样，内置的len函数返回切片中有效元素的长度，内置的cap函数返回切片容量大小，容量必须大于或等于切片的长度。
 */
func sliceExe() {

	var (

		g = make([]int,3) //有3个元素的切片, len和cap都为3
		h = make([]int,2,3) //有2个元素的切片, len为2, cap为3
		i = make([]int,0,3) //有0个元素的切片, len为0, cap为3
	)

	fmt.Println(g,h,i)

	var a []int

	a = append(a,1)
	a = append(a,1,2,3)
	//添加一个切片
	a = append(a,[]int{1,2,3}...)

	//在开头添加1个元素
	/**
	 * 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此，
	 * 从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
	 */
	a = append([]int{3},a...)

	//删除尾部一个元素
	a = a[:len(a)-1]

	//删除开头一个元素
	a = a[1:]


}

func trimSpace(s []byte) []byte {

	b := s[:0]

	for _, x := range s{

		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}


var sortArray = []int{4,2,5,7,2,1,88,1}

func sortNum(s []int) {
	// 以int方式给float64排序
	sort.Ints(s)
}




