package youbeacomm

import (
	"encoding/json"
	"net/http"
)

type Phrases []Phrase

type Phrase struct {
	Id      string       `json:"id"`
	Strings []L10nString `json:"strings"`
}

type L10nString struct {
	Lang   string `json:"lang"`
	String string `json:"string"`
}

func PhrasePhraseIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PhrasePhraseIdResponsesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PhraseSuggestionsFrequentGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var phrases = []Phrase{
		Phrase{
			Id: "16286557-8ace-4023-8407-a7cc5a25f930",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "この電車は湯の川に止まりますか？",
				},
				L10nString{
					Lang:   "en",
					String: "Does this train stop at Yunokawa?",
				},
			},
		},
	}

	json.NewEncoder(w).Encode(phrases)

}

func PhraseSuggestionsLocationDeviceIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
