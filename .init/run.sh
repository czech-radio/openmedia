#!/bin/bash -u
# -u : fail on unbound variable

GITHUB_API_URL="https://api.github.com/repos"
REPOS_GROUP="${REPOS_GROUP:-czech-radio}"
REPO_NAME="${REPO_NAME:-openmedia-archive}"
BINARY_NAME=${BINARY_NAME:-openmedia-archive}
MAIN_COMMAND="openmedia-archive -V"
REPO_URL="${GITHUB_API_URL}/${REPOS_GROUP}/${REPO_NAME}"
ASSET_DOWNLOAD_URL="https://github.com/${REPOS_GROUP}/${REPO_NAME}/releases/download"
SERVICE_NAME="$(basename "$PWD")"
AUTO_UPDATE_SERVICE="${AUTO_UPDATE_SERVICE:-false}"
RELEASE_TAG=${RELEASE_TAG:-latest}

ServiceActivate(){
  systemctl --user enable "${PWD}/${SERVICE_NAME}.service"
  systemctl --user start "${SERVICE_NAME}.service"
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
  curl -s "${REPO_URL}/releases/latest" | jq -r ".tag_name"
}


DownloadAsset(){
  local tag="$1"
  local asset="$2"
  local assets_url="${ASSET_DOWNLOAD_URL}/${tag}"
  if ! curl -s -L -O "${assets_url}/${asset}" ; then
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
  DownloadAsset "$tag" "run.sh"
}


ServiceRun(){
  local tag="$RELEASE_TAG"
  if [[ "$tag" == "latest" ]] ; then
    tag="$(GetNewReleaseTag)"
  fi
  if [[ ! "$tag" =~ v.* ]] ; then
    echo "Cannot get tag name. Asset files not downloaded." >&2
    return 1
  fi
  if [[ ! -f "$BINARY_NAME" ]]; then
    DownloadTagReleaseFiles "$tag"
  fi
  # Check if binary is present
  if [[ "$AUTO_UPDATE_SERVICE" == "true" ]]; then
    #TODO: gracefull handling of deactivation of runnig service: when the main command is still running. e.g. through service unit file directives
    DownloadTagReleaseFiles "$tag"
    systemctl --user daemon-reload
  fi
  
  # Run main command
  chmod u+x ./"$BINARY_NAME"
  ./${MAIN_COMMAND}
}

"$@"
