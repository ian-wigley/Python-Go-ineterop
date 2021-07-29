import ctypes
import os

## Python <-> Go interop example

go_object = ctypes.cdll.LoadLibrary(os.getcwd() + '/_go-add.so')
print(go_object.add(10))