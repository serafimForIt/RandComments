package main

import (
	"config"
	"front"
)

func main() {
	Starter()
}

var Settings config.Settings

func Starter() {

	Settings.Load()

	front.AppMain(
		Settings.Emoticons_num,
		Settings.Standard_phrases_num,
		Settings.Exclusive_phrases_num,
		Settings.Max_emoticon_multiplier,
		Settings.Min_quantity_from_set,
		Settings.Max_quantity_from_set,
		Settings.Delimiter,
		Settings.Split_in_line,
	)
}
