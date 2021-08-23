package stack

type StringStack []string

func (s *StringStack) Push(toadd string) {

	*s = StringStack(append([]string(*s), toadd))

}

func (s *StringStack) Pull() string {
	tempS := []string(*s)
	*s = StringStack(tempS[:len(tempS)-1])
	return tempS[len(tempS)-1]
}
