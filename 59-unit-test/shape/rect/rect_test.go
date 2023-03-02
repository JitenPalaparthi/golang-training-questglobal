package rect

import (
	"testing"
	//"time"
)

// actual result against the expected result
func TestArea(t *testing.T) {
	var expected float32 = 155.035
	r1 := new(Rect)
	r1.L = 10.10
	r1.B = 15.35
	//time.Sleep(time.Second * 6)
	if r1.Area() != expected {
		t.Errorf("test Area method actual result : %v and expected result : %v", r1.Area(), expected)
	}
}

func TestPerimeter(t *testing.T) {
	var expected float32 = 50.9
	r1 := new(Rect)
	r1.L = 10.10
	r1.B = 15.35
	//time.Sleep(time.Second * 6)
	if r1.Perimeter() != expected {
		t.Errorf("test Perimeter method actual result : %v and expected result : %v", r1.Area(), expected)
	}
}
