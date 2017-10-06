package youbeacomm 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"BeaconPassedPost",
		"POST",
		"/shield-9/Youbeacomm/1.0.0-alpha/beacon/passed/",
		BeaconPassedPost,
	},

	Route{
		"DevicePost",
		"POST",
		"/shield-9/Youbeacomm/1.0.0-alpha/device/",
		DevicePost,
	},

	Route{
		"RootGet",
		"GET",
		"/shield-9/Youbeacomm/1.0.0-alpha/",
		RootGet,
	},

	Route{
		"PhrasePhraseIdGet",
		"GET",
		"/shield-9/Youbeacomm/1.0.0-alpha/phrase/{phraseId}/",
		PhrasePhraseIdGet,
	},

	Route{
		"PhrasePhraseIdResponsesGet",
		"GET",
		"/shield-9/Youbeacomm/1.0.0-alpha/phrase/{phraseId}/responses/",
		PhrasePhraseIdResponsesGet,
	},

	Route{
		"PhraseSuggestionsFrequentGet",
		"GET",
		"/shield-9/Youbeacomm/1.0.0-alpha/phrase/suggestions/frequent/",
		PhraseSuggestionsFrequentGet,
	},

	Route{
		"PhraseSuggestionsLocationDeviceIdGet",
		"GET",
		"/shield-9/Youbeacomm/1.0.0-alpha/phrase/suggestions/location/{deviceId}/",
		PhraseSuggestionsLocationDeviceIdGet,
	},

}