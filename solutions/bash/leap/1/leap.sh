#!/usr/bin/env bash
usage () {
    echo "Usage: leap.sh <year>"
    exit 1
}
main () {
    [[ ! $1 =~ ^[0-9]+$ ]] || [[ $# -ne 1 ]] && usage
    if [[ "$(echo "$1 % 4" | bc)" == 0 ]]; then
      if [[ "$(echo "$1 % 100" | bc)" = 0 ]]; then
        if [[ "$(echo "$1 % 400" | bc)" = 0 ]]; then
          echo "true"
        else
          echo "false"
        fi
      else
        echo "true"
     fi
    else
      echo "false"
    fi
}

main "$@"