## Go + v8 integration sample

This sample integrates v8 with go utilizing a wrapper. V8 runs fine on its own but starts fail when context switches are introduced. We discovered this problem when we tried runnning v8 from a server. This sample reproduces the issue in a simpler manner by running a separte dummy goroutine. 

####Setup
 - move the v8runner folder into your GOROOT folder
 - add a copy of libv8wrapper.dylib to /usr/lib/ 
 - ```go run sample.go```

