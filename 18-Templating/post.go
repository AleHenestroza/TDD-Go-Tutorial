package blogrenderer

import "strings"

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
