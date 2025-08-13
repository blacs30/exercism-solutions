#!/usr/bin/env bash
main() {
    sum=
    for color in $1 $2; do
    case "$color" in
      black)
        sum="${sum}0";;
      brown)
        sum="${sum}1";;
      red)
        sum="${sum}2";;
      orange)
        sum="${sum}3";;
      yellow)
        sum="${sum}4";;
      green)
        sum="${sum}5";;
      blue)
        sum="${sum}6";;
      violet)
        sum="${sum}7";;
      grey)
        sum="${sum}8";;
      white)
        sum="${sum}9";;
      *)
        echo "invalid color"
        exit 1;;
    esac
    done
    printf "%01d" "${sum#0}"
}

main "$@"