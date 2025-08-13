#!/usr/loca/bin/env bash

throw_dice () {
    eyes=()
    for _ in {1..4}; do
      eyes+=( $(( RANDOM % 6 + 1 )) )
    done
    IFS=$'\n' sorted=($(sort -n <<<"${eyes[*]}"))
    unset IFS
    awk 'BEGIN {t=0; for (i=2; i<ARGC; i++) t+=ARGV[i]; print t}' "${sorted[@]}"
}

get_modifier() {
    constitution="$1"
    modifier="$(( constitution - 10 ))"
    if (( modifier < 0 && modifier % 2 != 0 )); then
       echo $(( modifier / 2 - 1 ))
    else
       echo $((  modifier / 2 ))
    fi
}

main() {
    if [ "$1" == "modifier" ]; then
      get_modifier "$2"
    else
      for ability in strength dexterity constitution intelligence wisdom charisma; do
        if [ "$ability" == "constitution" ]; then
          sum=$(throw_dice)
          echo "$ability $sum"
          echo "hitpoints $(( 10 + $(get_modifier "$sum") ))"
        else
          echo "$ability $(throw_dice)"
        fi
      done
    fi
}
main "$@"
