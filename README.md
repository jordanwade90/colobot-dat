colobot-dat: read game data files from Colobot
==============================================

Release builds of Ceebot and Colobot store their game files in DAT containers.
This project consists of (1) a library for reading DAT files and (2) a
command-line utility to extract files from a DAT container.

Library
=======

Read a container file by passing an `io.ReaderAt` and a `Codec` to `New()`.  A
`Codec` XORs data in the container with a predefined table; `Codec`s for the
demo and full versions of Ceebot and Colobot are included with the library.
(The tables were extracted from the Colobot source code at
[github.com/adiblol/colobot].)  The resulting `Container` contains a slice of
`File`s, each of which has an `io.SectionReader` which will read the data from
that file.

Command-line extractor
======================

`colobot-dat` is an extractor for DAT files.  Given a path to a DAT file, it
will extract each file into the current directory, unless a file with the same
name already exists.  (If the `-l` flag is given, only list the files.  If the
`-v` flag is also given, list the start and end offsets for each file.)  The
codec can be specified with the `-H` flag and is one of <q>ceebot</q>,
<q>ceebot-demo</q>, <q>colobot</q> (the default), <q>colobot-demo</q>, and
<q>none</q> (which does not decode the file).
