package sess

type Session struct {
	Stuff string
}

func (s *Session) Store(folder string) {
	s.Stuff = folder
}

func (s *Session) Clear() {
	s.Stuff = ""
}

func NewSession() *Session {
	s := Session{}
	return &s
}
