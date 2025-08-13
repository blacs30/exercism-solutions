#!/usr/bin/env bash
main () {
    len="${#1}"
    sum=0
    for (( i=0; i<len; i++ )); do
      digit=${1:i:1}
      sum=$(( (digit ** len) + sum ))
    done
    [[ sum -eq $1 ]] && echo "true" || echo "false"
}

main "$@"