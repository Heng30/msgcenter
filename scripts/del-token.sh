#!/bin/bash

db=$HOME/.local/share/syncsvr/syncsvr.db
sqlite3 $db "DELETE FROM accessTokens where token='$1'"
