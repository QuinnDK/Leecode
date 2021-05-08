package main

func isOneBitCharacter(bits []int) bool {
	for len(bits)>1{
		if bits[0]==1{
			bits=bits[2:]
		}else {
			bits=bits[1:]
		}
	}
	return len(bits)==1
}
