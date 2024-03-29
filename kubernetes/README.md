```kubectl cluster-info```

- if you are acceccsing more than one cluster..

```kubectl cluster-info --context kind-kind```

- set us the cluster context as the default cluster
```kubectl config use-context kind-kind-cluster2```

- list of pods running in all namespaces

```kubectl get pods -A```

- create/apply a pod using yaml file. This runs in the current contect

```kubectl apply -f sample-pod.yaml```

- to create a pod on a specific context/cluster

```kubectl apply -f sample-pod.yaml --context=kind-test```

- to create a pod in a namespace

```kubectl create -f sample-pod.yaml -n sample```

- to get list of pods in a namespace

```kubectl get pods -n default```

- to get logs of a container

```kubectl logs myapp -c  sampleapp```

- to get logs of a pod

```kubectl logs myapp```

- to get informaton about a pod

```kubectl describe pods myapp```

- to delete a pod 

```kubectl delete pod myapp```

- to create a namespace through command line 

```kubectl create ns demo```

- to create namespace through yaml file(can use apply or create command)

```kubectl create -f sample-namespace.yaml```

- edit any deplyment

```kubectl edit deply sample-deployent.yaml -n dev```

- deployment means there are 1 or more pods running.
- when user wants to access any pod, there must a service.
- the service is an abstraction layer for pods[deployments, replicasets,replicationControllers,normal pod etc]
- service has a load balancer by default

- Service types:
1. ClusterIP : By default the type of the service is .The service can be accessed inside the cluster.
2. NodePort: The service can be accessed using node ip address. Automatically port is picked by the kubernetes if not given. The port range is 30000-32767
The application can be accessed using node IP address.
3. LoadBalancer: On any cloud provider clusters, this gives a url/endpoint or ip address to access the service.


- to get services 

```kubectl get svc -n dev```

- to test the application(only for dev purpose)

```kubectl port-forward deploy/myappdeployment 58080:8080 -n dev```

- to scale n dev

```kubectl scale --replicas=10 deploy/myappdeployment -n dev```


- port-forward  (Only development and testing)

```kubectl port-forward service/myappservice 58080:8080 -n dev```

- to encode to base64

```echo "host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai" | base64```

- to decode from base64

```echo "aG9zdD1sb2NhbGhvc3QgdXNlcj1wb3N0Z3JlcyBwYXNzd29yZD1wb3N0Z3JlcyBkYm5hbWU9Y29udGFjdHNiZCBwb3J0PTU1NDMyIHNzbG1vZGU9ZGlzYWJsZSBUaW1lWm9uZT1Bc2lhL1NoYW5naGFpCg==" | base64 -d```


