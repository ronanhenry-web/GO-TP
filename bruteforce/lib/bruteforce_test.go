package lib_test

import (
	"bruteforce/lib"
	"testing"
)

func TestAkali(t *testing.T) {
	expected := "@kAl1"
	got := lib.FindPasswordNaive("@kAl1")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestShen(t *testing.T) {
	expected := "Sh3n"
	got := lib.FindPasswordNaive("Sh3n")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestZed(t *testing.T) {
	expected := "z3D"
	got := lib.FindPasswordAsync("z3D")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func BenchmarkZedNaive(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordNaive("z3d")
	}
}

func BenchmarkZedCombination(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordByCombination("z3d")
	}
}

func BenchmarkHeimerdingerCombination(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordByCombination("H31m3rd1ng3r")
	}
}

func BenchmarkHeimerdingerAsync(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordAsync("H31m3rd1ng3r")
	}
}

func TestHeimerdingerAsAsync(t *testing.T) {
	passwords := []string{
		"z3D",
		"Sh3n",
		"@kAl1",
	}

	for _, password := range passwords {
		got := lib.FindPasswordByCombinationAsync(password)
		if got != password {
			t.Errorf("Expected %v but got %v", password, got)
		}
	}
}
