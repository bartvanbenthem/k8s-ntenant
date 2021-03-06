package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bartvanbenthem/k8s-ntenant/sync"
)

func main() {
	// get from environment variables
	address := os.Getenv("K8S_SERVER_ADDRESS")
	cert := os.Getenv("K8S_SERVER_CERT")
	key := os.Getenv("K8S_SERVER_KEY")

	// http handler functions
	http.HandleFunc("/", HandlerDefault)
	http.HandleFunc("/credential/sync", HandlerCredentialSync())
	http.HandleFunc("/grafana/sync", HandlerGrafanaSync())
	http.HandleFunc("/ldap/sync", HandlerLDAPSync())

	// check if certs for tls server are provided
	if len(cert) != 0 && len(key) != 0 {
		// listen and serve https connections
		log.Printf("About to listen on https://%v/\n", address)
		err := http.ListenAndServeTLS(address, cert, key, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	// listen and serve http connections when no certs are provided
	log.Printf("About to listen on http://%v/\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// default handler
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"server":"running"}`)
}

// handler for credential synchronization service
func HandlerCredentialSync() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		err := sync.Credential()
		if err != nil {
			log.Printf("Error: %v", err)
			io.WriteString(w, `{"credential":"sync finished with errors inspect log"}`)
		} else {
			log.Printf("credential sync finished")
			io.WriteString(w, `{"credential":" sync finished"}`)
		}
	})
}

// handler for grafana synchronization service
func HandlerGrafanaSync() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		err := sync.Grafana()
		if err != nil {
			log.Printf("Error: %v", err)
			io.WriteString(w, `{"grafana":"sync finished with errors inspect log"}`)
		} else {
			log.Printf("Grafana synchronization finished")
			io.WriteString(w, `{"grafana":" sync finished"}`)
		}
	})
}

// handler for ldap synchronization service
func HandlerLDAPSync() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		err := sync.LDAP()
		if err != nil {
			log.Printf("Error: %v", err)
			io.WriteString(w, `{"ldap":"sync finished with errors inspect log"}`)
		} else {
			log.Printf("LDAP synchronization finished")
			io.WriteString(w, `{"ldap":" sync finished"}`)
		}
	})
}
