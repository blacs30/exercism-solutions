#!/usr/bin/env bash

main() {
    colors=("black" "brown" "red" "orange" "yellow" "green" "blue" "violet" "grey" "white")
    if [ "$1" = "colors" ]; then
      for color in "${colors[@]}"; do
        echo "$color"
      done
      exit
    fi
    sum=
    case "$2" in
      black)   sum="${sum}0";;
      brown)        sum="${sum}1";;
      red)        sum="${sum}2";;
      orange)        sum="${sum}3";;
      yellow)        sum="${sum}4";;
      green)        sum="${sum}5";;
      blue)        sum="${sum}6";;
      violet)        sum="${sum}7";;
      grey)        sum="${sum}8";;
      white)        sum="${sum}9";;
      colors)        sum="${sum}10";;
      *)
        echo "invalid color"
    esac
    printf "%01d" "${sum#0}"
}
main "$@"