package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(f io.Reader) (Post, error) {
	scanner := bufio.NewScanner(f)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	post := Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore one line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
