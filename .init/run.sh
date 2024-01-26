#!/bin/bash -u
# -u : fail on unbound variable

GITHUB_API_URL="https://api.github.com/repos"
REPOS_GROUP="${REPOS_GROUP:-czech-radio}"
REPO_NAME="${REPO_NAME:-openmedia-archive}"
BINARY_NAME=${BINARY_NAME:-openmedia-archive}
SERVICE_NAME=${SERVICE_NAME:-openmedia-archive}
MAIN_COMMAND="openmedia-archive -V"
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
  curl --silent "${REPO_URL}/releases/latest" | jq -r ".tag_name"
}

DownloadAsset(){
  local tag="$1"
  local asset="$2"
  local assets_url="${ASSET_DOWNLOAD_URL}/${tag}"
  if ! curl --silent -L -O "${assets_url}/${asset}" ; then
    # --clobber: overwrite destination files
    echo "Failed to download new version assets: $tag" >&2
    return 1
  fi
  echo "Download new version asset: ${tag}/${asset}"
}

DownloadTagReleaseFiles(){
  local tag="$1"
  local assets_url="${ASSET_DOWNLOAD_URL}/${tag}"  
  echo "Downloading assets"
  DownloadAsset "$tag" "${BINARY_NAME}"
  DownloadAsset "$tag" "${BINARY_NAME}.service"
  DownloadAsset "$tag" "${BINARY_NAME}.timer"
  DownloadAsset "$tag" "version.txt"
  # DownloadAsset "$tag" "run.sh"
  # chmod u+x "./run.sh"
  chmod u+x "./${BINARY_NAME}"
}

ServiceUpdate(){
  #TODO: graceful handling of deactivation of running service: when the main command is still running. e.g. through service unit file directives. Trap errors log.
  local tag="$1"
  if [[ "$AUTO_UPDATE_SERVICE" != "true" ]]; then
    return
  fi
  # local latest_tag="$(curl -s -L https://github.com/czech-radio/openmedia-archive/releases/download/v0.9.0/version.txt | yq '.tag')"
  # local current_tag=$(cat version.txt | yq '.tag')
  local latest_date="$(curl -s -L https://github.com/czech-radio/openmedia-archive/releases/download/v0.9.0/version.txt | yq '.date')"
  local current_date=$(cat version.txt | yq '.date')
 
  # if [[ "$current_tag" != "$latest_tag" ]]; then
  if [[ "$current_date" != "$latest_date" ]]; then
    echo Updating service assets
    DownloadTagReleaseFiles "$latest_tag"
    systemctl --user daemon-reload
  fi
}

ServiceServe(){
  local tag="$RELEASE_TAG"
  if [[ "$tag" == "latest" ]] ; then
    tag="$(GetNewReleaseTag)"
  fi
  if [[ ! "$tag" =~ v.* ]] ; then
    echo "Cannot get tag name. Asset files not downloaded." >&2
    return 1
  fi
  
  # Check if assets are present
  if [[ ! -f "version.txt" ]]; then
    DownloadTagReleaseFiles "$tag"
  fi
  ServiceUpdate "$tag"
  
  # Activate service
  service_status="$(systemctl --user is-enabled "$SERVICE_NAME")"
  if [[ "$SERVICE_NAME" != "enabled" ]]; then
    ServiceActivate
    return
  fi
  ./${MAIN_COMMAND} 
}

ServiceTrigger(){
  systemctl --user start "$SERVICE_NAME"
}

"$@"
