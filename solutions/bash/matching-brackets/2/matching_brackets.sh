#!/usr/local/bin/env bash
ORDER_LIST=
fail() {
        echo "false"
        exit
}
remove() {
        if [[ "${ORDER_LIST:((${#ORDER_LIST} - 1)):1}" == "$1" ]]; then
                ORDER_LIST="${ORDER_LIST:0:-1}"
        else
                fail
        fi
}

add() {
        ORDER_LIST="${ORDER_LIST}${1}"
}

main() {
        for ((i = 0; i < ${#1}; i++)); do
                case "${1:$i:1}" in
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
                echo "false"
        else
                echo "true"
        fi
}
main "$@"
