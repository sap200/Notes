// package notes
package notes

import "strings"

// Note is a struct with content and index of the note in a Note slice
type Note struct {
	Index   int
	Content string
}

// New Note generates New Note
func NewNote(content string, index int) *Note {
	return &Note{Content: content, Index: index}
}

// Get Title returns the title of a note
// Title of a note is first word of the note.Content
func (n *Note) GetTitle() string {
	arr := strings.Split(n.Content, " ")
	if len(arr) != 0 {
		if arr[0] == "" {
			return "untitled"
		} else {
			return arr[0]
		}
	} else {
		return "untitled"
	}
}
