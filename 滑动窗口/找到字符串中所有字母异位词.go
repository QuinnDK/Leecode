package 滑动窗口

func findAngrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	cn1 := [26]int{}
	cn2 := [26]int{}

	for i := 0; i < len(p); i++ {
		cn1[p[i]-'a']++
		cn2[s[i]-'a']++
	}
	res := []int{}

	for i := 0; i < len(s)-len(p); i++ {
		if cn1 == cn2 {
			res = append(res, i)
		}
		cn2[s[i]-'a']--
		if i+len(p) < len(s) {

		}
	}
}
