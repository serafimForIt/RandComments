package txt_handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(path_to_file string) []string {
	file, err := os.Open(path_to_file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	our_strings := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text() + "\n"
		our_strings = append(our_strings, i)
	}
	return our_strings
}

func WriteFile(path_to_file string, in_strings []string) {
	file, err := os.Create(path_to_file)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for i, s := range in_strings {
		if i == 0 {
			file.WriteString(s)
		} else if i == len(in_strings)-1 {
			file.WriteString("\n" + strings.TrimSuffix(s, "\n"))
		} else {
			file.WriteString("\n" + s)
		}
	}
}

func WriteRerollFile(path_to_file string, in_strings []string) {
	file, err := os.Create(path_to_file)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, s := range in_strings {
		// if i == 0 && s != "" {
		// 	file.WriteString(s)
		// } else if s != "" {
		// 	file.WriteString(s + "\n")
		// }
		if s != "" {
			file.WriteString(s)
		}
	}
}

func ArrayToString(our_strings []string) string {
	str := ""
	for _, s := range our_strings {
		str += s
	}
	return str
}

func StringToArray(our_string string) []string {
	strings := strings.Split(our_string, "\n")
	final_array := []string{}
	final_array = append(final_array, strings...)
	return final_array
}

func DeleteLastEnter(our_strings []string) []string {
	our_strings[len(our_strings)-1] = strings.TrimSuffix(our_strings[len(our_strings)-1], "\n")
	return our_strings
}
