#!/bin/bash

echo "Checking PostgreSQL ports..."
echo "Port 5432:"
sudo lsof -i :5432 || echo "Port 5432 is not in use"
echo ""

echo "Port 5433:"
sudo lsof -i :5433 || echo "Port 5433 is not in use"
echo ""

echo "Port 5434:"
sudo lsof -i :5434 || echo "Port 5434 is not in use"
echo ""

echo "Checking other app ports..."
echo "Port 8080 (backend):"
sudo lsof -i :8080 || echo "Port 8080 is not in use"
echo ""

echo "Port 3000 (frontend):"
sudo lsof -i :3000 || echo "Port 3000 is not in use"
echo ""

echo "Checking Docker status:"
docker ps -a
echo ""

echo "Docker networks:"
docker network ls