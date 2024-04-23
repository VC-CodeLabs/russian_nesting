requires GOTMPDIR setting  
bash: `export GOTMPDIR=~/Projects/tmp`  
cmd: `set GOTMPDIR=%USERPROFILE%\Projects\tmp`  

# Supporting Stdin

NOTE if executing from a bash shell, you *must* use cat to pipe manual keyboard input, otherwise ^D won't work to terminate input!

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

