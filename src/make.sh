#!/usr/bin/env bash

base=$1
if [ ! -f ${base}.latex ]; then
	exit 1;	
fi

tmp=$(mktemp --suffix=.latex)
>&2 echo tmpfile is ${tmp}
cat header.latex > ${tmp}
cat $base.latex >> ${tmp}
cat footer.latex >> ${tmp}
pandoc -f latex -t gfm ${tmp} > md/${base}.md

unlink ${tmp}
