#!/bin/bash
cd /home/eisandbar/ytlive
git pull origin main
docker compose up client --build -d