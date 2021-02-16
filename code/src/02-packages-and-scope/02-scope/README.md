### 
1. variable declaration constant declaration function declaration.

They are the same. 

1. file scope  
`import fmt` use in file scope where it declare.  
2. package scope  
* const ok = true  
* func main() {  
}  
So ok and main are the names which are visible throughout this main package.   
visible to all the files belong to package.  
but they won't call main of course because main function would be called by go runtime time itself.    
3. variable declare (block scope)
``` 
func main()
{
    var a = 1
}
``` 
##  
package scope, means that anything that is declared outside of any other functions body will
belong to the package the file belongs to.  
2. packet scope to share functionality between each other.  
3.Declarations need to be outside or any functions.  
   
### 
It says: "bye redeclared it in this block

the previous declaration at ./bye.go". 
1.  same scope, you can't have two declarations with the same name
### new 
importing only happens at the file level.
We import packages into a file not into a package.  
importing happens at the file scope only.  
multiple files in pakage level not use import fmt one file to another file.  

## rename import package
import "fmt" 
import f "fmt"