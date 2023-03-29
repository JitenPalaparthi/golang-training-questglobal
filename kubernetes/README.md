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




