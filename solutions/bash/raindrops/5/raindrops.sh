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
my_map=( [3]="Pling" [5]="Plang" [7]="Plong" )

main () {
    output=""

    for key in ${!my_map[@]}; do
        if [ "$(($1 % ${key}))" -eq 0 ]; then
          output+="${my_map[${key}]}"
        fi
    done

    echo "${output:-$1}"
}

main "$@"