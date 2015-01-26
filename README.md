##Adrenaline

A command line HTTP performance tool built in Go.

There are many other tools out there that do roughly the same thing.  In exploring Go, I thought it would be interesting to try and write my own.  The command line arguments provide a good summary of its existing capabilities:

    > ./adrenaline --help
    Usage of ./adrenaline:
      -count=1: The number of clients that should run in parallel.
      -duration=10: Duration of test run (in seconds).
      -har="": The HAR file to be used for playback.
      -host="localhost": Specify the endpoint for issuing requests.

You can specify the target host, test duration, degree of parallelism and a traffic source (in HAR format).

Coming soon:
* support for multiple hosts
* support for traffic "profiles"
* multi-machine traffic generation (over SSH)

