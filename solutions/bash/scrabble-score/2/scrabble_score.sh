#!/usr/bin/env bash
main () {
  sum=0
  word=${1^^}
  for (( char=0; char<${#word}; char++)); do
  case "${word:char:1}" in
    A|E|I|O|U|L|N|R|S|T)
        sum=$(( 1 + sum )) ;;
    D|G)
        sum=$(( 2 + sum )) ;;
    B|C|M|P)
        sum=$(( 3 + sum )) ;;
    F|H|V|W|Y)
        sum=$(( 4 + sum )) ;;
    K)
        sum=$(( 5 + sum )) ;;
    J|X)
        sum=$(( 8 + sum )) ;;
    Q|Z)
        sum=$(( 10 + sum )) ;;

    *) ;;
  esac
  done
  echo "$sum"

}
main "$@"