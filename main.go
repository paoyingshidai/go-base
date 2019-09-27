package main

import "workspace2/concurrent"

func main() {

	//_,numb,strs := function.Numbers() //只获取函数返回值的后两个
	//fmt.Println(numb,strs)
	//
	//array.PrintArray()
	//
	//point.MyPoint()
	//
	//_range.RangeTest()
	//
	//mystruct.PrintBook()
	//
	//mystruct.SetBookProperties()
	//
	//mymap.SetAndPrintMap()
	//
	//mymap.MapDelete()

	//fmt.Println(recursion.Factorial(12))

	//typechange.TypeChange()

	//go concurrent.Sleep("world")
	//concurrent.Sleep("hello")

	//concurrent.TestSum()

	concurrent.TestClose()

	// error

	// 正常情况
	//if result, errorMsg := myerror.Divide(100, 10); errorMsg == "" {
	//	fmt.Println("100/10 = ", result)
	//}
	//// 当被除数为零的时候会返回错误信息
	//if _, errorMsg := myerror.Divide(100, 0); errorMsg != "" {
	//	fmt.Println("errorMsg is: ", errorMsg)
	//}
}
