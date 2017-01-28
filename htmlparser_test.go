package htmlparser

import ("testing"
        "strings"
)

func TestHtmlParser(t *testing.T) {

foo := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`

  cases := []struct {
		in , want string
	}{
		{foo, "foo/bar/baz"},
    {"", ""},
	}

	for _, c := range cases {
		got:= Htmlparser(c.in)
		if !strings.Contains(got,c.want) {
			t.Errorf("Htmlparser(%q) returned %q, wanted  %q", c.in, got, c.want)
    }
	}
}
