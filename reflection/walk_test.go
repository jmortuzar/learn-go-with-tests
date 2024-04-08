package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Matias"},
			[]string{"Matias"},
		}, {
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Matias", "Malaga"},
			[]string{"Matias", "Malaga"},
		}, {
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Matias", 30},
			[]string{"Matias"},
		}, {
			"nested fields",
			Person{"Matias", Profile{30, "Malaga"}},
			[]string{"Matias", "Malaga"},
		}, {
			"pointers to things",
			&Person{"Matias", Profile{30, "Malaga"}},
			[]string{"Matias", "Malaga"},
		}, {
			"slices",
			[]Profile{{30, "Malaga"}, {34, "Santiago"}},
			[]string{"Malaga", "Santiago"},
		}, {
			"arrays",
			[2]Profile{{30, "Malaga"}, {34, "Santiago"}},
			[]string{"Malaga", "Santiago"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "Malaga"}
			aChannel <- Profile{34, "Santiago"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Malaga", "Santiago"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{30, "Malaga"}, Profile{34, "Santiago"}
		}

		var got []string
		want := []string{"Malaga", "Santiago"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
