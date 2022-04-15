package main

import "testing"

func TestMain(t *testing.T) {
	main()
	if !test {
		t.Errorf("main failed!")
	}
}
