#!/bin/bash

PgcliConnectLocal(){
  pgcli postgresql://admin@0.0.0.0:5432/rundowns
}

TestDBlisten(){
  sudo nping -c 2 --tcp -p 5432 0.0.0.0
}

"$@"
