# Ascii-Art-Web

***
```
                      _   _                           _                                _      
                     (_) (_)                         | |                              | |     
  __ _   ___    ___   _   _   ______    __ _   _ __  | |_   ______  __      __   ___  | |__   
 / _` | / __|  / __| | | | | |______|  / _` | | '__| | __| |______| \ \ /\ / /  / _ \ | '_ \  
| (_| | \__ \ | (__  | | | |          | (_| | | |    \ |_            \ V  V /  |  __/ | |_) | 
 \__,_| |___/  \___| |_| |_|           \__,_| |_|     \__|            \_/\_/    \___| |_.__/  
                                                                                              
                                                                                              
```         

## Table of Contents
1. [Description](#description)
2. [Authors](#authors)
3. [Usage:how to run](#usage-how-to-run)
4. [Implementation details: algorithm](#implementation-details-algorithm)

### Description:
***
Hi *Talent*!
Ascii-art-web consists in creating and running a server which, based on the data entered by the user via a form, displays the result in a graphic representation using ASCII.
This project's servor (back-end) is written using GO and the web documents (front-end) using HTML & CSS

#### Go,Html & Css:
***
- **GO**, also called Golang or Go language, is an open source programming language that Google developed.
* **HTML**,The HyperText Markup Language is the standard markup language for documents designed to be displayed in a web browser.
+ **CSS**,Cascading Style Sheets, form a computer language that describes the presentation of HTML documents.

#### Ascii-art-fs:
***
As a reminder,Ascii-art-fs is a program which consists in receiving a string and a banner as an argument and outputting the string in a graphic representation using ASCII.If you want an overview, consult it at https://learn.zone01dakar.sn/git/vindour/ascii-art-fs


### Authors:
***
- Vincent Félix Ndour (our super capitain)-Backend- https://learn.zone01dakar.sn/git/vindour
* Masseck Thiaw (the genius) -Frontend- https://learn.zone01dakar.sn/git/mthiaw
+ Seynabou Niang (writer & debugger) -Test- https://learn.zone01dakar.sn/git/sniang

### Usage:how to run
***
#### server
A little intro about the installation.
```
$ git clone https://learn.zone01dakar.sn/git/vindour/ascii-art-web
$ cd ascii-art-web
$ go run main.go
```
#### tests
to run the tests, open a new terminal and enter the command
```
$ go test -v
```
For a specific test:
```
$ go test -v -run=nameOfTheTest
```
you can check the test coverage with:
```
$ go test -cover
```

### Implementation details: algorithm
***
here a brief overview of our structure
```
.
├── templates
│   ├── error400.html
│   └── error404.html
│   └── error405.html
    └── error500.html
│   └── index.html
├── main.go
```
#### Starting a web server in the main function of our main.go file:
- Creation of the server with the ListenAndServe() method of the net/http package which takes the port in string format as an argument, note that the two previous points are very important, and a handler which is not useful to us here so it will be nil. The Fatal() method of the log package will allow us to catch any error when creating our server.
* display on the terminal after the launch of the server with the command 
```
go run main.go
```
a message with the link that sends us to a browser
```
log.Println("Server started on http://localhost:8080")
```
+ creating a static file server: We create the server with the FileServer() method which takes an http.FileSystem as an argument. In order for this server to host our files in the templates folder, it must be transformed with the Dir() method. You must specify the Handle route.
```
fs := http.FileServer(http.Dir("templates"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```
- add route handlers to the web server with the method HandleFunc():
```
http.HandleFunc("/", indexHandler)
http.HandleFunc("/ascii-art", asciiArtHandler)
```
##### indexHandler & asciiArtHandler functions:
###### indexHandler:
```
func indexHandler(w http.ResponseWriter, r *http.Request) {
    ...
}
```
- Check if the path included in the request is valid . If it is not valid, use the ParseFiles() method of the template package to find our error404.html file ready for this purpose and pass it in response with the Execute() method of the template.If the method are not allowed => error405.html
* we continue if the path is good by using the template package again to retrieve our index.html file. If an error occurs, return our index500.html page as response.


###### asciiArtHandler
```
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
    ...
}
```
- check if method is different from POST if so return error405.html.
+ check if the text is not printable => error400.html 
* if the method is POST, check if the text is not empty and if the given banner exists. If only one condition is missing return error404.html .
+ now if there is no error, use our function of transformation of a string into ascii graphic representation to return it in our index.
