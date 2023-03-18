#!/bin/bash

db=$HOME/.local/share/syncsvr/syncsvr.db
sqlite3 $db "INSERT INTO accessTokens(token) values('$1')"
