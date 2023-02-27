the hard ware is 1 core processor

1 thread from hw perspective

the OS multiplexes the one thread to multiple OS threads

N:M(N us number of cores/hw threads : M is OS threads)

Thread in Java --> It does a kernal call askes the OS to get a thread

Thread contains some memory as well.

Everytime the switching between threads is every time to make a call to the OS

- From performance perspective , calling system calls is a sluggish operation.

- this concept is called scheduling

- green threads : the runtime creates and schedules the thread.

- Go routines are green threads(coroutines)
- Go runtime has a schedular
- Gorountines are very small in size.

- 64 mb of memory --> 65 mb : 64 mb of memory Total 128 mb, therefore 63 mb memory is unused. (hypo)
- Gorountes are small in size and then can grow or shrink based on demand.
- You can create fee thousand Go routines per thread(internally it uses a thread because OS knows only about a thread)