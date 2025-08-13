#!/usr/bin/env bash
main () {
    bin=$(echo "obase=2; $1" | bc|tr -d "\n"|rev )
    secret=
    for (( i=0; i<${#bin}; i++)); do
      if [ "${bin:i:1}" -eq 1 ]; then
      case "$i" in
        0)
          secret="wink";;
        1)
          secret+=",double blink";;
        2)
          secret+=",close your eyes";;
        3)
          secret+=",jump";;
        esac
      fi
    done
    if [[ ${bin:4:1} -eq 1 ]]; then
      rev=$(echo -n "${secret#,}," | tac -s ",")
      echo ${rev%,}
    else
     s="${secret#,}"
     s="${s%,}"
     echo "$s"
    fi
}
 main "$@"