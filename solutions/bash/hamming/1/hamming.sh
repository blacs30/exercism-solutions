#!/usr/bin/env bash
usage () {
    echo "Usage: hamming.sh <string1> <string2>"
    exit 1
}
compare_str () {
    left="$1"
    right="$2"
    diffs=0
    for (( i=0; "$i" <${#left}; i++  )); do
      if [ "${left:$i:1}" != "${right:$i:1}" ]; then
        (( diffs++ ))
      fi
    done
    echo $diffs
}
main () {
    if [ $# -lt  2 ]; then
      usage
    elif  [ ${#1} -ne ${#2} ]; then
      echo "strands must be of equal length"
      exit 1
    elif [ "$1" = "$2" ]; then
      echo 0
      exit
    fi
    compare_str "$1" "$2"
}

main "$@"