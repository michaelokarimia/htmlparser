package htmlparser

import ("testing"
        "strings"
)

func TestHtmlParser(t *testing.T) {

  cases := []struct {
		in , want string
	}{
		{foo, "foo/bar/baz"},
    {htmlsample, "github.com"},
    {htmlsample, "golang"},
	}

	for _, c := range cases {
		got:= Htmlparser(c.in)
		if !strings.Contains(got,c.want) {
			t.Errorf("Htmlparser(%q) returned %q, wanted  %q", c.in, got, c.want)
    }
	}
}



  const foo = `<p>Links:</p><ul><li><a href="foo">
  Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
  const htmlsample =`<!DOCTYPEhtml><htmllang="en">
  <headprofile="http://a9.com/-/spec/opensearch/1.1/"><metacharset="utf-8">
  <metaname="viewport"content="width=device-width,initial-scale=1.0">
  <link href="/-/bootstrap.min.css?v=8bec1bba3e23ecba22cffb197a2d440af410b15d"rel="stylesheet">
  <link href="/-/site.css?v=7d81f4104c89dbe376345f6bfe3e62b4e40d3d06"rel="stylesheet">
  <title>html-GoDoc</title>
  <metaname="description"content="PackagehtmlimplementsanHTML5-complianttokenizerandparser.">
  </head><body><navclass="navbarnavbar-default"role="navigation"
  <divclass="container"><divclass="navbar-header">
  <buttontype="button"class="navbar-toggle"data-toggle="collapse"data-target=".navbar-collapse"
  <spanclass="sr-only">Togglenavigation</span><spanclass="icon-bar"></span
  <spanclass="icon-bar"></span><spanclass="icon-bar"></span></button
  <a class="navbar-brand"href="/"><strong>GoDoc</strong></a></div
  <divclass="collapsenavbar-collapse"><ulclass="navnavbar-nav"><li>
  <a href="/">Home</a></li><li><a href="/-/about">About</a></li></ul
  <formclass="navbar-navnavbar-formnavbar-right"id="x-search"action="/"role="search">
  <inputclass="form-control"id="x-search-query"type="text"name="q"placeholder="Search"
  </form></div></div></nav><divclass="clearfix"id="x-projnav">
  <a href="https://github.com/golang/net/"><strong>net:</strong></a
  <a href="/golang.org/x/net">golang.org/x/net</a><spanclass="text-muted"
  /</span><spanclass="text-muted">html</span></body>`
