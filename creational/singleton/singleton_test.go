package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to instance after calling ,not nil")
	}
	expectedCounter := counter1

	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		t.Errorf("After calling for the first time to count , the counet must be one but it is %d\n",currentCounter)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter{
		t.Error("expected same instance but it returned different instance")
	}
	currentCounter = counter2.AddOne()
	if currentCounter != 2 {
		t.Errorf("after calling AddOne using second counter ,the currnt count must be 2 but it is %d\n",currentCounter)
	}
}
