package reflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Ahmed"},
			ExpectedCalls: []string{"Ahmed"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name:          "struct with non string field",
			Input:         Person{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name:          "pointers to things",
			Input:         &Person{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "slices",
			Input: []Person{
				{"Chris", 33},
				{"Ahmed", 25},
			},
			ExpectedCalls: []string{"Chris", "Ahmed"},
		},
		{
			Name: "arrays",
			Input: [2]Person{
				{"Chris", 33},
				{"Ahmed", 25},
			},
			ExpectedCalls: []string{"Chris", "Ahmed"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}

		})
	}

	t.Run("maps", func(t *testing.T) {
		input := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		expectedCalls := []string{"Moo", "Baa"}
		var got []string
		walk(input, func(input string) {
			got = append(got, input)
		})

		if !(slices.Contains(got, "Moo") && slices.Contains(got, "Baa")) {
			t.Errorf("got %v, want %v", got, expectedCalls)
		}
	})

	t.Run("channels", func(t *testing.T) {
		input := make(chan Person)

		go func() {
			input <- Person{"Ahmed", 25}
			input <- Person{"Mohammed", 56}
			close(input)
		}()

		var got []string
		expectedCalls := []string{"Ahmed", "Mohammed"}

		walk(input, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expectedCalls) {
			t.Errorf("got %v, want %v", got, expectedCalls)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Person, Person) {
			return Person{"Ahmed", 25}, Person{"Mohammed", 56}
		}

		var got []string
		want := []string{"Ahmed", "Mohammed"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
