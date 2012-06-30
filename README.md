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
