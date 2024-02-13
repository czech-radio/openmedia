#!/bin/bash -u
# -u : fail on unbound variable

GITHUB_API_URL="https://api.github.com/repos"
REPOS_GROUP="${REPOS_GROUP:-czech-radio}"
REPO_NAME="${REPO_NAME:-openmedia-archive}"
BINARY_NAME=${BINARY_NAME:-openmedia-archive}
SERVICE_NAME=${SERVICE_NAME:-openmedia-archive}
MAIN_COMMAND="${MAIN_COMMAND:-openmedia-archive -V}"
REPO_URL="${GITHUB_API_URL}/${REPOS_GROUP}/${REPO_NAME}"
ASSET_DOWNLOAD_URL="https://github.com/${REPOS_GROUP}/${REPO_NAME}/releases/download"
AUTO_UPDATE_SERVICE="${AUTO_UPDATE_SERVICE:-false}"
RELEASE_TAG=${RELEASE_TAG:-latest}

ServiceActivate(){
  systemctl --user enable "${PWD}/${SERVICE_NAME}.service"
  # systemctl --user start "${SERVICE_NAME}.service"
  systemctl --user enable "${PWD}/${SERVICE_NAME}.timer"
  systemctl --user start "${SERVICE_NAME}.timer"
}

ServiceDeactivate(){
  systemctl --user stop "${SERVICE_NAME}.timer"
  systemctl --user disable "${SERVICE_NAME}.timer"
  systemctl --user stop "${SERVICE_NAME}.service"
  systemctl --user disable "${SERVICE_NAME}.service"
}

ServiceStatus(){
  journalctl --user -u "${SERVICE_NAME}" -f
}

GetNewReleaseTag(){
  curl --silent --show-error "${REPO_URL}/releases/latest" | jq -r ".tag_name"
}

NeedsUpdate(){
  local latest_tag="$(GetNewReleaseTag)"
  local latest_date="$(curl --silent --show-error -L "https://github.com/${REPOS_GROUP}/${REPO_NAME}/releases/download/${latest_tag}/version.txt" | head -1)"
  local current_date="$(head -1 version.txt)"
  if [[ "$current_date" != "$latest_date" ]] && "${AUTO_UPDATE_SERVICE}" ; then
    echo true
    return
  fi
  echo false
}

DownloadReleaseFile(){
  local release_tag="$1"
  local filename="$2"
  local filename_url="$ASSET_DOWNLOAD_URL/$release_tag/$filename"
  if ! curl --silent --show-error -L -O "${filename_url}" ; then
    echo "Failed to download release file: $release_tag/$filename" >&2
    rm "$filename"
  fi
  echo "Downloaded release file: $release_tag/$filename"
}

DownloadReleaseFiles(){
  local release_tag="$1"
  declare -a ReleaseFiles=(
    "${BINARY_NAME}"
    "${BINARY_NAME}.service"
    "${BINARY_NAME}.timer"
    "version.txt"
  )
  update="$(NeedsUpdate)"
  for file in "${ReleaseFiles[@]}"; do
    if [[ ! -f "$file" ]] || "$update" ; then
      DownloadReleaseFile "$release_tag" "$file"
    fi
  done
  chmod u+x "./${BINARY_NAME}"
  if "$update" ; then
    systemctl --user daemon-reload
  fi
}

ServiceServe(){
  #TODO: graceful handling of deactivation of running service: when the main command is still running. e.g. through service unit file directives. Trap errors log.
  #
  # systemctl is-active --quiet service
  local release_tag="$RELEASE_TAG"
  if [[ "$release_tag" == "latest" ]] ; then
    release_tag="$(GetNewReleaseTag)"
  fi
  if [[ ! "$release_tag" =~ v.* ]] ; then
    echo "Cannot get tag name. Asset files not downloaded." >&2
    return 1
  fi
  DownloadReleaseFiles "$release_tag"
  
  # Activate service
  service_status="$(systemctl --user is-enabled "$SERVICE_NAME")"
  if [[ "$service_status" != "enabled" ]]; then
    ServiceActivate
    return
  fi
  printenv
  eval "./${MAIN_COMMAND}"
}

ServiceTrigger(){
  systemctl --user start "$SERVICE_NAME"
}

"$@"
