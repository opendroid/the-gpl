#/usr/bin/env zsh
# Example: https://iridakos.com/programming/2018/03/01/bash-programmable-completion-tutorial
thegpl_completions() {
  if [ "${COMP_WORDS[1]}" == "bits" ]; then
    COMPREPLY=($(compgen -W "-n=0x0badcafe" -- "bits"))
    return
  fi
  if [ "${COMP_WORDS[1]}" == "mas" ]; then
    COMPREPLY=($(compgen -W "-fn=array -fn=comp" -- "mas"))
    return
  fi
  if [ "${COMP_WORDS[1]}" == "temp" ]; then
    COMPREPLY=($(compgen -W "-c=12 -f=12 -k=12" -- "temp"))
    return
  fi
  if [ "${COMP_WORDS[1]}" == "degrees" ]; then
    COMPREPLY=($(compgen -W "-c=12°F -f=12°K -k=12°C" -- "degrees"))
    return
  fi
  if [ "${COMP_WORDS[1]}" == "du" ]; then
    COMPREPLY=($(compgen -W "-dir=." -- "du"))
    return
  fi
  if [ "${COMP_WORDS[1]}" == "parse" ]; then
    if [ "$COMP_CWORD" -eq 3 ]; then
      COMPREPLY=($(compgen -W "-type=outline type=links -type=images type=pretty -type=crawl" -- "parse"))
      return
    fi
    COMPREPLY=($(compgen -W "-site=https://www.google.com -site=https://www.duckduckgo.com" -- "site"))
    return
  fi
  if [ "${COMP_WORDS[1]}" == "service" ]; then
    COMPREPLY=($(compgen -W "-sp=clock:9999 -cp=clock:9999 -sp=reverb:9998 -cp=reverb:9998 -sp=chat:9997" -- "server"))
    return
  fi
    if [ "${COMP_WORDS[1]}" == "server" ]; then
    COMPREPLY=($(compgen -W "-port=8080" -- "service"))
    return
  fi
  COMPREPLY=($(compgen -W "bits mas temp du lissajous parse service client server" "${COMP_WORDS[1]}"))
}

complete -F thegpl_completions the-gpl
