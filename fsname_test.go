package fssup

import (
	"testing"
)

func TestFSName(t *testing.T) {

	tests := []struct{
		Dir string
		Name string
		Expected string
	}{
		{
			Dir:  "something",
			Name: "something/apple",
			Expected:       "apple",
		},
		{
			Dir:  "something",
			Name: "something/apple/banana",
			Expected:       "apple/banana",
		},
		{
			Dir:  "something",
			Name: "something/apple/banana/cherry",
			Expected:       "apple/banana/cherry",
		},



		{
			Dir:  "something/",
			Name: "something/apple",
			Expected:       "apple",
		},
		{
			Dir:  "something/",
			Name: "something/apple/banana",
			Expected:       "apple/banana",
		},
		{
			Dir:  "something/",
			Name: "something/apple/banana/cherry",
			Expected:       "apple/banana/cherry",
		},



		{
			Dir:  "something",
			Name: "something/apple/",
			Expected:       "apple/",
		},
		{
			Dir:  "something",
			Name: "something/apple/banana/",
			Expected:       "apple/banana/",
		},
		{
			Dir:  "something",
			Name: "something/apple/banana/cherry/",
			Expected:       "apple/banana/cherry/",
		},



		{
			Dir:  "something/",
			Name: "something/apple/",
			Expected:       "apple/",
		},
		{
			Dir:  "something/",
			Name: "something/apple/banana/",
			Expected:       "apple/banana/",
		},
		{
			Dir:  "something/",
			Name: "something/apple/banana/cherry/",
			Expected:       "apple/banana/cherry/",
		},



		{
			Dir:  "something",
			Name: "something/apple.html",
			Expected:       "apple.html",
		},
		{
			Dir:  "something",
			Name: "something/apple/banana.html",
			Expected:       "apple/banana.html",
		},
		{
			Dir:  "something",
			Name: "something/apple/banana/cherry.html",
			Expected:       "apple/banana/cherry.html",
		},



		{
			Dir:  "something/",
			Name: "something/apple.html",
			Expected:       "apple.html",
		},
		{
			Dir:  "something/",
			Name: "something/apple/banana.html",
			Expected:       "apple/banana.html",
		},
		{
			Dir:  "something/",
			Name: "something/apple/banana/cherry.html",
			Expected:       "apple/banana/cherry.html",
		},
	}

	for testNumber, test := range tests {

		actual, err := fsName(test.Dir, test.Name)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DIR: %q", test.Dir)
			t.Logf("Name: %q", test.Name)
			continue
		}

		{
			expected := test.Expected

			if  expected != actual{
				t.Errorf("For test #%d, the actual 'file-system name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL :  %q", actual)
				t.Logf("DIR: %q", test.Dir)
				t.Logf("Name: %q", test.Name)
				continue
			}
		}
	}
}
