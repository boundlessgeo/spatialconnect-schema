#!/bin/bash
# script dir
if [ `uname` == "Linux" ]; then
    SCRIPT_DIR=$(dirname $(readlink -f $0))
else
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
fi
# Project directory is parent of script directory
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"
# source generation script is in this script's directory
GEN_SRC_SCRIPT="${SCRIPT_DIR}/index.sh"
# Java project root is in PROJECT_DIR/java
JAVA_ROOT_DIR="${PROJECT_DIR}/java"
# Java generated source directory is "src/main/java/com/boundlessgeo/schema" under JAVA_ROOT_DIR
JAVA_SRC_DIR="${JAVA_ROOT_DIR}/src/main/java/com/boundlessgeo/schema"
# Protobuf source directory is "src/main/java" under JAVA_ROOT_DIR
PROTOBUF_SRC_DIR="${JAVA_ROOT_DIR}/src/main/java"
# ensure the source directory exists
mkdir -p ${JAVA_SRC_DIR}
# remove any old source code laying around that may not get overwritten
if [ -d ${JAVA_ROOT_DIR}/src -a -w ${JAVA_ROOT_DIR}/src ]; then
  find ${JAVA_ROOT_DIR}/src -type f -exec rm {} \;
fi
# Protobuf root
PKG_DIR="${PROJECT_DIR}/proto/boundlessgeo_schema"
# source directory for proto
PROTO_DIR="${PROJECT_DIR}/proto/"
# generate java source
${GEN_SRC_SCRIPT} java ${JAVA_SRC_DIR}
# generate GRPC -- not now
#protoc "${PKG_DIR}/worm.proto" "${PKG_DIR}/Command.proto" "${PKG_DIR}/Event.proto" "${PKG_DIR}/Metadata.proto" \ 
#--proto_path="${PROTO_DIR}" --java_out=plugins=grpc:"${PROTOBUF_SRC_DIR}"
# generate protobuf source
protoc "${PKG_DIR}/Metadata.proto" --proto_path="${PROTO_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
protoc "${PKG_DIR}/Command.proto" "${PKG_DIR}/Metadata.proto" --proto_path="${PROTO_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
protoc "${PKG_DIR}/Event.proto" "${PKG_DIR}/Metadata.proto" --proto_path="${PROTO_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
protoc "${PKG_DIR}/Msg.proto" --proto_path="${PROTO_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
protoc "${PKG_DIR}/Feature.proto" --proto_path="${PROTO_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
# build Java artifacts
# generate the pom from the version.json
# get the version from version.json
VERSION="$(grep version "${SCRIPT_DIR}"/version.json | awk '{print $2}'| sed -e 's/\"//g')"
# fill in pom.xml with version
cat "${JAVA_ROOT_DIR}/pom.template" | sed -e 's/\${from\.script}/'${VERSION}'/' > "${JAVA_ROOT_DIR}/pom.xml"
mvn clean install -f "${JAVA_ROOT_DIR}/pom.xml"
