package main 

import (
	"testing"
	)

func TestGreeting(t *testing.T) {
	result := Greeting("Code.education Rocks!");

	runes := []rune(result);

	start := string(runes[:3]);
	end := string(runes[len(result)-4:]);

	if start != "<b>" {
		t.Errorf("Valor inicial diferente de <b>");
	} else if end != "</b>" {
		t.Errorf("Valor final diferente de </b>");
	}
}