#!/usr/bin/env bash

str="[[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]],[[5,7,0],[0,3,1],[0,5,0]]"
echo ${str}  | sed -e 's/\[/\{/g' -e 's/\]/\}/g'

