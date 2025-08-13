#!/usr/loca/bin/env bash

CONSTITUTION=

throw_dice () {
    eyes=()
    for _ in {1..4}; do
      number=$(( ( RANDOM % 6 )  + 1 ))
      eyes+=( "$number" )
    done
    echo "${eyes[@]}" | tr " " "\n" | sort -n | tail -n +2 | awk '{s+=$1} END {print s}'
}

get_modifier() {
    subst="$(( CONSTITUTION - 10 ))"
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
        eyes_sum=$(throw_dice)
        echo "$a $eyes_sum"
        if [ "$a" == "constitution" ]; then
          CONSTITUTION="$eyes_sum"
        fi
      done
      modifier=$(get_modifier)
      hitpoints=$(( 10 + modifier ))
      echo "hitpoints $hitpoints"
    fi
}
main "$@"
