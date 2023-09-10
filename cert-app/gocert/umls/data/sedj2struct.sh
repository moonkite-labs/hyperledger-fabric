#!/bin/bash

if [[ "$1" == "UML" ]]; then 
    grep -e 'private ' 'test.txt' | sed -E 's/.+private (.+) ([a-zA-Z]+)\;?[[:space:]]*/    - \u\2: \u\1/gm';
elif [[ "$1" == "GO" ]]; then
    grep -e 'private ' 'test.txt' | sed -E 's/.+private (.+) ([a-zA-Z]+)\;?[[:space:]]*/    \u\2 \L\1/gm';
elif [[ "$1" == "PARAM" ]]; then # Construct parameter list from model struct
    cat test.txt | tr -s ' ' | cut -d ' ' -f 1,2 | sed -E 's/ *([[:alpha:]]+) *([a-zA-Z0-9\.]+)/\l\1 \2/gm' | tr '\n' ', ' | tr '\t' ' ';
elif [[ "$1" == "CONSTRUCT" ]]; then # Construct constructor input from model struct 
    cat test.txt | tr -s ' ' | cut -d ' ' -f 1,2 | sed -E 's/ *([[:alpha:]]+) *([a-zA-Z0-9\.]+)/\1: \l\1,/gm';
elif [[ "$1" == "OPTIONS" ]]; then # Construct struct field for options
    if [[  -z "$2" ]]; then
        echo "Please provide the structure name"
        exit 1
    fi
    cat test.txt | tr -s ' ' | cut -d ' ' -f 1,2 | tr '\t' ' ' | sed -E "s/ *([[:alpha:]]+) *([a-zA-Z0-9\.]+)/func With\u\1 \(\l\1 \l\2) ${2}Option \{\n\treturn func\(i \*models.${2}\) \{\n\t\ti\.\u\1 = \l\1\n\t\}\n\}/gm";
fi