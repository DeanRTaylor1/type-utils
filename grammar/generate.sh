#!/bin/sh

alias antlr4='java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -package parser -o ../parser HCLLikeDSL.g4
