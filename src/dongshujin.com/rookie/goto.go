package rookie

import "fmt"

func GotoFunc() {
	i := 0

Here:
	fmt.Println("i=", i)
	if i += 1; i < 10 {
		goto Here
	}

}
