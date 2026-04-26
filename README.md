# make
use case of make keyword in Golang

## Why make exists separately

Because **slices/maps/channels** have internal runtime structures.   
Example slice internally has:  
* pointer to array   
* len   
* cap   

**Map** has hash table internals.   
**Channel** has queue + synchronization primitives.

So Go needs runtime setup.