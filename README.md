# dictionary
A CLI dictionary app. Create your own dictionary and save everything you want. Anything but text only :)

Golang version 1.18 
# How to run this project
```shell
git clone https://github.com/kosher9/dictionary.git
cd dictionary
go mod download
go build -o dict
```
// Add to the dictionary
```shell
./dict -action add "Golang" "What a beautiful language"
#'Golang' added to the dictionary
```

// List element in your dictionary

```shell
./dict -action list
#Golang          Un super langage                                        Jul  4 23:27:22
```