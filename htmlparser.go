package htmlparser

import ("golang.org/x/net/html"
        "strings"
)



func Htmlparser(inputHtml, domain string) (urls []string){

doc, err := html.Parse(strings.NewReader(inputHtml))
if err != nil {
}
var f func(*html.Node)
f = func(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" && strings.Contains(a.Val, domain) {
                urls = append(urls, a.Val)
                break
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        f(c)
    }
}
f(doc)

return urls
}
