#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username bosca --dbname bosca <<-EOSQL
	CREATE DATABASE spicedb owner bosca;
EOSQL