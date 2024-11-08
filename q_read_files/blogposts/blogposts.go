package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(filesystem, f.Name())
		if err != nil {
			return nil, err //todo: needs clarification, should we totally fail if one file fails? or just ignore ?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, f string) (Post, error) {
	postFile, err := filesystem.Open(f)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
