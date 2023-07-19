# Python Go interop Test
Small examples of Python interoperation with Go!
<br>
<br>
## Example 1
> Example of passing Integers and Strings between Python and Go!
<br>
<br>
## Example 2
> Recursive binary tree example.
> For more information about binary trees take a look at the following site: https://en.wikipedia.org/wiki/Binary_search_algorithm
> <br>
> This example was originally written for a small coding challenge, rather than using Python, I thought it would be good to have a play with Go! The challenge was to return an ordered & formatted range of numbers from any sequence, i.e. 100,1,2,3,4,5,8,9
> would yield 1..5, 8..9, 100
> <br>
> <br>
> Clone the repo, cd ./BinarySearchTree/ folder before open the local version in VSCode (or similar) then build the Golang static object library (.so) enter:
> <br> go build -buildmode=c-shared -o _interop.so
> <br>
> <br>Finally, to see the Python - Go interop in action run the GoPythonInteropTest.py code.
