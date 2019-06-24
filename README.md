# Swarm Node implementer spec

Documentation on how to create a custom Swarm node implementation

[https://github.com/ethersphere/user-stories/issues/50](https://github.com/ethersphere/user-stories/issues/50)

## Contents

The documents are in latex format, and are found in `./src`

Tools used for message serializations used in the documents can be found in `./tools`. They are written in `golang` and use libraries from the official `geth` and `swarm` implementations.

## Build

To build pdf with bibliography:

```
cd $REPO/src
pdflatex spec.latex
bibtex spec
pdflatex spec.latex
pdflatex spec.latex
```
