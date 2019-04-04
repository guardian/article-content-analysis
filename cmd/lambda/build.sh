#!/usr/bin/env bash

# Script creates the following files
# and then copies the artifact to S3
# so that the lambda can be deployed by RiffRaff.

# packages
# |-- riff-raff.yaml
# |-- build.json
# |-- article-entity-analysis-cfn
# |   |-- cfn.yaml
# |-- article-entity-analysis
# |   |-- lambda.zip

set -ve

# article entity root directory
# Rest of script is executed with this as the working directory.
ROOT_DIR=$(dirname "$0")/../..
cd ${ROOT_DIR}

    PROJ_DIR=cmd/lambda

mkdir -p packages
cp ${PROJ_DIR}/riff-raff.yaml packages

CFN_DIR=packages/article-entity-analysis-cfn
mkdir -p ${CFN_DIR}
cp ${PROJ_DIR}/cfn.yaml ${CFN_DIR}

# Build binary
DOCKER_WD="/root"
docker run --rm \
    -v ${PWD}:${DOCKER_WD} \
    -w ${DOCKER_WD} \
    -e GOOS=linux \
    -e GOARCH=amd64 \
    -e CGO_ENABLED=0 \
    golang:1.11 \
    bash -c "go build ${PROJ_DIR}/main.go"

# Output of docker run command will be an executable file named main. Zip it.
ZIP_FILE=article-entity-analysis.zip
zip ${ZIP_FILE} main

LAMBDA_DIR=packages/article-entity-analysis
mkdir -p ${LAMBDA_DIR}
cp ${ZIP_FILE} ${LAMBDA_DIR}

PROJECT_NAME=article-entity-analysis
BUILD_NUMBER=${BUILD_NUMBER=DEV}
BUILD_NAME=article-entity-analysis-build.${BUILD_NUMBER}
BUILD_START_DATE=$(date +"%Y-%m-%dT%H:%M:%S.000Z")

# Note BUILD_NUMBER and BUILD_VCS_NUMBER are predefined Team City build parameters.
# https://confluence.jetbrains.com/display/TCD9/Predefined+Build+Parameters

cat >build.json << EOF
{
   "projectName":"${PROJECT_NAME}",
   "buildNumber":"${BUILD_NUMBER}",
   "startTime":"${BUILD_START_DATE}",
   "revision":"${BUILD_VCS_NUMBER}",
   "vcsURL":"git@github.com:guardian/article-entity-analysis.git",
   "branch":"${BRANCH_NAME}"
}
EOF

cat build.json

aws s3 cp \
    --acl bucket-owner-full-control \
    --region=eu-west-1 \
    build.json s3://riffraff-builds/${PROJECT_NAME}/${BUILD_NUMBER}/build.json

aws s3 cp \
    --acl bucket-owner-full-control \
    --region=eu-west-1 \
    --recursive \
    packages s3://riffraff-artifact/${PROJECT_NAME}/${BUILD_NUMBER}

# clean up
rm -r packages
