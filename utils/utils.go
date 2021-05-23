// package utils provide functions for persistence
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sap200/notes/notes"
)

// Saves the current noteAll array in the file storage
// This solves persistence
func SaveState(backupFile string, noteAll *[]*notes.Note) {
	dat, err := json.Marshal(noteAll)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(backupFile, dat, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// Loads the data from initial file
func LoadState(backupFile string, noteAll *[]*notes.Note) {
	if _, err := os.Stat(backupFile); os.IsNotExist(err) {
		fmt.Println("No backup file")
		return
	}
	content, err := ioutil.ReadFile(backupFile)
	if err != nil {
		fmt.Println(err)
	}
	var initNote []*notes.Note
	err = json.Unmarshal(content, &initNote)
	if err != nil {
		fmt.Println(err)
	}
	*noteAll = initNote
	fmt.Println(initNote)
}
