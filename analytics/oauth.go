package analytics

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func OauthConnect() *http.Client {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	conf := &jwt.Config{
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
		PrivateKey: []byte("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDkKHYnxKZLoCQB\nrykgDKl3GsszKhrDpm9uCEDRe8AvDpq2Hiv9nXI3AxpnLWVIvL5cvkOGdF+biTVW\nwA6L4UvnlkOnnSGdGsjHmCfLLHACy3htuoeMtWEHeTq564TGa/asgoE//fXrTw+B\ni2FRUPp+DORRdPnwX/FLodK9NEQRKlptOmeGjZG6GBzzIyXGBjy+ilghIQOHgW6Q\nAKyWFjt+uzamesU01HqvorS/t+/2uf6YUiVm0tyNrY6nOUAiKGNGsyQ99NeHaHiV\nqvjmvcXboy2TL8f6vdiSX/JNrc98JwVwvC88w2V8dBZ7zMQFaMxtVAlYuK2gbaQ4\nBFJSlqXtAgMBAAECggEBAJcyEQrNkH4xQ2syToiEm5DoWpwvXMkm9FTwdzPbVEJa\n5T+mwvQwcE3waoDRXntTSLthApKJcAo8ZPJ2F62uwYRY2cCrqmjYJRHPzR8g9aGE\ngOwGXmlqkUBIoBgw8S4kwbUUTt6/R0aGf/oTCYy/oyABHxSz9iXLDrZWKLdleW39\naF0+1uaPmq+clkNh7iLXNrKiULFbSCUG+Isj54uZrDuHwc9ilCVrMwxv8sVVEpkI\n1iYldBT2WGx4HXCq1HvMvqe4agAiMDUrt3xIGLSo3BxxcgAvPdz9j6DFHk8ELkEg\nyRUF6H1OTArKLZNW05LAAw1NZM43/J4x/2Jv+BZ9EYECgYEA/2qXVOq12NYh5IYA\ndnlXHN9ZoOHhMsajIOHJNO+Y7PQBlt9c5pM09M3NyUU43NVJ4RszOh+88MlZllV1\n5aY1/kMQd4LXiEHhhX8VegWKBT2+0nciJeV/5ir7pEaT6uTwxpWDudkbam/n0uNj\nVVBSJoUA+XlAOre6L8XW6pdFf70CgYEA5K3s4h+wVt5s2s8yKqr0FOKjdoPfqqY0\n6JmqaCyYTgciqDA/3D/9Yopl4hM2j3s0EP+RwVRkmt2KrAEHpA5fMYFD83lXkb4D\nFbNzts17SLbOEmiRop+4UF9rAaU8r9OpAtPSrKlBGeUHk2MBQjrKet6ekgncALUA\nkdtyJ/M2yfECgYEAi/qR12MOqDiquIBMDqHvgQ5TI53HerlsO6AiUlWhXPNZT6M4\nWdGBa0zDJDRAtp5Z7/mlOIyeUuQ2qfJKO8K8wo1kn+rgTne0riH41jfk7XxqcRQD\nbft5/d2+rchUVp6D6oGUgK0gnlOYHHasQNxVOXEZGKA/+C4q5vWpFEY1O40CgYAN\nNQJOWPgtPDFbPcYu8qqgHirEqH8dcvXxTRarZ4/wSbc1sXuus79lloql0OhAxyPN\nq4T2fnJFA3CD2JdhYTccO9P51tM45pl1AmU3dIyfOXeTVpi6pLDXa3tZn/puSTIk\nqKktMaVB4plaaMhk09Jn4D2WP52GKO38fokMR56tcQKBgBEhwYBvEBHzxDudSDjB\nBqIwlYi+J1Lr5hRL/BtzSYbPkSwSQmnaoozsPSiVfo7p945Lze3xEDuEf008Cnk1\nrQJhhsFuy7oa4O1z++LiIieU5iVEAH5ZNBHdXSb8GcTA/9g1v3bzZX3hGna4UZpM\n0cBXFp6/84oxJCQn79uOwN6M\n-----END PRIVATE KEY-----\n"),
		Scopes: []string{
			"https://www.googleapis.com/auth/analytics",
		},
		TokenURL: google.JWTTokenURL,
	}
	client := conf.Client(oauth2.NoContext)
	return client
}
