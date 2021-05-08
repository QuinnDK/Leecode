package main

/*
思路

将 nn 视作一个长为 3232 的二进制串，从低位往高位枚举 nn 的每一位，将其倒序添加到翻转结果 \textit{rev}rev 中。

代码实现中，每枚举一位就将 nn 右移一位，这样当前 nn 的最低位就是我们要枚举的比特位。当 nn 为 00 时即可结束循环。

需要注意的是，在某些语言（如 \texttt{Java}Java）中，没有无符号整数类型，因此对 nn 的右移操作应使用逻辑右移。

*/

func reverseBits(n uint32)(rev uint32){
	for i:=0;i<32&&n>0;i++{
		rev |=n&1<<(31-i)
		n>>=1
	}
	return
}