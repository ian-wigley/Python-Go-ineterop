import ctypes
import os

## Python <-> Go interop example

go_object = ctypes.cdll.LoadLibrary(os.getcwd() + '/_go-add.so')
print(go_object.add_test(10))

string_test = go_object.string_test
string_test.restype = ctypes.c_void_p
ptr = string_test('Hello from Python !!'.encode('utf-8'))
out = ctypes.string_at(ptr)
print(out.decode('utf-8'))