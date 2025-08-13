#!/usr/bin/env bash
main () {
    local abc="abcdefghijklmnopqrstuvwxyz"   
    if [ "$1" == "" ]; then echo "false"; exit; fi
    for char in $(echo "$1" | fold -w1); do   
        abc=$(echo ${abc} | sed "s/$char//I")
    done
    if [ $(printf "%s" "$abc" | wc -c) -eq 0 ]; then echo "true"; else echo "false"; fi
}

main "$@"