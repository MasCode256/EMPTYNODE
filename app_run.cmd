@echo off

pushd apps\%1\
g++ main.cpp -o main.exe
popd
