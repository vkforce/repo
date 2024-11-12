#!/bin/bash


POD_NAME="test-pod"
NAMESPACE="insights"

start_time=$(date +%s)

echo "Start Time: $start_time"

# kubectl apply -f pod.yaml -n $NAMESPACE
while true; do
  pod_status=$(kubectl get pod $POD_NAME -n $NAMESPACE -o=jsonpath='{.status.phase}')
  if [ "$pod_status" == "Running" ]; then
    break
  fi
  sleep 1
done

end_time=$(date +%s)

creation_duration=$((end_time - start_time))

formatted_duration=$(date -u -d @$creation_duration +'%H:%M:%S')

echo "Time taken for pod $POD_NAME to become ready: $formatted_duration"

# kubectl delete pod $POD_NAME -n $NAMESPACE
