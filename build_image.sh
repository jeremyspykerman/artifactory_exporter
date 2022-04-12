#/bin/sh

VERSION="0.2.0_sky"
SOURCE_COMMIT=`git rev-parse HEAD`
SOURCE_BRANCH=`git rev-parse --abbrev-ref HEAD`
BUILD_DATE=`date "+%Y%m%d%H%M.%S"`
IMAGE="ghcr.io/sky-uk/gtvd-skyeu-shared-services-artifactory/artifactory_exporter"

docker build -t "${IMAGE}:${VERSION}" --build-arg VERSION=${VERSION} --build-arg SOURCE_COMMIT=${SOURCE_COMMIT} --build-arg SOURCE_BRANCH=${SOURCE_BRANCH} --build-arg BUILD_DATE=\'${BUILD_DATE}\' .
