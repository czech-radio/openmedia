#!/usr/bin/env sh
# Build docker image with non root user. User will have name, ID, GID as user running the docker commands.

SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
DEVEL_DOCKERFILE="${SCRIPT_DIR}/../deploy/Dockerfile_devel"
DEVEL_IMAGENAME="localhost/openmedia_reduce_devel"
# TAG="v0.1.0"

devel_build(){
  if [ ! -f "$DEVEL_DOCKERFILE" ] ; then
    echo "File ${DEVEL_DOCKERFILE} does not exist"
    exit 1
  fi

  DOCKER_BUILDKIT=1 docker build  -f "$DEVEL_DOCKERFILE" -t "$DEVEL_IMAGENAME" \
    --progress=plain \
    --build-arg MY_UID="$(id -u)" \
    --build-arg MY_GID="$(id -g)" \
    --build-arg MY_GROUP="$(id -g -n)" \
    --build-arg MY_USER="$USER" .
}

# devel_run(){
  # docker run -ti $DEVEL_IMAGENAME
# }

devel_run(){
  service="openmedia_reduce_devel"
  docker-compose -f ./deploy/docker-compose_devel.yml run --rm "$service"
}

"$@"
