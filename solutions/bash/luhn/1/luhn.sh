#!/usr/bin/env bash
main () {
    remove_space=${1//[[:space:]]/}
    [[ $remove_space =~ [^0-9\ ] ]] || [[ ${#remove_space} -lt 2 ]] && echo "false" && exit 
    reversed=$(echo "$remove_space" | rev)
    local new_number=
    for (( i=0; i<"${#reversed}"; i++ )); do
      if [ $i -gt 0 ] && [ "$(echo "($i+1)%2"| bc)" == 0 ]; then
        double=$(( "${reversed:$i:1}" * 2 ))
        [[ $double -gt 9 ]] && double=$(( double -9 ))
        new_number+="$double"
      else
        new_number+="${reversed:$i:1}"
      fi
    done
    sum=
    for (( i=0; i<"${#new_number}"; i++ )); do
      sum=$(( sum + ${new_number:$i:1} ))
    done
    [[ "$(echo "${sum}%10"| bc)" == 0 ]] && echo "true" || echo "false"
}

main "$@"