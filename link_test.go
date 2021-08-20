package link

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink(t *testing.T) {

	type test struct {
		testFile string
		want     []Link
	}
	tests := []test{
		{
			testFile: "testfiles/ex1.html",
			want: []Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			testFile: "testfiles/ex2.html",
			want: []Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on Github!",
				},
			},
		},
		{
			testFile: "testfiles/ex3.html",
			want: []Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			testFile: "testfiles/ex4.html",
			want: []Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				},
			},
		},
	}

	for _, tc := range tests {
		bytes, err := ioutil.ReadFile(tc.testFile)
		assert.NoError(t, err, fmt.Sprintf("Error reading file %s", tc.testFile))
		links := ParseLinks(string(bytes))
		assert.Equal(t, tc.want, links, fmt.Sprintf("Failed on file %s", tc.testFile))
	}
}
