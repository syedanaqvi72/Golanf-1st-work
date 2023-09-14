/*API stands for Application Programming Interface. The most important letter is I. An API is an interface.

An interface is like a frontier between two things, a “shared boundary across which two or more separate components of a computer system exchange information” (Wikipedia).

With this definition we can say that the go package fmt exposes an API to the go programmer to interact with it. Its API represents the set of functions that you can use for instance Println. It also covers the set of exported identifiers of the package (constants, variables, types).

Go Modules expose an API that is composed of all exported identifiers of the package(s) that the module is composed of.



Vocabulary: tag, revision, prerelease, release 
Tagged : it means that a “tag” has been created with a Version Control system
A tag is a string (it’s name ) that designate a specific revision of a project maintained by a Version Control System.
A tag can designate a released version or a pre-released version of the software

A pre-release version (or release candidate) is considered to be ready. It is made available for last minutes tests.

Yet, it is not considered as a stable release (or just “release”).

Pre-release versions have specific tags with appended characters :

For instance : 1.0.0-alpha, 1.0.0-alpha.1

Each project has different conventions about that

An untagged version is a specific state of the program at a given time.

In the Git VCS, it’s a “commit”