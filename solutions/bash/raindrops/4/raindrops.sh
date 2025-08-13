#!/usr/bin/env bash
main () {
    local keys=(3 5 7)
    local sounds=("Pling" "Plang" "Plong")

    local output=""
    for key in "${!keys[@]}"; do
        (( $1 % ${keys[${key}]} == 0 )) && output+="${sounds[${key}]}"
    done

    echo "${output:-$1}"
}

main "$@"