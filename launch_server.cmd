@echo off

taskkill /IM server.exe /F
nircmd-x64\nircmd.exe exec hide "server.exe"