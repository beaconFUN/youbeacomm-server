package youbeacomm

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

//type Phrases []Phrase

type Phrase struct {
	Id      string       `json:"id"`
	Strings []L10nString `json:"strings"`
}

type L10nString struct {
	Lang   string `json:"lang"`
	String string `json:"string"`
}

type PhraseResponse struct {
	Id        string   `json:"id"`
	Responses []string `json:"responses"`
}

var phrases = []Phrase{
	Phrase{
		Id: "16286557-8ace-4023-8407-a7cc5a25f930",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "おすすめを教えてください。",
			},
			L10nString{
				Lang:   "en",
				String: "What do you recommend?",
			},
		},
	},
	Phrase{
		Id: "520bf1d4-21ca-4014-a274-37cfaadf0c8b",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "海外に送れますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Can you ship overseas?",
			},
		},
	},
	Phrase{
		Id: "b10cb08e-4d8f-49f6-8b75-8b86c9d98ae9",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "賞味期限はいつまでですか？",
			},
			L10nString{
				Lang:   "en",
				String: "When is the expiration date?",
			},
		},
	},
	Phrase{
		Id: "275f3d49-d32a-4e7f-90a8-ac7be054dc61",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "たくさん買ったら、値引きはありますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Is there a volume discount?",
			},
		},
	},
	Phrase{
		Id: "dec26bae-fdf5-49fe-80e6-dde45e12a3ac",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "試食はできますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Can I sample it?",
			},
		},
	},
	Phrase{
		Id: "91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "クレジットカードは使えますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Can I use a credit card?",
			},
		},
	},
	Phrase{
		Id: "8d3b865a-fe6b-43ba-a20d-484a7fd70e5c",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "袋を多めにもらえますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Can I have more bags?",
			},
		},
	},
	Phrase{
		Id: "66b4b880-5a28-4cc7-b935-c9b8e5be7a3f",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "ここでしか買えないものはありますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Is there anything I can only buy here?",
			},
		},
	},
	Phrase{
		Id: "354e538a-7259-47f2-a7d3-dc72a0182d21",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "プレゼント用に包装してください。",
			},
			L10nString{
				Lang:   "en",
				String: "Can you gift-wrap this, please?",
			},
		},
	},
	Phrase{
		Id: "0cf79b03-d8fd-41e2-88cb-f9ff6af5a0c1",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "はい。",
			},
			L10nString{
				Lang:   "en",
				String: "Yes.",
			},
		},
	},
	Phrase{
		Id: "c965aa62-9019-4cd6-a362-fa94e32ca734",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "いいえ。",
			},
			L10nString{
				Lang:   "en",
				String: "No.",
			},
		},
	},
	Phrase{
		Id: "582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "いいですよ。",
			},
			L10nString{
				Lang:   "en",
				String: "Okay.",
			},
		},
	},
	Phrase{
		Id: "d55924d6-13d8-4926-9801-cad754d5c7ff",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "ごめんなさい。",
			},
			L10nString{
				Lang:   "en",
				String: "Sorry.",
			},
		},
	},
	Phrase{
		Id: "66504b47-db33-4f1a-9629-ec64c1a2d7d4",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "今忙しいので。",
			},
			L10nString{
				Lang:   "en",
				String: "Sorry, I'm busy.",
			},
		},
	},
}

var responseIds = []PhraseResponse{
	PhraseResponse{
		Id:        "16286557-8ace-4023-8407-a7cc5a25f930",
		Responses: []string{},
	},
	PhraseResponse{
		Id: "520bf1d4-21ca-4014-a274-37cfaadf0c8b",
		Responses: []string{
			"0cf79b03-d8fd-41e2-88cb-f9ff6af5a0c1",
			"c965aa62-9019-4cd6-a362-fa94e32ca734",
		},
	},
	PhraseResponse{
		Id:        "b10cb08e-4d8f-49f6-8b75-8b86c9d98ae9",
		Responses: []string{},
	},
	PhraseResponse{
		Id: "275f3d49-d32a-4e7f-90a8-ac7be054dc61",
		Responses: []string{
			"0cf79b03-d8fd-41e2-88cb-f9ff6af5a0c1",
			"c965aa62-9019-4cd6-a362-fa94e32ca734",
		},
	},
	PhraseResponse{
		Id: "dec26bae-fdf5-49fe-80e6-dde45e12a3ac",
		Responses: []string{
			"0cf79b03-d8fd-41e2-88cb-f9ff6af5a0c1",
			"c965aa62-9019-4cd6-a362-fa94e32ca734",
		},
	},
	PhraseResponse{
		Id: "91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
		Responses: []string{
			"0cf79b03-d8fd-41e2-88cb-f9ff6af5a0c1",
			"c965aa62-9019-4cd6-a362-fa94e32ca734",
		},
	},
	PhraseResponse{
		Id: "8d3b865a-fe6b-43ba-a20d-484a7fd70e5c",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
			"66504b47-db33-4f1a-9629-ec64c1a2d7d4",
		},
	},
	PhraseResponse{
		Id:        "66b4b880-5a28-4cc7-b935-c9b8e5be7a3f",
		Responses: []string{},
	},
	PhraseResponse{
		Id: "354e538a-7259-47f2-a7d3-dc72a0182d21",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
		},
	},
	PhraseResponse{
		Id: "67965004-ae7f-446f-8dc4-082e6909aa32",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
			"66504b47-db33-4f1a-9629-ec64c1a2d7d4",
		},
	},
}

var frequentIds = []string{
	"16286557-8ace-4023-8407-a7cc5a25f930",
	"520bf1d4-21ca-4014-a274-37cfaadf0c8b",
	"275f3d49-d32a-4e7f-90a8-ac7be054dc61",
	"dec26bae-fdf5-49fe-80e6-dde45e12a3ac",
	"91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
	"354e538a-7259-47f2-a7d3-dc72a0182d21",
}

func filterPhraseById(phrase []Phrase, id string) (*Phrase, error) {
	for _, v := range phrase {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, errors.New("No phrase found.")
}

func filterPhrasesById(phrases []Phrase, ids []string) (filtered []Phrase) {
	for _, id := range ids {
		if p, err := filterPhraseById(phrases, id); err == nil {
			filtered = append(filtered, *p)
		}
	}

	return
}

func PhrasePhraseIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)

	phrase, err := filterPhraseById(phrases, vars["phraseId"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phrase)
}

func PhrasePhraseIdResponsesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)

	phrase, err := filterPhraseById(phrases, vars["phraseId"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&[]Phrase{})
		return
	}

	// Should be moved to SQL
	var resps []string
	for _, v := range responseIds {
		if v.Id == phrase.Id {
			resps = append(resps, v.Responses...)
		}
	}

	if len(resps) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&[]Phrase{})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filterPhrasesById(phrases, resps))
}

func PhraseSuggestionsFrequentGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(filterPhrasesById(phrases, frequentIds))

}

func PhraseSuggestionsLocationDeviceIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
