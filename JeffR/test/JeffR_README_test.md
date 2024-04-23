requires GOTMPDIR setting  
bash: `export GOTMPDIR=~/Projects/tmp`  
cmd: `set GOTMPDIR=%USERPROFILE%\Projects\tmp`  

# Supporting Stdin

NOTE if executing from a bash shell, you *must* use cat pipe output, otherwise ^D won't work to terminate input!

@see https://stackoverflow.com/questions/15673120/how-can-i-signal-eof-to-close-stdin-under-the-git-bash-terminal-on-windows-ctrl

Bash, manual input:  
`cat | go run JeffR_RussianNesting_Solution.go`  
// with cat piped, ^D works to terminate manual input

this is not an issue from Windows cmd line or when piping / redirecting file contents

Windows, manual input:  
`go run JeffR_RussianNesting_Solution.go`  
// ^Z works to terminate manual input

Read input from stdin (file redirection), bash or Windows:

`go run JeffR_RussianNesting_Solution.go < {filespec}`

e.g.

`go run JeffR_RussianNesting_Solution.go < samples/JeffR_Alek_Example1.txt`

(flip the slash(es) in the filepath on Windows)

-OR-

(bash specific:)

`cat samples/JeffR_Alek_Example1.txt | go run JeffR_RussianNesting_Solution.go`

(windows specific:)

`type samples\JeffR_Alek_Example1.txt | go run JeffR_RussianNesting_Solution.go`

# Envelope: Struct vs Array

When considering how to implement the solution, the basic question of whether to represent each envelope as a struct:

    type Envelope struct {
        width int
        height int
    }

or, as an array:

    type Envelope [2]int

it seems probable that, even tho the size of the array variant is pre-defined, array operations in general might be
less efficient than the equivalent struct ops (more overhead in general for arrays??)

Some tests were coded to attempt to prove this out; while the differences were relatively small, they were remarkably consistent
and seemed to support the theory that struct is the better approach here.

Early versions of my solution included support for testing creation (since we have to read from stdiin) 
and reading of envelope-as-array and envelope-as-struct;
test runs were generated using the following from Windows cmd prompt:  
```
    go run JeffR_RussianNesting_Test.go -t=3 | find "***" >JeffR_Test_Struct_vs_Array.log
    go run JeffR_RussianNesting_Test.go -t=10 | find "***" >>JeffR_Test_Struct_vs_Array.log
    go run JeffR_RussianNesting_Test.go -t=100 | find "***" >>JeffR_Test_Struct_vs_Array.log
    go run JeffR_RussianNesting_Test.go -t=1000 | find "***" >>JeffR_Test_Struct_vs_Array.log
    go run JeffR_RussianNesting_Test.go -t=10000 | find "***" >>JeffR_Test_Struct_vs_Array.log
    go run JeffR_RussianNesting_Test.go -t=100000 | find "***" >>JeffR_Test_Struct_vs_Array.log
```    
test run output can be found in JeffR_Test_Struct_vs_Array.log

