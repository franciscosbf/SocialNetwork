package utils

import "testing"

func TestValidSetAny(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Unexpected panic: %v", err)
		}
	}()

	var i int

	SetAny(1, &i)

}

func TestInvalidSetAny(t *testing.T) {
	catchPanic := func(t *testing.T, from any, to AnyPtr) {
		defer func() {
			if err := recover(); err == nil {
				t.Error("Expecting panic")
			}
		}()

		SetAny(from, to)

	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestInvalidFrom",
			test: func(t *testing.T) {
				var j string
				l := "hello"
				catchPanic(t, &l, &j)
			},
		},
		{
			name: "TestInvalidTo",
			test: func(t *testing.T) {
				var j string
				catchPanic(t, "hello", j)
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}
