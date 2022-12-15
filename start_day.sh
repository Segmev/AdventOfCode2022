#!/bin/bash

mkdir day${1}
cp ./template/main.go day${1}/day${1}.go
cp ./template/input.txt day${1}/input.txt
sed -i "s/package main/package day${1}/" day${1}/day${1}.go
sed -i "7 i \ \ \"github.com/Segmev/AdventOfCode2022/day${1}\"" main.go
sed -i "45 i \ \ \ \ \ \ \ \ \ \ \ \ \"day${1}\": day${1}.Main," main.go
