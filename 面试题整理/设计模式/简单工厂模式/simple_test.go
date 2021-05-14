package 简单工厂模式

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi,Tom" {
		t.Fatal("type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello,Tom" {
		t.Fatal("type1 test fail")
	}
}
