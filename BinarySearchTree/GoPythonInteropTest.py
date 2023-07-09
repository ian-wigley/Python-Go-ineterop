#################################################################
#       Go! binary search tree - Python interoperation          #
#                                                               #
# Build the Go shared object first with the following command : #
#                                                               #
#         go build -buildmode=c-shared -o _interop.so           #
#                                                               #
#             Author : Ian Wigley 8th July 2022                 #
#################################################################

import ctypes
from ctypes import *

so = ctypes.cdll.LoadLibrary('./_interop.so')

pyarr = [1,2,3,5,7,8,55,10,100,101,102,103]
so.binaryTree.argtypes = [c_int * len(pyarr)]
so.binaryTree.restype = ctypes.c_void_p

# Make C array from py array
arr = (c_int * len(pyarr))(*pyarr)
ptr = so.binaryTree(arr, len(arr), 2)
out = ctypes.string_at(ptr)
print('Result from Go! : ' + out.decode('utf-8'))
