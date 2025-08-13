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
HITPOINTS=10
CONSTITUTION=

# MODIFIER=constituion - 10 / 2 round down
# HITPOINTS = constituion + modifier
RND_NUMBER='$(( ( RANDOM % 6 )  + 1 ))'
throw_dice () {
    ability="$1"
    eyes=()
    for i in {1..4}; do
      number=$(eval echo $RND_NUMBER)
      eyes+=( "$number" )
    done
    eyes=$(echo "${eyes[@]}" | tr " " "\n" | sort -n | tail -n +2 | awk '{s+=$1} END {print s}')
    echo "$ability $eyes" | tee -a log
    if [ "$ability" == "constitution" ]; then
      CONSTITUTION="$eyes"
    fi
}

get_modifier() {
    subst=$(echo "$(( ($CONSTITUTION - 10) ))")
    if (( subst < 0 && subst % 2 != 0 )); then
       echo $(( subst / 2 - 1 ))
    else
       echo $((  subst / 2 ))
    fi
}

main() {
    if [ "$1" == "modifier" ]; then
      CONSTITUTION="$2"
      get_modifier
    else
      abilities=(strength dexterity constitution intelligence wisdom charisma)
      for a in "${abilities[@]}"; do
        throw_dice "$a"
      done
      modifier=$(get_modifier)
      hitpoints=$(( 10 + modifier ))
      echo "hitpoints $hitpoints" | tee -a log
    fi
}
touch log
echo "$@" >> log
main "$@"
