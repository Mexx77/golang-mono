package lib1

import "fmt"

type Helloer struct {
}

func (h Helloer) Hello(name string) string {
	return fmt.Sprintf("Nice to meet you, %s", name)
}
