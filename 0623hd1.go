package main

import "fmt"

func main() {
	var n int
	s:= []int{}
	fmt.Println("请输入一个数n：")
	fmt.Scanln(&n)
	fmt.Println("n以内的质数为：")
	s = Prime(n)
	for i:= 0;i<len(s);i++{
		fmt.Println(s[i])
	}
	

}
func Prime(n int) []int {
	s:= []int{}
	for i:= 0;i<=n;i++{
		if(isPrime(i)){
			s=append(s,i)
		}
	}
	return s

 }
func isPrime(n int) bool{
	if n<=1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i:=2;i*i<=n;i++{
		if(n % i == 0){
			return false
		}
	}
	return true
}
