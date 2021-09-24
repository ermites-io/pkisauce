#!/bin/sh

#
# lame but to not worry about the templates generation.
# this generate the template files for hasher, marshal & error associated.
#
# shall we generate 3 templates or only one files by parsing reputting together
# not sure, i ll do the least surprising things because it's LESS error prone
# generate 3 files.
#
FILES="error.go marshal.go hash.go"
SED=sed
CAT=cat
PACKAGE=$1

for i in $FILES 
do 
#	echo ${PWD}
	echo generating "../templates/helpers_scramble_$i.template"
	$CAT $i | $SED s/"package scramble"/"package main"/g > ${PWD}/../templates/helpers_scramble_$i.template
done
#cat error.go  | sed s/"package scramble"/"package main"/g
