package some_package

import "fmt"

type stringSet interface {
	Add(key string)
	Contains(key string) bool
}

type someService struct {
	stringSet stringSet
}

func New(stringSet stringSet) *someService {
	return &someService{stringSet}
}

func (s *someService) Work() {
	key := "testing"
	fmt.Printf("adding %s to my set...\n", key)
	s.stringSet.Add(key)
	fmt.Printf("checking if %s is contained in the set: %t\n", key, s.stringSet.Contains(key))
}
