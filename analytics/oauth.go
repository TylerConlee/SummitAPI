package analytics

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// OauthConnect grabs the private key out of the local private-key.txt file,
// and sends an authetication request to Google for authetication
func OauthConnect() *http.Client {
	dat, _ := ioutil.ReadFile("private-key.txt")
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	conf := &jwt.Config{

		// Every account that this email address has access to will be included
		// in data requests for this application
		Email: "analytics-access@summitppc-1216.iam.gserviceaccount.com",
		// The contents of your RSA private key or your PEM file
		// that contains a private key.
		// If you have a p12 file instead, you
		// can use `openssl` to export the private key into a pem file.
		//
		//    $ openssl pkcs12 -in key.p12 -passin pass:notasecret -out key.pem -nodes
		//
		// The field only supports PEM containers with no passphrase.
		// The openssl command will convert p12 keys to passphrase-less PEM containers.
		PrivateKey: dat,
		Scopes: []string{
			"https://www.googleapis.com/auth/analytics",
		},
		TokenURL: google.JWTTokenURL,
	}

	// Create an oauth autheticated client
	client := conf.Client(oauth2.NoContext)
	return client
}
