package main

import(
  "crypto/rsa"
  "crypto/tls"
  "crypto/x509"
  "fmt"
  "log"
  "net/http"
  "net/url"

  "github.com/crewjam/saml/samlsp"
)

func main() {

  keyPair, err := tls.LoadX509KeyPair("./certs/public.crt", "./certs/private.key")
  if err != nil {
    panic(err) // TODO handle error
  }
  keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
  if err != nil {
    panic(err) // TODO handle error
  }

  idpMetadataURL, err := url.Parse("https://sso.tamacorp.co/saml/metadata")
  if err != nil {
    panic(err) // TODO handle error
  }

  rootURL, err := url.Parse("http://localhost:3000")
  if err != nil {
    panic(err) // TODO handle error
  }

  samlSP, _ := samlsp.New(samlsp.Options{
    URL:            *rootURL,
    Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
    Certificate:    keyPair.Leaf,
    IDPMetadataURL: idpMetadataURL,

  })
  app := http.HandlerFunc(login)
  http.Handle("/login", samlSP.RequireAccount(app))
  http.Handle("/saml/", samlSP)
  log.Println("RP is running on port 3000")
  http.ListenAndServe(":3000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
  Token := samlsp.Token(r.Context())
  log.Printf("%+v", Token)
  fmt.Fprintf(w, "Return, Token: %s, Attributes: %s", Token.StandardClaims, Token.Attributes)
}
