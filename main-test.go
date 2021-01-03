package main

import (
	"fmt"
	"log"

	"github.com/bartvanbenthem/k8s-credsync/proxy"
	"github.com/bartvanbenthem/k8s-credsync/tenant"
)

func TestGetProxyCredentials() {
	// Prints the current proxycredentials
	pcurrent, err := proxy.AllProxyCredentials()
	if err != nil {
		log.Printf("\n%v\n")
	}
	fmt.Printf("\nProxy\n-----\n")
	for _, pc := range pcurrent.Users {
		fmt.Printf("User: %v Password: %v org: %v\n",
			pc.Username, pc.Password, pc.Orgid)
	}
}

// Prints the current tenant and proxy credentials
func TestMainFunctions() {
	// Prints the current tenant credentials
	tcurrent, err := tenant.AllTenantCredentials()
	fmt.Printf("\nTenant\n------\n")
	for _, tc := range tcurrent {
		fmt.Printf("User:%v Password:%v\n",
			tc.Client.BasicAuth.Username,
			tc.Client.BasicAuth.Password)
	}

	// Prints the current proxycredentials
	pcurrent, err := proxy.AllProxyCredentials()
	if err != nil {
		log.Printf("\n%v\n")
	}
	fmt.Printf("\nProxy\n-----\n")
	for _, pc := range pcurrent.Users {
		fmt.Printf("User:%v Password:%v org:%v\n",
			pc.Username, pc.Password, pc.Orgid)
	}
}
