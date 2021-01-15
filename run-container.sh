#!/bin/bash


docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/mathapp mathapp
