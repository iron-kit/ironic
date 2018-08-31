package file

import (
	"os"
)

// FileInfo 文件信息
type FileInfo struct {
	os.FileInfo
	Path string
}

// FileSlice 文件切片
type FileSlice []FileInfo

// ModTime升序
// sort.Interface
// Len is the number of elements in the collection.
func (files FileSlice) Len() int {
	return len(files)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (files FileSlice) Less(i, j int) bool {
	a := files[i]
	b := files[j]
	if a.ModTime().UnixNano() < b.ModTime().UnixNano() {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (files FileSlice) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}
