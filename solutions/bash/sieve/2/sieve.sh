#!/usr/bin/env bash
calcMultiples () {
    local start=$1
    local end=$2
    local marked="$3"
    local list=()
    for (( i=(start+start);i<="$end";i+="$start" )); do
      if ! [[ " ${marked[*]} " =~ [[:space:]]${i}[[:space:]] ]]; then
        list+=("$i")
      fi
    done
    echo "${list[@]}"
}
main () {
    [[ $# -lt 1 ]] || [[ $1 = "" ]] || [[ $1 =~ ^[^0-9]+$ ]] || [ "$1" -lt 2 ] && exit 
    marked=()
    local prime=()
    for ((i=2;i<=$1;i++)); do
      if [[ " ${marked[*]} " =~ [[:space:]]${i}[[:space:]] ]]; then
        continue; 
      fi
        prime+=("$i")
        multiples=$(calcMultiples "$i" "$1" "${marked[@]}")
        marked+=("$multiples")
    done
    echo "${prime[@]}"
}
main "$@"