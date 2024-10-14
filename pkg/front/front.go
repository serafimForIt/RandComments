package front

import (
	"controller"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AppMain(
	Emoticons_num,
	Standard_phrases_num,
	Exclusive_phrases_num,
	Max_emoticon_multiplier,
	Min_quantity_from_set,
	Max_quantity_from_set float64,
	Delimiter,
	Split_in_line string,
) {
	app := app.New()
	mainWindow := app.NewWindow("shut the fuck up and do your work")
	mainWindow.Resize(fyne.NewSize(1000, 700))
	mainWindow.CenterOnScreen()

	clipBoard := mainWindow.Clipboard()
	clipBoardString := ""

	settingsLabel := widget.NewRichTextWithText("Settings: ")
	x_label := widget.NewLabel(string("                 "))

	emoticons_x_text, emoticons_x_input := SettingsEntryCreator(
		"Emoticon multiplier ", "num", StrConv(Emoticons_num),
	)

	standard_phrase_x_text, standard_phrase_x_input := SettingsEntryCreator(
		"Standard phrase multiplier ", "num", StrConv(Standard_phrases_num),
	)

	exclusive_phrase_x_text, exclusive_phrase_x_input := SettingsEntryCreator(
		"Exclusive phrase multiplier ", "num", StrConv(Exclusive_phrases_num),
	)

	emoticons_in_a_row_text, emoticons_in_a_row_input := SettingsEntryCreator(
		"Max emotcions in a row multiplier ", "num", StrConv(Max_emoticon_multiplier),
	)

	min_quantity_from_set_text, min_quantity_from_set_input := SettingsEntryCreator(
		"Min quantity from the set ", "num", StrConv(Min_quantity_from_set),
	)

	max_quantity_from_set_text, max_quantity_from_set_input := SettingsEntryCreator(
		"Max quantity from the set ", "num", StrConv(Max_quantity_from_set),
	)

	emoticons_delimiter_text, emoticons_delimiter_input := SettingsEntryCreator(
		"Statement delimiter ", "delimiter", Delimiter,
	)

	emoticons_split_text, emoticons_split_input := SettingsEntryCreator(
		"Emoticon delimiter ", "splitLineEntry", Split_in_line,
	)

	right_text := widget.NewMultiLineEntry()
	reroll_result := controller.Reroll(
		Emoticons_num,
		Standard_phrases_num,
		Exclusive_phrases_num,
		Max_emoticon_multiplier,
		Min_quantity_from_set,
		Max_quantity_from_set,
		Delimiter,
		Split_in_line,
		"emoticons.txt",
		"standard_phrases.txt",
		"exclusive_phrases.txt",
		"result.txt",
	)
	for i, s := range reroll_result {
		if i == 0 && s != "" {
			right_text.SetText(s)
		} else if s != "" {
			right_text.Append(s)
		}
		controller.WriteTXT("for_writing.txt", reroll_result)
	}

	btn_save_settings := widget.NewButton("Save settings", func() {
		emoticons_x_input.Entry.Refresh()
		standard_phrase_x_input.Entry.Refresh()
		exclusive_phrase_x_input.Entry.Refresh()
		emoticons_delimiter_input.Entry.Refresh()
		min_quantity_from_set_input.Entry.Refresh()
		max_quantity_from_set_input.Entry.Refresh()
		emoticons_delimiter_input.Entry.Refresh()
		emoticons_split_input.Entry.Refresh()

		// validator
		if emoticons_x_input.Entry.Text != emoticons_x_input.Entry.PlaceHolder && emoticons_x_input.Entry.Text != "" {
			Emoticons_num, _ = strconv.ParseFloat(emoticons_x_input.Entry.Text, 64)
		}
		if standard_phrase_x_input.Entry.Text != standard_phrase_x_input.Entry.PlaceHolder && standard_phrase_x_input.Entry.Text != "" {
			Standard_phrases_num, _ = strconv.ParseFloat(standard_phrase_x_input.Entry.Text, 64)
		}
		if exclusive_phrase_x_input.Entry.Text != exclusive_phrase_x_input.Entry.PlaceHolder && exclusive_phrase_x_input.Entry.Text != "" {
			Exclusive_phrases_num, _ = strconv.ParseFloat(exclusive_phrase_x_input.Entry.Text, 64)
		}
		if emoticons_in_a_row_input.Entry.Text != emoticons_in_a_row_input.Entry.PlaceHolder && emoticons_in_a_row_input.Entry.Text != "" {
			Max_emoticon_multiplier, _ = strconv.ParseFloat(emoticons_in_a_row_input.Entry.Text, 64)
		}
		if min_quantity_from_set_input.Entry.Text != min_quantity_from_set_input.Entry.PlaceHolder && min_quantity_from_set_input.Entry.Text != "" {
			Min_quantity_from_set, _ = strconv.ParseFloat(min_quantity_from_set_input.Entry.Text, 64)
		}
		if max_quantity_from_set_input.Entry.Text != max_quantity_from_set_input.Entry.PlaceHolder && max_quantity_from_set_input.Entry.Text != "" {
			Max_quantity_from_set, _ = strconv.ParseFloat(max_quantity_from_set_input.Entry.Text, 64)
		}
		if emoticons_delimiter_input.Entry.Text != emoticons_delimiter_input.Entry.PlaceHolder && emoticons_delimiter_input.Entry.Text != "" {
			Delimiter = emoticons_delimiter_input.Entry.Text
		}
		if emoticons_split_input.Entry.Text != emoticons_split_input.Entry.PlaceHolder {
			Split_in_line = emoticons_split_input.Entry.Text
		}
		// end of validator

		controller.WriteSettingsChanges(
			Emoticons_num,
			Standard_phrases_num,
			Exclusive_phrases_num,
			Max_emoticon_multiplier,
			Min_quantity_from_set,
			Max_quantity_from_set,
			Delimiter,
			Split_in_line,
		)
		emoticons_x_input.Entry.SetPlaceHolder(StrConv(Emoticons_num))
		emoticons_x_input.Entry.SetText("")
		standard_phrase_x_input.Entry.SetPlaceHolder(StrConv(Standard_phrases_num))
		standard_phrase_x_input.Entry.SetText("")
		exclusive_phrase_x_input.Entry.SetPlaceHolder(StrConv(Exclusive_phrases_num))
		exclusive_phrase_x_input.Entry.SetText("")
		emoticons_in_a_row_input.Entry.SetPlaceHolder(StrConv(Max_emoticon_multiplier))
		emoticons_in_a_row_input.Entry.SetText("")
		min_quantity_from_set_input.Entry.SetPlaceHolder(StrConv(Min_quantity_from_set))
		min_quantity_from_set_input.Entry.SetText("")
		max_quantity_from_set_input.Entry.SetPlaceHolder(StrConv(Max_quantity_from_set))
		max_quantity_from_set_input.Entry.SetText("")
		emoticons_delimiter_input.Entry.SetPlaceHolder(Delimiter)
		emoticons_delimiter_input.Entry.SetText("")
		emoticons_split_input.Entry.SetPlaceHolder(Split_in_line)
	})

	btn_reroll := widget.NewButton(" Reroll ", func() {
		clipBoardString = ""
		reroll_result := controller.Reroll(
			Emoticons_num,
			Standard_phrases_num,
			Exclusive_phrases_num,
			Max_emoticon_multiplier,
			Min_quantity_from_set,
			Max_quantity_from_set,
			Delimiter,
			Split_in_line,
			"emoticons.txt",
			"standard_phrases.txt",
			"exclusive_phrases.txt",
			"result.txt",
		)
		for i, s := range reroll_result {
			if i == 0 && s != "" {
				right_text.SetText(s)
				clipBoardString += s
			} else if s != "" {
				right_text.Append(s)
				clipBoardString += s
			}
		}
		controller.WriteTXT("for_writing.txt", reroll_result)
	})

	btn_copy := widget.NewButton(" Copy ", func() {
		clipBoard.Content()
		clipBoard.SetContent(right_text.Text)
	})

	v1 := container.NewVBox(
		settingsLabel,
		emoticons_x_text,
		standard_phrase_x_text,
		exclusive_phrase_x_text,
		emoticons_in_a_row_text,
		min_quantity_from_set_text,
		max_quantity_from_set_text,
		emoticons_delimiter_text,
		emoticons_split_text,
		btn_save_settings,
		x_label,
	)

	v2 := container.NewVBox(
		x_label,
		&emoticons_x_input.Entry,
		&standard_phrase_x_input.Entry,
		&exclusive_phrase_x_input.Entry,
		&emoticons_in_a_row_input.Entry,
		&min_quantity_from_set_input.Entry,
		&max_quantity_from_set_input.Entry,
		&emoticons_delimiter_input.Entry,
		&emoticons_split_input.Entry,
		x_label,
		x_label,
	)

	v3 := container.NewHBox(btn_reroll)

	v4 := container.NewHBox(btn_copy)

	l_settings := container.NewVBox(container.NewHBox(v1, v2))

	lb := container.NewVBox(container.NewHBox(v3, v4))

	rScroll := container.NewStack(right_text)
	leftPart := container.NewBorder(nil, nil, nil, nil, l_settings)
	rightPart := container.NewBorder(nil, lb, nil, nil, rScroll)

	e_b_btn_S, e_b_MEntry := FileBorderCreator("emoticons.txt")
	s_b_btn_S, s_b_MEntry := FileBorderCreator("standard_phrases.txt")
	ex_b_btn_S, ex_b_MEntry := FileBorderCreator("exclusive_phrases.txt")

	// faq := widget.NewMultiLineEntry()
	// faq.SetText(FAQ_text)
	// faq.Disabled()
	// faq.
	//, fyne.TextAlign(fyne.TextWrapWord), fyne.TextStyle{})
	//faq.Resize(fyne.MeasureText(FAQ_text, 60, widget.Adaptive))
	faq := container.NewVScroll(widget.NewLabel(FAQ_text))

	tabs := container.NewAppTabs(
		container.NewTabItem("Main", container.NewBorder(nil, nil, leftPart, nil, rightPart)),
		container.NewTabItem("Emoticons", container.NewBorder(nil, e_b_MEntry, nil, nil, e_b_btn_S)),
		container.NewTabItem("Standart phrases", container.NewBorder(nil, s_b_MEntry, nil, nil, s_b_btn_S)),
		container.NewTabItem("Exclusive phrases", container.NewBorder(nil, ex_b_MEntry, nil, nil, ex_b_btn_S)),
		container.NewTabItem("FAQ", container.NewBorder(nil, nil, nil, nil, faq)),
	)

	mainWindow.SetContent(tabs)
	mainWindow.ShowAndRun()
}
