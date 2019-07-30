#!/bin/bash
path=`pwd` 
goGetFile="/interfaces/files/"
goAPI="/interfaces/"
goEmailSender="/services/MailSender/"

cd $path$goGetFile
if [ ! -e "./files" ]; then
  echo "Build files..."
  go build ./
  echo "Running files"
  ./files &
else
  echo "Running files"
  ./files &
fi

cd $path$goAPI
if [ ! -e "./interfaces" ]; then
  echo "Build API interfaces..."
  go build ./
  echo "Running API interfaces on port :8921"
  ./interfaces &
else
  echo "Running API interfaces on port :8921"
  ./interfaces &
fi
cd $path$goEmailSender
if [ ! -e "./MailSender" ]; then 
  echo "Build Email Sender..."
  go build ./
  echo "Running Email Sender on :8225"
  ./MailSender &
else
  echo "Running Email Sender on :8225"
  ./MailSender &
fi
