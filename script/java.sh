#!/bin/sh
node script/index.js java
mkdir -p src/main/java/com/boundlessgeo/schema
mv Actions.java src/main/java/com/boundlessgeo/schema/
protoc ./proto/Msg.proto --java_out="./src/main/java"
protoc ./proto/Feature.proto --java_out="./src/main/java"

./java/gradlew -p java build
./java/gradlew -p java install
./java/gradlew -p java uploadArchives

rm -rf src
rm -rf build

