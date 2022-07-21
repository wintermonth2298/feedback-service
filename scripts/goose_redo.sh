#!/bin/sh

goose -dir ./migrations postgres "user=postgres password=postgres host=postgres dbname=test sslmode=disable" redo
