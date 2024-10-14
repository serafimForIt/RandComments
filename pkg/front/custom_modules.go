package front

import (
	"controller"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var FAQ_text string = controller.FAQ_text("FAQ.txt")

type newEntry struct {
	Entry widget.Entry
	Val   string
}

func NewEntry(s string) *newEntry {
	entry := &newEntry{}
	entry.Entry.ExtendBaseWidget(&entry.Entry)
	entry.Val = s
	return entry
}

func (e *newEntry) TypedRune(r rune) {
	if e.Val == "num" {
		if (r >= '0' && r <= '9') || r == '.' || r == ',' {
			e.Entry.TypedRune(r)
		}
	} else if e.Val == "delimiter" {
		if r == '~' || r == '!' || r == '*' || r == '&' || r == '%' || r == '/' {
			e.Entry.TypedRune(r)
		}
	} else if e.Val == "splitLineEntry" {
		if r == ' ' || r == 0 {
			e.Entry.TypedRune(r)
		}
	}
}

func (e *newEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func FileBorderCreator(path_to_file string) (*widget.Entry, *widget.Button) {
	b_MEntry := widget.NewMultiLineEntry()
	b_text := controller.ReadTXT(path_to_file)
	b_MEntry.SetText(controller.ArrayToString(b_text))

	b_btn_S := widget.NewButton(" Save ", func() {
		controller.WriteTXT(path_to_file, controller.StringToArray(b_MEntry.Text))
	})

	return b_MEntry, b_btn_S
}

func StrConv(str float64) string {
	return strconv.FormatFloat(str, 'g', 10, 64)
}

func SettingsEntryCreator(eText, eType, e string) (*widget.Label, *newEntry) {
	label := widget.NewLabel(eText)
	input := NewEntry(eType)
	input.Entry.SetPlaceHolder(e)
	return label, input
}
