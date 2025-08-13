#!/usr/bin/env bash
main () {
    if [ $# -lt 2 ]; then echo "false" && exit 1; fi
    case "$1" in 
      "encode")
        encode "$2";;
      "decode")
        decode "$2";;
      *)
        echo "wrong action"; exit 1;;
    esac
}
flip_char() {
    declare -A abc
    abc=( 
    [1]="a" 
    [2]="b" 
    [3]="c" 
    [4]="d" 
    [5]="e"
    [6]="f" 
    [7]="g" 
    [8]="h" 
    [9]="i" 
    [10]="j" 
    [11]="k" 
    [12]="l" 
    [13]="m" 
    [14]="n" 
    [15]="o" 
    [16]="p" 
    [17]="q" 
    [18]="r" 
    [19]="s" 
    [20]="t" 
    [21]="u" 
    [22]="v" 
    [23]="w" 
    [24]="x" 
    [25]="y" 
    [26]="z"
    )
    char="$1"
    pos=

    for key in "${!abc[@]}"; do
      if [ "${abc[$key]}" = "$char" ]; then
        pos=$key
        break
      fi
    done
    new_pos=$(( (pos -1 - 26) * -1 ))
    echo "${abc[$new_pos]}"
}

encode () {
    # echo "Received $input"
    sanitzed=${1//[[:space:]|[:punct:]]/}
    sanitzed=${sanitzed,,}
    # echo "Sanitzed $sanitzed"
    new_string=
    for (( i=0; i<${#sanitzed}; i++)); do
        char="${sanitzed:$i:1}" 
        if [ $(( "$i" % 5)) -eq 0 ] && [ $i -gt 0 ]; then
          new_string+=" "
        fi
        [[ $char =~ [[:alpha:]] ]] && new_string+=$(flip_char "$char") || new_string+="$char"
    done
    echo "$new_string"

}

decode () {
    sanitzed=${1//[[:space:]|[:punct:]]/}
    sanitzed=${sanitzed,,}
    new_string=
    for (( i=0; i<${#sanitzed}; i++)); do
        char="${sanitzed:$i:1}" 
        [[ $char =~ [[:alpha:]] ]] && new_string+=$(flip_char "$char") || new_string+="$char"
    done
    echo "$new_string"
}
main "$@"