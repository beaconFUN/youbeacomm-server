package youbeacomm

import (
	"encoding/json"
	"errors"
	"fmt"
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
				String: "朝市でのおすすめのスポットは何ですか？",
			},
			L10nString{
				Lang:   "en",
				String: "What do you recommend at Morning Market?",
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
		Id: "477948a4-a6bd-4214-a6a3-fad3a25cc597",
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
		Id: "28c7acab-d366-493a-8ac9-b1380570b4e8",
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
		Id: "8216d5fb-4f3c-4fb7-8bd7-0559b82215d6",
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
		Id: "4603cb11-d2df-4665-ba3d-a482a50824d4",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "新函館北斗駅に行くにはどうすればいいですか？",
			},
			L10nString{
				Lang:   "en",
				String: "How can I get to Shin-Hakodate-Hokuto Station?",
			},
		},
	},
	Phrase{
		Id: "46c70b37-516e-4da2-8ff6-e377006b0fe4",
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
	Phrase{
		Id: "ffc1358c-e1d0-4a4d-ba6a-ad06310671eb",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "駅二市場のイカ釣りです。",
			},
			L10nString{
				Lang:   "en",
				String: "Squid fishing at Ekini Market is recommended.",
			},
		},
	},
	Phrase{
		Id: "70676dad-cd70-443a-be19-4b4e7eacd9d2",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "ウニやイクラの海鮮丼です。",
			},
			L10nString{
				Lang:   "en",
				String: "A bowl of rice with sea urchin and salmon roe on top is recommended.",
			},
		},
	},
	Phrase{
		Id: "5169fe03-7e6e-4ac9-836b-543f3259b162",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "朝ご飯のおすすめはありますか？",
			},
			L10nString{
				Lang:   "en",
				String: "What do you recommend for breakfast?",
			},
		},
	},
	Phrase{
		Id: "05503fdd-c145-463b-86f9-c0825824a147",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "近くにXXXというパン屋さんがありますよ。",
			},
			L10nString{
				Lang:   "en",
				String: "How about visiting the bakery \"XXX\" near here?",
			},
		},
	},
	Phrase{
		Id: "18da1eea-cf80-4710-8cb1-f730bbd2f5ce",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "チェックアウトした後、荷物を預かってもらえますか？",
			},
			L10nString{
				Lang:   "en",
				String: "Could you keep my baggage for a few hours?",
			},
		},
	},
	Phrase{
		Id: "4c177f10-6bad-41e8-9548-27ba7a3471c1",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "[Debug] 検知したビーコン: ホテル",
			},
			L10nString{
				Lang:   "en",
				String: "[Debug] Detected Beacon: Hotel",
			},
		},
	},
	Phrase{
		Id: "699db86d-5f17-4c6f-bb20-b662ede5aa37",
		Strings: []L10nString{
			L10nString{
				Lang:   "ja",
				String: "[Debug] 検知したビーコン: 駅",
			},
			L10nString{
				Lang:   "en",
				String: "[Debug] Detected Beacon: Station",
			},
		},
	},
}

var responseIds = []PhraseResponse{
	PhraseResponse{
		Id: "16286557-8ace-4023-8407-a7cc5a25f930",
		Responses: []string{
			"ffc1358c-e1d0-4a4d-ba6a-ad06310671eb",
			"70676dad-cd70-443a-be19-4b4e7eacd9d2",
		},
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
	PhraseResponse{
		Id: "520bf1d4-21ca-4014-a274-37cfaadf0c8b",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
			"66504b47-db33-4f1a-9629-ec64c1a2d7d4",
		},
	},
	PhraseResponse{
		Id: "b10cb08e-4d8f-49f6-8b75-8b86c9d98ae9",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
			"66504b47-db33-4f1a-9629-ec64c1a2d7d4",
		},
	},
	PhraseResponse{
		Id: "275f3d49-d32a-4e7f-90a8-ac7be054dc61",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
			"66504b47-db33-4f1a-9629-ec64c1a2d7d4",
		},
	},
	PhraseResponse{
		Id: "354e538a-7259-47f2-a7d3-dc72a0182d21",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
			"66504b47-db33-4f1a-9629-ec64c1a2d7d4",
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
	PhraseResponse{
		Id: "5169fe03-7e6e-4ac9-836b-543f3259b162",
		Responses: []string{
			"70676dad-cd70-443a-be19-4b4e7eacd9d2",
			"05503fdd-c145-463b-86f9-c0825824a147",
		},
	},
	PhraseResponse{
		Id: "18da1eea-cf80-4710-8cb1-f730bbd2f5ce",
		Responses: []string{
			"582058db-ac4b-4baf-b9fe-b6bdc23e81fa",
			"d55924d6-13d8-4926-9801-cad754d5c7ff",
		},
	},
}

var frequentIds = []string{
	"477948a4-a6bd-4214-a6a3-fad3a25cc597",
	"28c7acab-d366-493a-8ac9-b1380570b4e8",
	"8216d5fb-4f3c-4fb7-8bd7-0559b82215d6",
	"4603cb11-d2df-4665-ba3d-a482a50824d4",
	"46c70b37-516e-4da2-8ff6-e377006b0fe4",
}

var locationalIds = map[int64][]string{
	int64(0x0000)<<16 + int64(0x0000): { // Debug
		"16286557-8ace-4023-8407-a7cc5a25f930",
		"520bf1d4-21ca-4014-a274-37cfaadf0c8b",
		"275f3d49-d32a-4e7f-90a8-ac7be054dc61",
		"dec26bae-fdf5-49fe-80e6-dde45e12a3ac",
		"91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
		"354e538a-7259-47f2-a7d3-dc72a0182d21",
	},
	int64(0x0000)<<16 + int64(0xFFFF): { // Hotel
		"4c177f10-6bad-41e8-9548-27ba7a3471c1",
		"5169fe03-7e6e-4ac9-836b-543f3259b162",
		"46c70b37-516e-4da2-8ff6-e377006b0fe4",
		"18da1eea-cf80-4710-8cb1-f730bbd2f5ce",
		"8216d5fb-4f3c-4fb7-8bd7-0559b82215d6",
	},
	int64(0xFFFF)<<16 + int64(0xFFFF): { // Station
		"699db86d-5f17-4c6f-bb20-b662ede5aa37",
		"4603cb11-d2df-4665-ba3d-a482a50824d4",
		"28c7acab-d366-493a-8ac9-b1380570b4e8",
		"8216d5fb-4f3c-4fb7-8bd7-0559b82215d6",
		"91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
	},
	int64(0x0001)<<16 + int64(0x0001): { // Market
		"16286557-8ace-4023-8407-a7cc5a25f930",
		"520bf1d4-21ca-4014-a274-37cfaadf0c8b",
		"275f3d49-d32a-4e7f-90a8-ac7be054dc61",
		"dec26bae-fdf5-49fe-80e6-dde45e12a3ac",
		"91a16bc8-fdce-4b4d-bc09-cc5431cab9c7",
		"354e538a-7259-47f2-a7d3-dc72a0182d21",
	},
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

	var lastBeacon Beacon = GetLatestBeaconLog(0)
	fmt.Printf("Last Location: major=%d minor=%d\n", lastBeacon.Major, lastBeacon.Minor)

	if ids, ok := locationalIds[int64(lastBeacon.Major)<<16+int64(lastBeacon.Minor)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&[]Phrase{})
		return
	} else {
		json.NewEncoder(w).Encode(filterPhrasesById(phrases, ids))
	}
}
