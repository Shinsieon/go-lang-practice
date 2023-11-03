package main

type set map[int]struct{}

func (s set) Add(v int) {
	s[v] = struct{}{}
}
func (s set) Remove(v int) {
	delete(s, v)
}
func (s set) Contain(v int) bool {
	_, ok := s[v]
	return ok
}
func (s set) Length() int {
	return len(s)
}
func solution(elements []int) int {

	answer := set{}
	elements_ := append(elements, elements[0:len(elements)-1]...)
	for startIdx := 0; startIdx < len(elements); startIdx++ {
		sum := 0
		for i := 0; i < len(elements); i++ {
			sum += elements_[startIdx+i]
			answer.Add(sum)
		}
	}
	return answer.Length()

}