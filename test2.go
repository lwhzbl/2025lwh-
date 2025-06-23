package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s []int
	
    scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("请输入数字（每行一个，输入空行结束）：")
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("无效输入：%s，忽略该值\n", line)
			continue
		}
		s = append(s, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取输入时出错：%v\n", err)
	}

	fmt.Println("这一串数去重后为：")
	s = Deduplicate(s)
	for _, num := range s{
		fmt.Println(num)
	}
	
	


}

func Deduplicate(nums []int) []int{
	s := make(map[int]bool)
	result := []int{}
	for _,num:=range nums{
		if !s[num]{
			s[num]=true
			result=append(result,num)
		}
	}
	return result

}
