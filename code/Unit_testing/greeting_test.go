package Unit_testing

import (
	"fmt"
	"testing"
)

func Test_SayHello_ValidArgument(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Hello %s", name)
	result := sayHello(name)

	if result != expected {
		t.Errorf("\"sayHello('%s')\" FAILED, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayHello('%s')\" SUCCEDED, expected -> %v, got -> %v", name, expected, result)
	}
}

func Test_SayHello_ValidArgument_array_inputs(t *testing.T) {
	inputs := []struct {
		name   string
		result string
	}{
		{name: "Yemeksepeti", result: "Hello Yemeksepeti"},
		{name: "Banabi", result: "Hello Banabi"},
		{name: "Yemek", result: "Hello Yemek"},
	}

	for _, item := range inputs {

		result := sayHello(item.name)
		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v", item.name, item.result, result)
		}
	}
}
