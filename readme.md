atomic variable implemenation took 3 sec


mutex(multithread) implementation took 27 sec


single threaded implemenation took 37 seconds


we used atomic var bcoz we were not context switching many times over, else this would have not been a good option


in this program, our countLetters () is fast, we have to just wait for pages to load that is taking all the time.

mutex, when we apply lock to a process, we give control to the OS, then OS manages the resources of those mutexes and it wastes a few CPU cycles. Spinning locks improves this. It minimises control being transferred to OS