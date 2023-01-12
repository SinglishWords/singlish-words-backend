#!/bin/bash

echo $(date +"%T") >> ~/detect-categories.log
cd /root/deployment/backend/src/scripts
./scripts detect-categories | head >> ~/detect-categories.log
