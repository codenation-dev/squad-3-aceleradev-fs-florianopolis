#!/bin/bash
path=`pwd` 
goGetFile="/interfaces/files/"
goAPI="/interfaces/"
goEmailSender="/services/MailSender/"
reactpath="/frontend/"
uiAddUser="/adduser/"
killall MailSender
killall interfaces
killall files
killall adduser
case $1 in
   "-d") 
        cd $path$uiAddUser
        if [ -e "./adduser" ]; then
            echo "Removed adduser"
            rm ./adduser
        fi
        cd $path$goGetFile
        if [ -e "./files" ]; then
            echo "Removed files..."
            rm ./files
        fi
        cd $path$goAPI
        if [ -e "./interfaces" ]; then
            echo "Removed API interfaces..."
            rm ./interfaces
        fi
        cd $path$goEmailSender
        if [ -e "./MailSender" ]; then 
            echo "Removed Email Sender..."
            rm ./MailSender
        fi
         ;;
   *) echo ""
      exit 1
      ;;
esac