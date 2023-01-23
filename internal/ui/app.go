package ui

type Page struct {
	Content string
	User    string
}

type Header struct {
	Href string
	Icon string
	Text string
}

type Content struct {
	Header string
	Body   string
}
