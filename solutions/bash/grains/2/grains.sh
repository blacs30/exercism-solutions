#!/usr/bin/env bash
TOTAL_NUMBER=18446744073709551615
exponent () {
  echo "2^$(($1 - 1 ))" | bc
}
main () {
    if [ "$1" = "total" ]; then
      # This number won't change so we can hard code it
      echo "$TOTAL_NUMBER"
    elif [ "$1" -lt 1 ] || [ "$1" -gt 64 ]; then
      echo "Error: invalid input"
      exit 1
    else
      exponent "$1"
    fi

}

main "$@"