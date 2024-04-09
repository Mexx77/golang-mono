package some_package

type stringSet interface {
	Add(key string)
}

type SomeService struct {
	stringSet stringSet
}

func New(setProvider stringSet) *SomeService {
	return &SomeService{setProvider}
}

func (s *SomeService) Work() {
	s.stringSet.Add("testing")
}
