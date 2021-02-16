## What is package?  
* First rule of being a package is that all files that belong to a package should be in the same directory.  
* Second rule is that all files that are in the same folder should belong to the same package.   
* The package clause should be at the first line of code. So you can't use it anywhere beyond the first line.   

## How to execute same package of multiple go scource files.  
1.`go run .`    
2. ` go run *.go`   
3. `go run main.go hey.go bye.go`    
* example    
```  
maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code/src/02-packages-and-scope/exercises/Package$ go run .
Hi 
goodbye
maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code/src/02-packages-and-scope/exercises/Package$ go run *.go
Hi 
goodbye
maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code/src/02-packages-and-scope/exercises/Package$ go run main.go hey.go bye.go
Hi 
goodbye
```   

## there are only two kinds of packages in Go.  
Executable packages and library packages.     
1. a file should belong to the package "main"   
2. that file should have a function named "main".      
1. . As you know, its name should always be "package main", and executable package 
should also contain a "main function", but only once.    
2. A library package is created only for usability. So you can import them from other packages    

## difference library packages and Executable packages    

library | Executable | 
--- | --- |
created for resusability | created for running  
non-executing |executing  
importable | non-importable  
can have any name | name should be main 
no func main | func main   


   
