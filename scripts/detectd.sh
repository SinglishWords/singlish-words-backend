#!/bin/bash

echo $(date +"%T") >> detect-categories.log
./scripts detect-categories | head >> detect-categories.log
