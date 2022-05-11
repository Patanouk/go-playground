package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

func newPostsFromFs(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPostFromFile(filesystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPostFromFile(filesystem fs.FS, fileName string) (Post, error) {
	postFile, err := filesystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return NewPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func NewPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tag string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tag)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagSeparator), ", "),
		Body:        extractBody(scanner),
	}, nil
}

func extractBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}

type Post struct {
	Title, Description, Body string
	Tags                     []string
}
