#!/bin/bash
set -e

echo "stdout: message"
>&2 echo -e "stderr: error"
