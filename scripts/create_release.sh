#!/bin/bash

GITHUB_TOKEN=$1
TAG_NAME="$(cut -d '/' -f 3 <<< $2)" # Get the tag name from "refs/tags/<tag_name>"
DESCRIPTION=$3

GITHUB_REPO="AntonyChR/orus-media-server"

# Create a release

CREATE_RELEASE_URL="https://api.github.com/repos/$GITHUB_REPO/releases"

RESP="$(curl -L \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  "$CREATE_RELEASE_URL"\
  -d "{\"tag_name\":\"$TAG_NAME\",\"target_commitish\":\"main\",\"name\":\"$TAG_NAME\",\"body\":\"$DESCRIPTION\",\"draft\":false,\"prerelease\":false,\"generate_release_notes\":false}")"


RELEASE_ID="$(echo $RESP | grep -o "\"id\":\s*[0-9]*" | grep -o "[0-9]*" | head -1)"

echo "Release ID: $RELEASE_ID"

if [ -z "$RELEASE_ID" ]; then
  echo "Failed to create release"
  exit 1
fi

# Generate zip file
ZIP_FILE="orus-media-server_${TAG_NAME}_amd64.zip"
mv dist/app orus-media-server
zip $ZIP_FILE orus-media-server

# Upload the zip file
UPLOAD_ASSETS_URL="https://uploads.github.com/repos/$GITHUB_REPO/releases/$RELEASE_ID/assets?name=$ZIP_FILE"

curl -L \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  -H "Content-Type: application/octet-stream" \
  "$UPLOAD_ASSETS_URL" \
  --data-binary "@$ZIP_FILE"