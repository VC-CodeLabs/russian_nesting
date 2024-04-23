NOTE this folder contains r&d stuff, not part of the actual solution

see the JeffR/test folder for a test harness that validates the actual solution with different test cases.

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

