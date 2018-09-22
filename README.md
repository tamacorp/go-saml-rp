# GO SAML RP (Relying Party)

### Project Summary

GO SAML RP is a Go repository contains exmaple of code that works as an RP connecting to TAMA via SAML 2.0 Protocol

The example code uses [saml] for SAML 2.0 library

### Installation
```bashp
go get ./...
```

### Run
```bashp
go run server.go
```

### Usage
To try using SAML, please go to
```bashp
http://localhost:3000/login
```

[saml]: https://github.com/crewjam/saml
