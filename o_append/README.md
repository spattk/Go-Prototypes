The idea is inspired from the GFS paper which implements atomic record append. I have tried to implement a similar structure using Unix's `O_APPEND` mode.

#### How to run
```
cd o_append
go run o_append.go
```

#### Code Explanation
`o_append.go` -> driver code / main function

`func no_o_append()` -> writes to a log file (named log_no_oappend.txt) without `O_APPEND` mode 

`func o_append()` -> writes to a log file (named log_oappend.txt) with `O_APPEND` mode 

#### Output Explanation
* Both functions are run through a for-loop from 1 to 10, so the file should effectively contain 10 lines of output.

* The order is not maintained because the writing sequence depends on how the OS spawns different goroutines.

* However, the interesting observation is that while `log_oappend.txt` contains all 10 lines in an interspersed order, log_no_oappend.txt does not have all 10 lines in the output.

* This happens because the data gets overridden when multiple writers write simultaneously, and the ones finishing later perform the final flush.

* This can be verified by adding some sleep between each goroutine call, which will ensure that `log_no_oappend.txt` prints output from all 10 goroutine calls.