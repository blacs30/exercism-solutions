#!/usr/bin/env bash
usage () {
    echo "wrong input" 
    exit 1
}
calc() {
   [[ $(echo "($1)^2 + ($2)^2 <= $3^2" | bc ) == 1 ]]
}
main () {
    [[ $# -lt 2 ]] && usage
    [[ $1 =~ [^0-9.-]+ ]] || [[ $2 =~ [^0-9.-]+ ]] && usage
    if calc "$1" "$2" 1; then
      echo 10
    elif calc "$1" "$2" 5; then
      echo 5
    elif calc "$1" "$2" 10; then
      echo 1
    else 
    echo 0
    fi

}
main "$@"