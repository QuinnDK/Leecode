package main

func decode(encoded []int) []int {
	n := len(encoded)
	var sum1, sum2 int
	for i := 0; i < n; i = i + 2 {
		sum1 ^= encoded[i] //encoded[0]^encoded[2]^encoded[4]....==res[0]^res[1] ^ res[2]^res[3] ^ res[4]^res[5]...
	}
	for i := 0; i <= n+1; i++ {
		sum2 ^= i //0^1^2^3...
	}

	res := make([]int, n+1)
	res[n] = sum1 ^ sum2 //

	for i := n - 1; i >= 0; i-- {
		res[i] = encoded[i] ^ res[i+1]
	}
	return res
}
