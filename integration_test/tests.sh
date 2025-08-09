#!/bin/bash

#first test

bash scripts/echo.sh

DIFF=$(diff results/output_app.txt results/output_bash.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST echo PASSED"
else
    echo "TEST echo FAIL"
fi

#second test
bash scripts/cd_pwd.sh

DIFF=$(diff results/output_app.txt results/output_bash.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST cd and pwd PASSED"
else
    echo "TEST cd and pwd FAIL"
fi

#third test

bash scripts/pipeline.sh

DIFF=$(diff results/output_app.txt results/output_bash.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST pipeline PASSED"
else
    echo "TEST pipeline FAIL"
fi

#fourth test

bash scripts/external.sh

DIFF=$(diff results/output_app.txt results/output_bash.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST external commands PASSED"
else
    echo "TEST external commands FAIL"
fi