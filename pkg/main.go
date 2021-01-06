package main

import (
	"fmt"
	"github.com/trisolaria/raphael/pkg/crypt"
	"github.com/trisolaria/raphael/pkg/session"
)

func main() {
	//Some small tests
	userSession := session.UserSession{}
	printAuthentication(&userSession,"a","b")

	indeterminantAuth := crypt.IndeterminantAuthenticator{}
	printAuthentication(&indeterminantAuth,"a","b")
}

func printAuthentication(a crypt.Authenticator, user, pass string){
	fmt.Printf("auth = %v \n", a.Authenticate(user, pass))
}