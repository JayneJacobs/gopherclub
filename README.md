# gopherclub.jaynejacobs.com
# gopherclub



usermod -aG wheel gopherclub
  139  visud
  140  visudo
  141  usermod -aG sudoer gopherclub
  142  visudo
  143  usermod -aG sudo gopherclub


Go prevents cross-sight scripting attacks
GopherJS converts Go code into JavaScript code

GopherJS compatibility table:

https://github.com/gopherjs/gopherjs/blob/master/doc/packages.md

```sh
go get -u github.com/gopherjs/gopherjs
```

* Build Js source file
    
```sh
gopherjs build
```

* Build and minify JS source file:

```sh
gopherjs build -m
```
    
   Compiles Go packages and start Http server
    
```sh
    gopherjs serve
```

   * use int instead of int8/16/32/64

   * Use float64 not float32

*JavaScript:*

```js
element = document.getElementById("navConteiner");
```
GopherJS:

```go
element := js.Global.Get("document").Call("getElementById", "navContainer")
```

* DOM Binding:

```go
element := dom.GetWindow().Document().GetElementByID("navContainer")
```

* Alias calls to js.Global and Dom

```go
var JS = js.Global
var D = dom.GetWindow().Document()
```
### Changing CSS Style

* JavaScript

```js
element = document.GetElementByID("modalLayer");
element.style.display = "none";
```

* GopherJS:
```go
var JS = js.Global
element := js.Global
element := JS.Get("document").Call("getElementById", "modalLayer")
element.Get("style").Set(display", "none")
```

Move .js generated files to /static in project root

# Making XHR Calls
* XHR object
  
  ### XML HTTP Request Object

  Data retreieved from a URL without causing full refresh
  xhr binding
  to install
  
  ```sh
  go get -u honnef.co/go/js/xhr 
```
* build file by referencing the root directory
* run the app from the root of the project

Let's Encrypt Automated Certificate Management
   * Allows automatic deployment of free SSL trusted certs

   ```sh
     65  cd /etc/letsencrypt/live/
   77  ls /etc/letsencrypt/live/gopherface.jaynejacobs.com/
   80  ls /etc/letsencrypt/live/gopherface.jaynejacobs.com/
   92  cd /opt/letsencrypt
    ./letsencrypt-auto certonly --standalone -d gopherclub.jaynejacobs.com
    ```
Performs domain name validation

## To build the binary for Linux
```sh
go build -o $GOPHERCLUB_APP_ROOT/builds/gopherclub-linux64

file builds/gopherclub-linux64
builds/gopherclub-linux64: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

#Building Binarys for other os's

https://golang.org/doc/install/source#go14

[Preparing the Bundle](PreparingBundle.md)

## To make cookied work
root@gopherclub:/tmp# chown gopherclub gopherclub-sessions