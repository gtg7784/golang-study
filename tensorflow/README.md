# Tensorflow with Go

## install and test tensorflow package

before install, you should install Tensorflow C Library for Tensorflow go.

install Tensorflow C Library at this [link](https://www.tensorflow.org/install/lang_c).

after install Tensorflow C Library, You should execute following command lines.

```sh
sudo update_dyld_shared_cache # OS X
```

if you waana test for install Tensorflow C Lib?

make tensorflow.c file and write following code.

```c
#include <stdio.h>
#include <tensorflow/c/c_api.h>

int main() {
  printf("Hello from TensorFlow C library version %s\n", TF_Version());
  return 0;
}
```

and, build it with follow command line.
```sh
gcc -I/usr/local/include -L/usr/local/lib tensorflow.c -ltensorflow
```

you can build a.c successfully?
now, installing Tensorflow C Lib is finish.

let's try to install Tensorflow Go Lib.

```sh
go get github.com/tensorflow/tensorflow/tensorflow/go
```

test
```sh
go test github.com/tensorflow/tensorflow/tensorflow/go
```

