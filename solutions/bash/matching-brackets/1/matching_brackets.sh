#!/usr/local/bin/env bash
ORDER_LIST=
expect() {
    list="$1"
    require="$2"
    len_list="${#list}"
    ((--len_list))
    if [ "${list:$len_list:1}" = "$require" ]; then
      echo 0
    else
      echo 1
    fi
}
remove () {
    char="$1"
    if [[ "$(expect "$ORDER_LIST" "${char}")" -eq "0" ]]; then
            ORDER_LIST="${ORDER_LIST:0:-1}"
          else
            echo "false"
            exit
          fi
}

add () {
    char="$1"
    ORDER_LIST="${ORDER_LIST}${char}"
}

main () {
    for ((i=0; i< ${#1}; i++)); do
      char="${1:$i:1}"
      case "$char" in
        "[") add "[" ;;
        "{") add "{" ;;
        "(") add "(" ;;
        "]") remove "[" ;;
        "}") remove "{" ;;
        ")") remove "(" ;;
        *) ;;
      esac
    done
    if [[ "${#ORDER_LIST}" -ne 0 ]]; then
      echo "false";
    else
      echo "true";
    fi
}

main "$@"
