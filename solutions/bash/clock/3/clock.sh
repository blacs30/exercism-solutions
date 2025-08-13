#!/usr/bin/env bash
global_hour=$1
global_min=$2
global_op=$3
global_mod=$4

debug() {
    echo "hour $global_hour"
    echo "min $global_min"
    echo "op $global_op"
    echo "mod $global_mod"
}

remove_sign() {
    if [ "$1" -lt 0 ]; then
      echo $(( $1 * -1 )) 
    else
      echo "$1"
    fi
}

validate() {
    input=$1
    regexp=$2
    [[ "$input" =~ ^${regexp}$ ]] && return || echo "invalid arguments $input"; debug; exit 1
}

main () {
  local hour=$1
  local min=$2
  local op=$3
  local mod=$4

  if [ "$#" -lt 2 ] || [ "$#" -gt 2 ] && [ "$#" -ne 4 ]; then
    echo "invalid arguments"
    exit 1
  fi

  validate "$hour" "[-0-9]+"
  validate "$min" "[-0-9]+"

  if [ "$#" -eq 4 ];then 
    validate "$mod" "[-0-9]+"
    validate "$op" "[-+]"
    [[ "$op" = "-" ]] && min=$(( min - mod )) || min=$(( min + mod ))
  fi

  local new_hour=$((( 24 + (hour % 24)) %24 ))
  local new_min=$((( 60 + (min % 60)) %60 ))
  
  if [ "$min" -lt 0 ]; then 
    # need to round negative hours up because we are still in that running hour
    add_hours=$(( (min - (60-1)) / 60 ))
  else
      add_hours=$(( min / 60 ))
  fi

  if [ $add_hours -lt 0 ]; then
    new_hour=$((  (( $(remove_sign $add_hours) + (24-1))  / 24 * 24 + new_hour - $(remove_sign $add_hours)) % 24 ))
  else
    new_hour=$(( (new_hour + add_hours) % 24 ))
  fi
  printf "%02d:%02d\n" $(( new_hour )) $(( new_min ))
}

main "$@"