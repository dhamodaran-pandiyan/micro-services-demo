
#### Introduction:
This configuration defines a deployment for the Go application, which will create three replicas of the application running in separate pods. 

It also defines a service for the application, which will expose the application to external traffic using a load balancer.

#### How to run:
To deploy this application to Kubernetes, you will need to have a Kubernetes cluster set up, and you will need to have the kubectl command-line tool installed.

You can install kubectl using the instructions in the official Kubernetes documentation: https://kubernetes.io/docs/tasks/tools/install-kubectl/.

Once you have kubectl installed, you can deploy the application to your cluster using the following command:

> kubectl apply -f deployment.yaml

This will create the deployment and service in your cluster, and the application will be accessible at the cluster IP address of the service. You can get the cluster IP address using the following command:

> kubectl get service go-app

This will display the details of the service, including the cluster IP address. You can then access the microservices at 

http://CLUSTER-IP:8000/service1, 
http://CLUSTER-IP:8000/service2, and 
http://CLUSTER-IP:8000/service3, 

where ```CLUSTER-IP``` is the cluster IP address of the service.