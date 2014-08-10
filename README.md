# Description

This program serves as an example of what the [GoController](http://github.com/thefryscorer/gocontroller) library is capable of, and a general reference for programming customised controllers using the library.
The controller is hosted locally as a webpage, which can then be accessed using any browser- mobile or otherwise- to control games over the network.

# Dependencies

Currently only works with Linux with the xte binary installed, I am looking into making it work for Windows, possibly utilising system calls to autohotkey.

# Installation

    # go get github.com/thefryscorer/gocontroller-example
    
Then to run the program, if your Gopath's bin directory is added to your path:
    
    # gocontroller-example
    
Or if your path does not include $GOPATH/bin:

    # cd $GOPATH/bin/
    # ./gocontroller-example
