# go-gin
Rest API in **golang** following best practices, built with gin, gorm(sqlite), swagger and MVC architecture.

## What is an interface?
Interface in golang provides an easy way to create abstraction between different layers inside our application. It helps to bind different resources loosely to a point where impl of one layer is completely independent of other. It ensures principle of `seperation of concerns`.
<br>
An interface neither have any memory associated nor can be pointer by a pointer, for reference see `controller/video.go` **controller** struct. 

## When to avoid passing pointers to functions?
Methods and function that bind or marshal JSON structure to golang struct would not execute properly when passed pointers, as they expect interface. To optimize the function pass, reference should be used using **&** operator.

## CRUD with gorm and some *TODO's*
- implementation of DB close method is left
- the update method requires whole video entity for update, **as we are using save method**, better way is to use **Update()** method.
- the error handling not setup right
- **Upsert** strategy is basically checking first that there are no users with that email, in that case insert that element, otherwise update the existing element. **fix for foreign key relation**

## Makefile specs
- **git** - git add - commit - push commands
- **start** - runs the main.go file

## References
[gindump](https://github.com/tpkeeper/gin-dump)
[jwt-go](https://github.com/dgrijalva/jwt-go)
[gorm-docs](https://gorm.io/docs)
[upsert-conflict](https://gorm.io/docs/create.html#upsert)

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License