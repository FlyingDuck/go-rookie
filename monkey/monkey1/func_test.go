package monkey1

import (
	"bou.ke/monkey"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDiscuss(t *testing.T) {
	Discuss()
	fmt.Println("---------------------------------")
	monkey.Patch(LucySay, func() string {
		return "FakeLucy: Yoho..."
	})
	Discuss()
}

func TestFmtPrint(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}
