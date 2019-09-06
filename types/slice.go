package types

func Covert(s1, s2 []int) {
	for _, value := range s1 {
		for i := range s2 {
			s2[i] = value
		}
	}
}
