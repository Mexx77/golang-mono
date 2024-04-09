package some_package

type stringSet interface {
	Add(key string)
}

type someService struct {
	stringSet stringSet
}

func New(stringSet stringSet) *someService {
	return &someService{stringSet}
}

func (s *someService) Work() {
	s.stringSet.Add("testing")
}
