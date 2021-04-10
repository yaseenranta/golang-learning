# golang-learning

```
//create module
go mod init example.com/greetings
//replace module with your local module
go mod edit -replace=example.com/greetings=../greetings
//synchronize the module
go mod tidy
