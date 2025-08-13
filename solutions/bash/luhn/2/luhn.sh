#!/usr/bin/env bash
main () {
    remove_space=${1//[[:space:]]/}
    [[ $remove_space =~ [^0-9\ ] ]] || [[ ${#remove_space} -lt 2 ]] && echo "false" && exit 
    reversed=$(echo "$remove_space" | rev)
    local sum=
    for (( i=0; i<"${#reversed}"; i++ )); do
      if [ $i -gt 0 ] && [ "$(echo "($i+1)%2"| bc)" == 0 ]; then
        double=$(( "${reversed:$i:1}" * 2 ))
        [[ $double -gt 9 ]] && double=$(( double -9 ))
        sum=$(( sum + double ))
      else
        sum=$((sum + ${reversed:$i:1} ))
      fi
    done
    [[ "$(echo "${sum}%10"| bc)" == 0 ]] && echo "true" || echo "false"
}

main "$@"