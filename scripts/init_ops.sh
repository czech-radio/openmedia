#!/usr/bin/env bash
MY_SERVICE=openmedia_reduce
INIT_DIR="../init"

enable_service(){
  systemctl --user enable ${INIT_DIR}/${MY_SERVICE}.service
}

enable_service_timer(){
  systemctl --user enable ${INIT_DIR}/${MY_SERVICE}.timer
  systemctl --user start ${MY_SERVICE}.timer
}

disable_service(){
  systemctl --user disable ${MY_SERVICE}.timer
  systemctl --user disable ${MY_SERVICE}.service
  systemctl --user stop ${MY_SERVICE}.service
}
