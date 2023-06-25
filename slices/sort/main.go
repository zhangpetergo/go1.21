package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 定义切片和要生成的数字数量
	var numbers []int
	numCount := 10

	// 生成随机数字并添加到切片中
	for i := 0; i < numCount; i++ {
		numbers = append(numbers, rand.Intn(100)) // 生成0到99之间的随机数
	}

	// 打印切片中的数字
	fmt.Println(numbers)

	// 默认按照升序排序
	slices.Sort(numbers)

	// 打印切片中的数字
	fmt.Println(numbers)
}
