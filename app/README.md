Installing
======
1. Download, compile and install nanomsg from http://nanomsg.org/download.html
2. Download, compile and install nanomsgxx from https://github.com/achille-roussel/nanomsgxx
3. Open the project in Qt Creator and build (or use qmake && make from the command line)

Running
======
Run from Qt Creator or execute the generated binary from command line. A nanomsg socket will
start listening with PAIR protocol on tcp port 40899. It will successfully parse data of the form-

10,-20,30

Which are the accelerometer readings with respect to x, y and z axes.

**NOTE** As of now an incorrect data string will cause the app to crash

