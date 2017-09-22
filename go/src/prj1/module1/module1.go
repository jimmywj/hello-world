package module1

import "fmt"

func PublicFunc(num int) {
        num =1000
	fmt.Println("in PublicFunc",num)
}
func PublicFunc1(s string) {
        s = "kkkk"
	fmt.Println("in PublicFunc1",s)
}
func privateFunc() {
	fmt.Println("in privateFunc")
}
