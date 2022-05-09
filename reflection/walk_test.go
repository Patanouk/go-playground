package reflection

import (
	"reflect"
	"testing"
)

func Test_walk(t *testing.T) {
	tests := []struct {
		name          string
		args          interface{}
		expectedCalls []string
	}{
		{"test one string interface",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{"test two string interface",
			struct {
				Name     string
				Lastname string
			}{"Chris", "Random"},
			[]string{"Chris", "Random"},
		},
		{"test string and int interface",
			struct {
				Name string
				Age  int
			}{"Chris", 10},
			[]string{"Chris"},
		},
		{"test nested structure",
			Person{"Chris", Profile{10, "Paris"}},
			[]string{"Chris", "Paris"},
		},
		{"test pointer",
			&Person{"Chris", Profile{10, "Paris"}},
			[]string{"Chris", "Paris"},
		},
		{"test slices",
			[]Profile{
				{33, "Paris"},
				{10, "London"},
			},
			[]string{"Paris", "London"},
		},
		{"test arrays",
			[2]Profile{
				{33, "Paris"},
				{10, "London"},
			},
			[]string{"Paris", "London"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]string, 0)
			walk(tt.args, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.expectedCalls) {
				t.Errorf("got %q calls, want %q", got, tt.expectedCalls)
			}
		})
	}

	t.Run("test maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Foo")
		assertContains(t, got, "Baz")
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{10, "Paris"}
			aChannel <- Profile{10, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Paris", "London"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q calls, want %q", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

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

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
