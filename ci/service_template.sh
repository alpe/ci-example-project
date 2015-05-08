#!/usr/bin/env bash
BUILD_NUMBER=$1
cat >> ci-example-$BUILD_NUMBER@.service <<EOF
[Unit]
Description=ci-example-service

[Service]
User=core
ExecStartPre=-/usr/bin/docker kill ci-example-%i
ExecStartPre=-/usr/bin/docker rm ci-example-%i
ExecStart=/usr/bin/docker run --rm --name ci-example-%i -p 8090:8080 alpetest/ci-example-project:$BUILD_NUMBER

ExecStop=/usr/bin/docker stop ci-example-%i

[X-Fleet]
Conflicts=ci-example-http@*.service
EOF
