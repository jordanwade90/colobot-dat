colobot-dat: read game data files from Colobot
==============================================

**This code is obsolete!** The [open-source Colobot project][colobot] hasn't
used the DAT container format since [colobot/colobot@6cce7ec6][commit]; this
code is of use to the historically curious only and shouldn't actually be used
for anything. It especially shouldn't be used for anyone wanting a container
file format, because this format is a terrible one; use tarballs or zip files
instead.

[colobot]: https://github.com/colobot/colobot
[commit]: https://github.com/colobot/colobot/commit/6cce7ec6fd262178ce91d218f9287363842349cd

Release builds of Ceebot and Colobot store their game files in DAT containers.
This project consists of (1) a library for reading DAT files and (2) a
command-line utility to extract files from a DAT container.

Library
=======

Read a container file by passing an `io.ReaderAt` and a `Codec` to `New()`.  A
`Codec` XORs data in the container with a predefined table; `Codec`s for the
demo and full versions of Ceebot and Colobot are included with the library.
(The tables were extracted from the [src/metafile.cpp][metafile] of the
original Colobot source.)  The resulting `Container` contains a slice of
`File`s, each of which has an `io.SectionReader` which will read the data from
that file.

[metafile]: https://github.com/colobot/colobot/blob/colobot-original/src/metafile.cpp

Command-line extractor
======================

`colobot-dat` is an extractor for DAT files.  Given a path to a DAT file, it
will extract each file into the current directory, unless a file with the same
name already exists.  (If the `-l` flag is given, only list the files.  If the
`-v` flag is also given, list the start and end offsets for each file.)  The
codec can be specified with the `-H` flag and is one of <q>ceebot</q>,
<q>ceebot-demo</q>, <q>colobot</q> (the default), <q>colobot-demo</q>, and
<q>none</q> (which does not decode the file).
