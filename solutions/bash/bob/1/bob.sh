#!/usr/bin/env bash
main () {
    # echo "$1"
    if [[ $1 =~ ^[[:blank:]|[:space:]]+$ ]] || [ "$1" = "" ]; then
      echo "Fine. Be that way!"
    elif [[ $1 =~ ^[1-9|A-Z|\*|\,|\!|\%|\^|\@|\#|\$\(|[:space:]]+[A-Z\!]+$ ]]; then
      echo "Whoa, chill out!"
    elif [[ $1 =~ ^[A-Z|1-9|[:blank:]|[:space:]|\']+[A-Z]+\?+$ ]]; then
      echo "Calm down, I know what I'm doing!"
    elif [[ $1 =~ [1-9|a-z|A-Z|[:space:]]+\?+[[:space:]]*$ ]]; then
      echo "Sure."
    else
      echo "Whatever."
    fi
}
main "$@"