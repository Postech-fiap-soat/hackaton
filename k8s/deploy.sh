#!/bin/bash

#kind create cluster --config=kind.yml --name=soatcluster

kubectl apply -f soatdb-configmap.yml
kubectl apply -f soatdb-deploy.yml
kubectl apply -f soatdb.yml

kubectl apply -f mshackaton-configmap.yml
kubectl apply -f mshackaton-deploy.yml
kubectl apply -f mshackaton.yml