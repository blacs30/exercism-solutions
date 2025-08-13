#!/usr/bin/env bash
main() {
    local sentence="${1//[-_\\*]/ }"
    acronym=
    for word in ${sentence}; do
      acronym+=$(echo "$word" | fold -w 1 | head -n 1)
    done

    echo "${acronym^^}"
}

main "$@"