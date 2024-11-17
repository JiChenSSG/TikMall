#!/bin/bash

. scripts/list_app.sh

get_app_list

readonly root_path=`pwd`
for app_path in ${app_list[*]}; do
  if [[ "${app_path}" = "app/common" ]]; then
    continue
  fi
  echo "copy ${app_path} log file"
  ln -s "../${app_path}/log/kitex.log" "log/${app_path##*/}.log"
  echo "Done!"
done