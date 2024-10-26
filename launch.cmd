@echo off
taskkill /IM server.exe /F

nircmd-x64\nircmd.exe exec hide "server.exe"
start http://127.0.0.1:3223/index.html
