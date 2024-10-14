package back

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var alpha string = `abcdefghijklmnopqrstuvwxyz1234567890абвгдеёжзийклмнопрстуфхцчшщъыьэюя~!@#№$;%^:&?*()-_=+|?/.>,<'"`
var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RunReroll(
	emoticons_num, standard_phrases_num, exclusive_phrases_num, max_emoticon_multiplier, min_quantity_from_set, max_quantity_from_set float64,
	delimiter, split_in_line string,
	emoticons_strs, standart_phrases_strs, exclusive_phrases_strs []string,
) []string {
	emoticons := RandShuffle(
		Emoticon_multiplier(
			Clipper(
				RandShuffle(
					Delimiter_handler(emoticons_strs, delimiter, min_quantity_from_set, max_quantity_from_set),
				), emoticons_num),
			max_emoticon_multiplier, split_in_line),
	)
	standard_phrases := RandShuffle(
		Clipper(
			Delimiter_handler(
				standart_phrases_strs, delimiter, min_quantity_from_set, max_quantity_from_set),
			standard_phrases_num),
	)
	exclusive_phrases := RandShuffle(
		Clipper(
			Delimiter_handler(exclusive_phrases_strs, delimiter, min_quantity_from_set, max_quantity_from_set),
			exclusive_phrases_num),
	)
	finalStrings := RandShuffle(AppendArrays(emoticons, standard_phrases, exclusive_phrases))
	return finalStrings
}

func Creator_of_rand_set(Min_quantity_from_set, Max_quantity_from_set float64, technical_variable []string) []string {
	min, max := int(Min_quantity_from_set), int(Max_quantity_from_set)
	result := []string{}
	our_rand_int := r.Intn(max-min+1) + min
	for i := 0; i < our_rand_int; i++ {
		j := technical_variable[r.Intn(len(technical_variable))]
		result = append(result, j)
	}
	return result
}

func Delimiter_handler(in_strings []string, delimiter string, Min_quantity_from_set, Max_quantity_from_set float64) []string {
	our_strings, technical_variable := []string{}, []string{}
	len_in_strings := len(in_strings)
	for i, j := range in_strings {
		if strings.ContainsAny(strings.ToLower(j), alpha) {
			if !strings.Contains(j, delimiter) {
				if i < len_in_strings-1 {
					if !strings.Contains(in_strings[i+1], delimiter) {
						our_strings = append(our_strings, j)
					} else if strings.Contains(in_strings[i+1], delimiter) {
						technical_variable = append(technical_variable, j)
					}
				} else {
					our_strings = append(our_strings, j)
				}
			} else if strings.Contains(j, delimiter) {
				if i < len_in_strings-1 {
					if !strings.Contains(in_strings[i+1], delimiter) {
						j, _ := strings.CutPrefix(j, delimiter)
						technical_variable = append(technical_variable, j)
						if int(Min_quantity_from_set) < len(technical_variable) {
							fin_tech_var := Creator_of_rand_set(Min_quantity_from_set, Max_quantity_from_set, technical_variable)
							our_strings = AppendArrays(our_strings, fin_tech_var)
						} else {
							our_strings = append(our_strings, technical_variable[r.Intn(len(technical_variable))])
						}
						technical_variable = []string{}
					} else if strings.Contains(in_strings[i+1], delimiter) {
						j, _ := strings.CutPrefix(j, delimiter)
						technical_variable = append(technical_variable, j)
					}
				} else {
					j, _ := strings.CutPrefix(j, delimiter)
					technical_variable = append(technical_variable, j)
					if int(Min_quantity_from_set) < len(technical_variable) {
						fin_tech_var := Creator_of_rand_set(Min_quantity_from_set, Max_quantity_from_set, technical_variable)
						our_strings = AppendArrays(our_strings, fin_tech_var)
					} else {
						our_strings = append(our_strings, technical_variable[r.Intn(len(technical_variable))])
					}
					technical_variable = []string{}
				}
			}
		}
	}
	return our_strings
}

func RandShuffle(in_strings []string) []string {
	our_strings := in_strings[:]
	for i := 0; i < len(our_strings); i++ {
		r := r.Intn(i + 1)
		our_strings[i], our_strings[r] = our_strings[r], our_strings[i]
	}
	return our_strings
}

func Emoticon_multiplier(in_strings []string, max_num float64, split_in_line string) []string {
	our_strings, our_ints := in_strings[:], []int{}
	for ids, i := range our_strings {
		for i := 1; i < int(max_num)+1; i++ {
			our_ints = append(our_ints, r.Intn(int(max_num)))
		}
		our_num := our_ints[r.Intn(len(our_ints))]
		if our_num >= 1 {
			f := strings.TrimPrefix(i, "\n")
			for j := 0; j < our_num; j++ {
				i = strings.Trim(i, "\n")
				i = i + split_in_line + f
			}
		}
		our_strings[ids] = i
	}
	return our_strings
}

func Clipper(in_strings []string, num float64) []string {
	our_strings := []string{}
	in_string_copy := in_strings[:]
	for i := 0.0; i <= num-1.0; i++ {
		our_strings = append(our_strings, in_string_copy...)
	}

	in_strings_len := len(in_string_copy)
	fractional_part := num - math.Floor(num)
	our_f_num := int(float64(in_strings_len) * fractional_part)
	for i := 0; i < our_f_num; i++ {
		our_strings = append(our_strings, in_string_copy[i])
	}
	return our_strings
}

func AppendArrays(in_strings ...[]string) []string {
	our_strings := []string{}
	for _, i := range in_strings {
		our_strings = append(our_strings, i...)
	}
	return our_strings
}
