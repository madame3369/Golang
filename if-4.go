package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000
	
	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈이...")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자.")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}
}

// package main

// import "fmt"

// func main() {
// 	foodcost = 40000
// 	friends = ""
// 	f.c=0

// 	if foodcost > 50000 && friends = "부자" {
// 		fmt.Println("신발끈을 묶는다")
// 	} else if foodcost >= 30000 || foodcost <= 50000 && f.c > 3 {
// 		fmt.Println("3명이 신발끈을 묶는다")
// 	} else foodcost < 30000 {
// 		fmt.Println("내가 낸다")
// 	}
// }