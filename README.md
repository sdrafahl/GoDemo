#GoPi

Simple Golang project that computes PI using the Leibniz infinite series.

https://en.wikipedia.org/wiki/Leibniz_formula_for_%CF%80

##Setup

1. Make Directory Structure

```

ProjectName -> bin

ProjectName -> src

src -> github.com

github.com -> <Github Username>

<Github Username> -> <Namespace>

```

2. Export OS Variables

For this project here are the commands run to setup the system variables.

```
export GOPATH=$HOME/code/goPi

```

```
export PATH=$PATH:GOPATH/code/goPi/bin

```


3. To Install

```

go install github.com/sdrafahl/pi

```

This installs the go project in the hello namespace and put it in the bin
directory.


4. To Run

After installing there should be an executable in the bin directory.


Tutorial Video

```
https://youtu.be/2KmHtgtEZ1s

```

##Running

After you have compiled the project you can navigate to the bin directory  
and execute the pi executable. The first command line argument to give is the
number of iterations.

For example

```
./pi 1000
```
Runs the program 1000 iterations for the Leibniz forumula.
