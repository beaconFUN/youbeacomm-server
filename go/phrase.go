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
		Phrase{
			Id: "520bf1d4-21ca-4014-a274-37cfaadf0c8b",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "写真を撮ってもらえますか？",
				},
				L10nString{
					Lang:   "en",
					String: "Could you take a picture?",
				},
			},
		},
		Phrase{
			Id: "b10cb08e-4d8f-49f6-8b75-8b86c9d98ae9",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "一番近いトイレはどこですか？",
				},
				L10nString{
					Lang:   "en",
					String: "Where is the nearest toilet?",
				},
			},
		},
		Phrase{
			Id: "275f3d49-d32a-4e7f-90a8-ac7be054dc61",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "ここに無料のWi-Fiはありますか？",
				},
				L10nString{
					Lang:   "en",
					String: "Can I use free Wi-Fi here?",
				},
			},
		},
		Phrase{
			Id: "dec26bae-fdf5-49fe-80e6-dde45e12a3ac",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "おすすめのレストランを教えてください。",
				},
				L10nString{
					Lang:   "en",
					String: "Could you tell me the recommended restaurant?",
				},
			},
		},
		Phrase{
			Id: "91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "函館は何が有名ですか？",
				},
				L10nString{
					Lang:   "en",
					String: "What is Hakodate famous for?",
				},
			},
		},
		Phrase{
			Id: "8d3b865a-fe6b-43ba-a20d-484a7fd70e5c",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "おすすめの日本の食べ物はありますか？",
				},
				L10nString{
					Lang:   "en",
					String: "What Japanese food should I try?",
				},
			},
		},
		Phrase{
			Id: "66b4b880-5a28-4cc7-b935-c9b8e5be7a3f",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "函館駅に行きたいのですが。",
				},
				L10nString{
					Lang:   "en",
					String: "I'd like to go to Hakodate station.",
				},
			},
		},
		Phrase{
			Id: "354e538a-7259-47f2-a7d3-dc72a0182d21",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "次のバスはいつですか？",
				},
				L10nString{
					Lang:   "en",
					String: "When is the next bus?",
				},
			},
		},
		Phrase{
			Id: "67965004-ae7f-446f-8dc4-082e6909aa32",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "一番近いコンビニはどこですか？",
				},
				L10nString{
					Lang:   "en",
					String: "Where is the nearest convenience store?",
				},
			},
		},
		Phrase{
			Id: "e6d7f905-55e7-437e-a048-82f500364dd4",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "どこに行けば日本円に両替ができますか？",
				},
				L10nString{
					Lang:   "en",
					String: "Where can I get a currency exchange into Japanese yen?",
				},
			},
		},
		Phrase{
			Id: "c5f22c13-f08e-4c92-9bed-5a39b2564b1a",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "観光案内所はどこですか？",
				},
				L10nString{
					Lang:   "en",
					String: "Where is the tourist information center?",
				},
			},
		},
		Phrase{
			Id: "12a85516-76a2-4aa2-b96e-197f511b4e9a",
			Strings: []L10nString{
				L10nString{
					Lang:   "ja",
					String: "電車に乗るにはどうすればいいですか？",
				},
				L10nString{
					Lang:   "en",
					String: "Please tell me how to get on a train.",
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
