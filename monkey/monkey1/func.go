package monkey1

import "fmt"

func Discuss() {
	fmt.Println(BobSay())
	fmt.Println(LucySay())
}

func BobSay() string {
	return "I'm Bob, a ba a ba..."
}

func LucySay() string {
	return "I'm Lucy, wu a wu a..."
}
