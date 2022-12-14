#!/usr/bin/env bash

# TODO: Makefile? Oder kann man latexmk das beibringen?

latexmk
cp build/*.pdf ../pdfs
