package strings

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	str := "Hello World"
	expected := "dlroW olleH"
	if Reverse(str) != expected {
		t.Errorf("expected result is %v but the actual result is %v", expected, Reverse(str))
	}
}

func TestToUpper(t *testing.T) {
	str1 := "Hello World$"
	expected1 := "HELLO WORLD$"
	if ToUpper(str1) != expected1 {
		t.Errorf("expected result is %v but the actual result is %v", expected1, ToUpper(str1))
	}
	str2 := "Hello, 世界"
	expected2 := "HELLO, 世界"
	if ToUpper(str2) != expected2 {
		t.Errorf("expected result is %v but the actual result is %v", expected2, ToUpper(str2))
	}
	fmt.Println(ToUpper(str2))
	str3 := "Hello,世j界"
	expected3 := "HELLO,世J界"
	if ToUpper(str3) != expected3 {
		t.Errorf("expected result is %v but the actual result is %v", expected3, ToUpper(str3))
	}
	fmt.Println(ToUpper(str3))
}

func TestVovelsCout(t *testing.T) {
	str1 := "Hello World"
	var expected1 uint = 3
	if expected1 != GetVovelCount(str1) {
		t.Errorf("expected result is %v but the actual result is %v", expected1, GetVovelCount(str1))
	}
	str2 := "why"
	var expected2 uint = 0
	if expected2 != GetVovelCount(str2) {
		t.Errorf("expected result is %v but the actual result is %v", expected2, GetVovelCount(str2))
	}
}

func TestVovelsCoutBDD(t *testing.T) {
	str1 := "Hello World"
	var expected1 uint = 3
	t.Log("Given a string", str1, "I want to test how many vovels are there")

	if expected1 != GetVovelCount(str1) {
		t.Fail()
		t.Log("expected result is", expected1, "but the actual result is", GetVovelCount(str1), "hence it is failure")
	} else {
		t.Log("if the actual result is eaual to the expected result it is successful.actual result:", GetVovelCount(str1), "expected result:", expected1, "since both are equal this test is successful")
	}
}
