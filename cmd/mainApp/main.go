package mainApp


import (
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"strings"
)






var cu Currency


func getCurrency(URLCurrency string) error {

	content, err := http.Get(URLCurrency)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(content.Body)
	if err != nil {
		return err
	}
	defer content.Body.Close()
	//Cu := Currency{}

	err = json.Unmarshal([]byte(body), &Cu)
	if err != nil {
		fmt.Printf("Unmarshaling currency failed %v :", err)
	}



	return nil
}

func handlerRequest(w http.ResponseWriter, r * http.Request) {

	url := strings.Split(r.URL.Path, "/")
	if len(url) != 3 {
		http.Error(w, "Splitscreen sadness", 400)
	}else if url[2] == "" {
		http.Error(w, "this is Sparta!", 400)
	}else{
		fmt.Fprintf(w, "Now we are going somewhere: %s \n", url[2])
		switch r.Method {
		case "GET":
		case "DELETE":
			w.WriteHeader(http.StatusAccepted)
		default:
			http.Error(w, "you have joined the darkside, darthVader is your father now", 400)
		}
	}


	/*
	w.Header().Set("Content-Type", "application/json")
	url := "http://api.fixer.io/latest?base=EUR"
	err := getCurrency(url)
	if err != nil {
		fmt.Printf("Something went wrong with extracting currency from URL %v", err)
	}
	json.NewEncoder(w).Encode(&Cu)


	switch r.Method {

	case http.MethodPost:



	}
*/
}
var Cu Currency


func RegistrationOfNewWebHook(w http.ResponseWriter, r * http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Something went wrong with reading webhook input %v", err)
	}
	defer r.Body.Close()
	payUrDues := Payload{}
	err = json.Unmarshal(body, &payUrDues)
	if err != nil {
		fmt.Printf("somethin went wrong in externalInput/Unmarshal body to Payload %v", err)
	}
	fmt.Fprintf(w, "%s, %s, %f", payUrDues.WebhookURL, payUrDues.TargetCurrency, payUrDues.MaxTriggerValue)


}
func latestCurrency(w http.ResponseWriter, r * http.Request) {


}
func averageCurrency(w http.ResponseWriter, r * http.Request) {


}

func addemdum(w http.ResponseWriter, r * http.Request) {


}


func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/root", handlerRequest)
	http.HandleFunc("/root", RegistrationOfNewWebHook)
	http.HandleFunc("/root/latest", latestCurrency)
	http.HandleFunc("/root/average", averageCurrency)
	http.HandleFunc("/root/evaluationtrigger", addemdum)


	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}