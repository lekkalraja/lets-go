

## Five Important Questions:

#### 1. How do we run the code in our project ?

###### `Go CLI`

* `go build` -> Compiles a bunch of go source code files and outputs executable file
    ```shellscript
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ls -ltr
        total 8
        -rw-rw-r-- 1 raja raja  75 May 28 17:13 main.go
        -rw-rw-r-- 1 raja raja 655 May 30 10:41 README.md
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ go build main.go 
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ls -ltr
        total 1904
        -rw-rw-r-- 1 raja raja      75 May 28 17:13 main.go
        -rw-rw-r-- 1 raja raja     655 May 30 10:41 README.md
        -rwxrwxr-x 1 raja raja 1937687 May 30 10:42 main
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ./main 
        Hello, Golang!
    ```
* `go run` -> Compiles and executes one or more files
    ```shellscript
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ go run main.go 
        Hello, Golang!
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ 
    ```
* `go fmt` -> Formats all the code in each file in the current directory
* `go install` -> Compiles and installs a package
* `go get` -> Downloads the raw source code of someone else's package
* `go test` -> Runs any tests associated with the current project

#### 2. What does `package main` mean ?

* `package` is workspace/project which contains one or more source files (*.go) which are created for specific task and each source file (.go) should contain it's `first line` as package declaration.

##### Types of Packages
1. Executable -> `package main` (`main` is special)
   *    Defines a package that can be compiled and then `executed`. **Must have a func called `main`**
   *    `package main` -> `go build` -> `main.exe` (if we ran this file, the function named `main` would be called automatically ran)
  ```shellscript
    raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ls -ltr
    total 8
    -rw-rw-r-- 1 raja raja   75 May 28 17:13 main.go
    -rw-rw-r-- 1 raja raja 2333 May 30 10:56 README.md
    raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ go build main.go 
    raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ls -ltr
    total 1904
    -rw-rw-r-- 1 raja raja      75 May 28 17:13 main.go
    -rw-rw-r-- 1 raja raja    2333 May 30 10:56 README.md
    -rwxrwxr-x 1 raja raja 1937687 May 30 10:56 main
    raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ./main 
    Hello, Golang!
    raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ 
  ```

2. Reusable
   * `package repository` -> Defins a package that can be used as a dependency (helper code)
   * `package utils` -> Defines a package that can be used as a dependency (helper code)
   *    `package repository` -> `go build repository.go` -> `nothing` (Compiling a non-main package gives)

    ```shellscript
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ls -ltr
        total 1908
        -rw-rw-r-- 1 raja raja      75 May 28 17:13 main.go
        -rwxrwxr-x 1 raja raja 1937687 May 30 10:56 main
        -rw-rw-r-- 1 raja raja    3098 May 30 10:57 README.md
        -rw-rw-r-- 1 raja raja      93 May 30 10:58 repository.go
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ go build repository.go
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$ ls -ltr
        total 1908
        -rw-rw-r-- 1 raja raja      75 May 28 17:13 main.go
        -rwxrwxr-x 1 raja raja 1937687 May 30 10:56 main
        -rw-rw-r-- 1 raja raja    3098 May 30 10:57 README.md
        -rw-rw-r-- 1 raja raja      93 May 30 10:58 repository.go
        raja@raja-Latitude-3460:~/Documents/coding/golang/learn-go/hello-world$
    ```

#### 3. What does `import fmt` mean ?

* Gives access to other package (standard/3rd party libraries) code to re-use

  [Standard Libraries](https://golang.org/pkg/)

#### 4. What's that `func` thing ?

* To define functions in golang

#### 5. How is the main.go file organized ?

1. Package Declaration
2. Import other packages that we need
3. Declare variables, functions

##### Questions :
1. What is the purpose of a package in Go?
   Ans) `To Group together code with a similar purpose`
2. What is the special name we use for a package to tell Go that we want it to be turned into a file that can be executed?
   Ans) `main`
3. The one requirement of packages named "main" is that we...
    Ans) `Define a function named "main", which is ran automatically when the programs runs`
4. Why do we use "import" statements?
    Ans) `To give our package access to code written in another package.`