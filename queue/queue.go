package queue

type StringQueue []string

func (s *StringQueue) Push(toadd string) {
	*s = StringQueue(append([]string(*s), toadd))
}

func (s *StringQueue) Pull() string {
	tempS := []string(*s)
	*s = StringQueue(tempS[1:])
	return tempS[0]
}
