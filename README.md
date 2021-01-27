# go-gin
Rest API in **golang** following best practices, built with gin, gorm, swagger and MVC architecture.

## What is an interface?
Interface in golang provides an easy way to create abstraction between different layers inside our application. It helps to bind different resources loosely to a point where impl of one layer is completely independent of other. It ensures principle of `seperation of concerns`.
<br>
An interface neither have any memory associated nor can be pointer by a pointer, for reference see `controller/video.go` **controller** struct. 

## When to avoid passing pointers to functions?
Methods and function that bind or marshal JSON structure to golang struct would not execute properly when passed pointers, as they expect interface. To optimize the function pass reference should be used using **&** operator.

## Makefile specs
- **git** - git add - commit - push commands
- **start** - runs the main.go file

## References
[gindump](https://github.com/tpkeeper/gin-dump)

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License