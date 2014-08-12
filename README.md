* lm

** Find the last modified files in a directory hierarchy.

I was on a server that used the file system as its datastore and I
wanted to see what files had been modified in the heirarchy of
files. I'm sure there is a bash incantation I could have come up with,
but this was more fun.

** Usage

    $ lm /path/to/files

To see the last N number just pipe it to head:

    $ lm /path/to/files | head -n 10
