package blog

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: dev, go
----
This is my body
Second body line`

		secondBody = `Title: Post 2
Description: Description 2
Tags: test`
	)

	testFs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello world2.md": {Data: []byte(secondBody)},
	}

	got, err := newPostsFromFs(testFs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, got[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"dev", "go"},
		Body:        "This is my body\nSecond body line",
	})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestFailingFileSystem(t *testing.T) {

	failingFs := stubFailingFs{}
	_, err := newPostsFromFs(failingFs)

	if err == nil {
		t.Errorf("Expected an error when opening the failing filesystem")
	}
}

type stubFailingFs struct {
}

func (s stubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("Failed to open file")
}
