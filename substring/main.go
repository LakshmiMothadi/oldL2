package main

import "fmt"

func substring(s string) {

	fmt.Println("string:")
	fmt.Scan(&s)

	for i:=0 ;i<len(s);i++{
		for j:=i+1;j<len(s);j++ {
			fmt.Println(i,j)
			//l= s[i]

		}
	}

}


func main(){
	var s string
	substring(s)
	var str string
	fmt.Println("Enter a  string to find longSubstring:")
	fmt.Scan(&str)
	lenOfNoRepeatingSubStr(str)
}