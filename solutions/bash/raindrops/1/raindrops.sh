#!/usr/bin/env bash

# The following comments should help you get started:
# - Bash is flexible. You may use functions or write a "raw" script.
#
# - Complex code can be made easier to read by breaking it up
#   into functions, however this is sometimes overkill in bash.
#
# - You can find links about good style and other resources
#   for Bash in './README.md'. It came with this exercise.
#
#   Example:
#   # other functions here
#   # ...
#   # ...
#
#   main () {
#     # your main function code here
#   }
#
#   # call main with all of the positional arguments
#   main "$@"
#
# *** PLEASE REMOVE THESE COMMENTS BEFORE SUBMITTING YOUR SOLUTION ***

main () {
    hit=false
    if [ "$(($1 % 3))" -eq 0  ]; then
        printf "Pling"
        hit=true
    fi

    if [ "$(($1 % 5))" -eq 0  ]; then
        printf "Plang"
        hit=true
    fi

    if [ "$(($1 % 7))" -eq 0  ]; then
        printf "Plong"
        hit=true
    fi

    if [ "$hit" != "true" ]; then
        echo $1
    fi
    printf "\n"
}

main "$@"