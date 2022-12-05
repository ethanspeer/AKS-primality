package main

import "testing"

func Test1(t *testing.T) {
	n := 3
	got := aks(n)
	want := "prime"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}

func Test2(t *testing.T) {
	n := 89
	got := aks(n)
	want := "prime"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}

func Test3(t *testing.T) {
	n := 631
	got := aks(n)
	want := "prime"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}

func Test4(t *testing.T) {
	n := 1303
	got := aks(n)
	want := "prime"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}

func Test5(t *testing.T) {
	n := 555
	got := aks(n)
	want := "composite"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}

func Test6(t *testing.T) {
	n := 10000
	got := aks(n)
	want := "composite"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}

func Test7(t *testing.T) {
	n := 1000000001913
	got := aks(n)
	want := "composite"
	
	if got != want {
		t.Errorf("Testing n=%d you find %q when you should find %q\n", n, got, want)
	}
}