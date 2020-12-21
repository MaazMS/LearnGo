### how to setup local and workspace environment  on ubuntu  
1. Install go on local machine.   
2. open home/.bashrc file.    
3. Inside bashrc first setup go path in /usr/local/go/bin.   
4. This is very important to setup the workspace.  
5. Because inside our workspace we create the application, deploy and run the application.  
6. Inside work space create three folders src, bin and pkg.     
Q1. what is scr folder?    
It is the location where source code is  available.   
Q2. what is bin folder?   
Every time we run the application binary file is created and store in bin dir.   
Q3. what is pkg folder?     
For compile any thing it generate intermediate binary which mean sit is not full compile.      
`example`      
If i take third party library and intergrated with our application then pkg directory intermediate      binary are store.   
for reason that because they are recompile every time.    
when you compile your go application it check any source file in **pkg** dir and check it compiled if   compiled then link them to our application.      
7. link the path of workspace with golan. It is very important.    
This the whole example to local environment and workspace  
``` 
export PATH=$PATH:/usr/local/go/bin  
export GOPATH=/home/maaz/github/LearningGO/code
export PATH=$PATH:$GOPATH/bin
```  
* This is local environment path `/usr/local/go/bin `
* Tis is my workspace `/home/maaz/github/LearningGO/code`  
* This is workspace path `PATH=$PATH:$GOPATH/bin`   

