package main

import (
	"fmt"
	"testing"
)

func TestNumber_Multi(t *testing.T) {
	type fields struct {
		a int
		b int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{name: "1", fields: fields{a: 0, b: 2}, want: 0},
		{name: "2", fields: fields{a: 3, b: -2}, want: -6},
		{name: "3", fields: fields{a: 5, b: 9}, want: 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := &Number{
				a: tt.fields.a,
				b: tt.fields.b,
			}
			if got := num.Multi(); got != tt.want {
				t.Errorf("Multi() = %v, want %v", got, tt.want)
			}
		})
	}
}
// 使用案例
func ExampleNumber_Multi() {
	num := Number{
		a: 2,
		b: 6,
	}
	fmt.Println(num.Multi())

	// Output:
	// 12
}

// 基准测试
func BenchmarkNumber_Multi(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		num := &Number{
			a: 33,
			b: 20,
		}
		if num.Multi() != 660 {
			b.FailNow()
		}
	}
}