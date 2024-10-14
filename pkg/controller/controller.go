package controller

import (
	"back"
	"config"
	"txt_handler"
)

var Settings config.Settings

func WriteSettingsChanges(
	emoticons_num, standard_phrases_num, exclusive_phrases_num, max_emoticon_multiplier, min_quantity_from_set, max_quantity_from_set float64,
	delimiter, split_in_line string,
) {
	Settings.Emoticons_num = emoticons_num
	Settings.Standard_phrases_num = standard_phrases_num
	Settings.Exclusive_phrases_num = exclusive_phrases_num
	Settings.Max_emoticon_multiplier = max_emoticon_multiplier
	Settings.Min_quantity_from_set = min_quantity_from_set
	Settings.Max_quantity_from_set = max_quantity_from_set
	Settings.Delimiter = delimiter
	Settings.Split_in_line = split_in_line
	Settings.Write()
}

func Reroll(
	emoticons_num, standard_phrases_num, exclusive_phrases_num, max_emoticon_multiplier, min_quantity_from_set, max_quantity_from_set float64,
	delimiter, split_in_line, emoticons_file, standart_phrases_file, exclusive_phrases_file, result_file string,
) []string {
	emoticons_strs := txt_handler.ReadFile(emoticons_file)
	standart_phrases_strs := txt_handler.ReadFile(standart_phrases_file)
	exclusive_phrases_strs := txt_handler.ReadFile(exclusive_phrases_file)
	result := back.RunReroll(
		emoticons_num,
		standard_phrases_num,
		exclusive_phrases_num,
		max_emoticon_multiplier,
		min_quantity_from_set,
		max_quantity_from_set,
		delimiter,
		split_in_line,
		emoticons_strs,
		standart_phrases_strs,
		exclusive_phrases_strs,
	)
	txt_handler.DeleteLastEnter(result)
	txt_handler.WriteRerollFile(result_file, result)
	return result
}

func WriteTXT(path_to_file string, in_strings []string) {
	txt_handler.WriteFile(path_to_file, in_strings)
}

func ReadTXT(path_to_file string) []string {
	return txt_handler.ReadFile(path_to_file)
}

func ArrayToString(our_strings []string) string {
	return txt_handler.ArrayToString(our_strings)
}

func StringToArray(str string) []string {
	return txt_handler.StringToArray(str)
}

func FAQ_text(path_to_file string) string {
	return ArrayToString(ReadTXT(path_to_file))
}
