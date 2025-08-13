#!/usr/bin/env bash
exponent () {
  echo "2^$(($1 - 1 ))" | bc
}

main () {
    # echo "$1"
    if [ "$1" = "total" ]; then
      sum=0
      for (( i=1; i<65; i++)); do
        calc=$(exponent "$i")
        sum=$(echo "$sum + $calc" | bc)
      done
      echo "$sum"
    elif [ "$1" -lt 1 ] || [ "$1" -gt 64 ]; then
      echo "Error: invalid input"
      exit 1
    else
      exponent "$1"
    fi

}

main "$@"