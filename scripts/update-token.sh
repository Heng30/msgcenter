#!/bin/bash

db=$HOME/.local/share/syncsvr/syncsvr.db
sqlite3 $db "UPDATE accessTokens set token='$2' where token='$1'"
