# Fifthsquare
An example web application for [my go foursquare API](https://github.com/aykutaras/gosquare). It mainly uses [net/http](http://golang.org/pkg/net/http/) and [html/template](http://golang.org/pkg/html/template/).

It gets registered user's checkIn data and saves it to a db and shows this data through a map.

## Usage
### Registration
* First you need to create an application for foursquare. Follow [these steps](https://developer.foursquare.com/overview/auth#registration) for registration.
* Replace your clientId and secret with example ones in server.go
* Google map or open map registrtion will be here

### Defaults
* Default port is **4001**. You can change it from server.go
* This application uses oauth registration. Default redirection address is **http://localhost:4001/code**
* Default db etc.

## TODO
Store user info first

### Database
BoltDB first maybe cayley after that

