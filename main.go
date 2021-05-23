package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sap200/notes/notes"
	"github.com/sap200/notes/utils"
)

// Global variable to store notes
var noteAll = []*notes.Note{}

// indicates current button pointed in the UI
var currentButton *widget.Button

// indicates current note index pointed by UI
var currentNote int

// current contentString in the input
var contentString = ""

// The persistence file
var backupFile = "notes.txt"

// Initialize the state of the program
func init() {
	// Initialize current button to nil
	currentButton = nil
	// no note is pointed when button is nil
	currentNote = -1

	// Load the current state of array
	utils.LoadState(backupFile, &noteAll)

}

func main() {
	a := app.New()

	w := a.NewWindow("Notes")
	w.SetContent(loadUI())
	w.Resize(fyne.NewSize(900, 700))

	w.ShowAndRun()
}

// -------------------------------------------------------------- AUXILARY FUNCTIONS -----------------------------------------------------------------------------------------------------------

// Load the UI component
func loadUI() fyne.CanvasObject {

	// contains the Notes references on lhs
	list := container.NewVBox()
	// The save button helps in saving
	save := widget.NewButton("save", func() {
		if currentNote != -1 {
			// get the content and save it
			// get the notes
			noteAll[currentNote].Content = contentString
			contentString = ""
			currentButton.SetText(noteAll[currentNote].GetTitle())
		}

		// Plus every time save is pressed get the data and save it to a file named notes.txt
		utils.SaveState(backupFile, &noteAll)
	})
	// bind the val to the text input multiline
	val := binding.NewString()
	callback := binding.NewDataListener(func() {
		var err error
		contentString, err = val.Get()
		if err != nil {
			log.Fatalln(err)
		}
	})
	// Register the listener
	val.AddListener(callback)

	// The multiline input box
	content := widget.NewMultiLineEntry()
	if currentNote == -1 || currentButton == nil {
		content.Disable()
	}
	content.Bind(val)
	rightPane := container.NewVSplit(content, save)
	rightPane.Offset = 0.95

	// upper toolbar to add or remove notes
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addNote(list, content)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			removeNote(list, content)

		}),
	)

	sidebar := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)

	split := container.NewHSplit(sidebar, rightPane)
	split.Offset = 0.25

	if len(noteAll) != 0 {
		renderNotes(list, content)
	}

	return split
}

// --------------------------------------------------------------------------------------------------------- UI AUXILARY FUNCTIONS -----------------------------------------------------------------------------------------------

// Add a note
func addNote(list *fyne.Container, content *widget.Entry) {
	note := notes.NewNote("", len(noteAll))
	noteAll = append(noteAll, note)
	var noteButton *widget.Button
	noteButton = widget.NewButton(note.GetTitle(), func() {
		content.Enable()
		currentButton = noteButton
		currentNote = note.Index
		content.SetText(note.Content)
		fmt.Println(noteAll)
	})

	list.Add(noteButton)
}

// Remove a note
func removeNote(list *fyne.Container, content *widget.Entry) {
	if currentNote != -1 && currentButton != nil {
		list.Remove(currentButton)
		if len(noteAll) != 1 {
			noteAll = append(noteAll[:currentNote], noteAll[currentNote+1:]...)
		} else {
			noteAll = []*notes.Note{}
		}
		currentNote = -1
		currentButton = nil
		content.SetText("")
		content.Disable()
		fmt.Println(noteAll)
	}
}

// Render notes if it is initially present
func renderNotes(list *fyne.Container, content *widget.Entry) {
	for _, n := range noteAll {
		note := n

		var noteButton *widget.Button
		noteButton = widget.NewButton(note.GetTitle(), func() {
			content.Enable()
			currentButton = noteButton
			currentNote = note.Index
			content.SetText(note.Content)
			fmt.Println(noteAll)
		})

		list.Add(noteButton)
	}

}
