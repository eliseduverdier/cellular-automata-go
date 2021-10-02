package parameters

import (
	"net/http"
	"reflect"
	"testing"
)

type Test struct {
	Query    string
	Expected Parameters
}

func TestGetParametersFromRequest(t *testing.T) {
	tests := []Test{
		{"/text?r=1&start=centered", Parameters{2, 1, 100, 100, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, "text"}},
		{"/text?r=110&o=1&s=2&w=4&h=4&start=custom&line=0001", Parameters{2, 1, 4, 4, []int{0, 0, 0, 1}, 110, "text"}},
		{"/image?r=70&o=1&s=3&w=4&h=4&start=centered", Parameters{3, 1, 4, 4, []int{0, 2, 1, 2}, 70, "image"}},
		{"/image?r=110&o=2&s=4&w=4&h=4&start=custom&line=1100", Parameters{4, 2, 4, 4, []int{1, 1, 0, 0}, 110, "image"}},
		{"/image?w=3&r=1", Parameters{2, 1, 3, 3, []int{1, 1, 1}, 1, "image"}},
	}

	for i := range tests {
		req, _ := http.NewRequest("GET", tests[i].Query, nil)
		if actual := GetFromRequest(req); !reflect.DeepEqual(actual, tests[i].Expected) {
			t.Errorf("parameters.GetFromRequest(%s) error: Got %v, expected %v.", tests[i].Query, actual, tests[i].Expected)
		}
	}
}

func TestGetParametersRandomFromRequest(t *testing.T) {
	query := "/image?w=3"
	req, _ := http.NewRequest("GET", query, nil)
	actualParams := GetFromRequest(req)
	if len(actualParams.Start) != 3 {
		t.Errorf("parameters.GetFromRequest(%s) error: Got first line of len %v, expected %v.", query, len(actualParams.Start), 3)
	}
	if actualParams.Rule == 0 {
		t.Errorf("parameters.GetFromRequest(%s) error: Expected a random rule number, got 0.", query)

	}
}

func TestCatchIncorrectRequest(t *testing.T) {
	// failingTests := []string{
	// 	"/image?r=a&o=Z&s=0&w=a&h=a&start=unknown&line=yop", //TODO catch errors
	// }
	// for i := range failingTests {
	// 	req, _ := http.NewRequest("GET", failingTests[i], nil)
	// 	actual, err := GetFromRequest(req)
	// 	if err == nil {
	// 		t.Errorf("Should have got a error from request %s, got %v", failingTests[i], actual)
	// 	}
	// }
}

func TestGetParametersFromShell(t *testing.T) {
	// tests := []Test{
	// 	{"-text ", Parameters{2, 1, 10, 10, []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, 110, "text"}},
	// 	{"-text -r=123 -w=4 -start=centered", Parameters{2, 1, 4, 4, []int{0, 0, 0, 1}, 123, "text"}},
	// 	{"-image -w=4", Parameters{2, 1, 4, 4, []int{0, 0, 0, 1}, 70, "image"}},
	// 	{"-image -w=4 -s=4 -o=2 -line=1100", Parameters{4, 2, 4, 4, []int{1, 1, 0, 0}, 110, "image"}},
	// }

	// for i := range tests {
	// 	input := strings.NewReader("-shell -r=123")
	// 	reader := io.Reader(input)
	// 	r := bufio.NewScanner(reader)
	// 	if actual := GetFromShell(); !reflect.DeepEqual(actual, tests[i].Expected) {
	// 		t.Errorf("parameters.GetFromShell(%s) error: Got %v, expected %v.", tests[i].Query, actual, tests[i].Expected)
	// 	}
	// }
}
