package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Emoticons_num           float64 `json:"emoticons_num"`
	Standard_phrases_num    float64 `json:"standard_phrases_num"`
	Exclusive_phrases_num   float64 `json:"exclusive_phrases_num"`
	Max_emoticon_multiplier float64 `json:"max_emoticon_multiplier"`
	Min_quantity_from_set   float64 `json:"min_quantity_from_set"`
	Max_quantity_from_set   float64 `json:"max_quantity_from_set"`
	Delimiter               string  `json:"delimiter"`
	Split_in_line           string  `json:"split_in_line"`
}

func (s *Settings) Load() {
	data, err := os.ReadFile("settings.json")

	if err != nil {
		fmt.Println("no settings loaded...")
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
}

func (s *Settings) Write() {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	_ = os.WriteFile("settings.json", data, 0644)
}
