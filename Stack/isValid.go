package Stack

type Stack struct {
	stack []any
}

func NewStack() *Stack {
	return &Stack{
		[]any{},
	}
}

func (s *Stack) Push(val any) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() any {
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

func isValid(s string) bool {
	if len(s)%2 != 0 { // s 长度必须是偶数
		return false
	}
	stack := NewStack()
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
		} else {
			if stack.Empty() {
				return false
			}
			switch c {
			case ')':
				if stack.Pop() != '(' {
					return false
				}
			case ']':
				if stack.Pop() != '[' {
					return false
				}
			case '}':
				if stack.Pop() != '{' {
					return false
				}
			}
		}
	}
	if !stack.Empty() {
		return false
	}
	return true
}
