#!/usr/bin/env bash
main () {
    length=${#1}
    reversed=""
    for (( i=length; i>=0; i-- )); do
      reversed+="${1:i:1}"
    done
    echo "$reversed"
}

main "$@"