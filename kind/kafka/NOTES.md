

```

docker exec -it kind-control-plane  /bin/bash

k exec -it task-pv-pod -- /bin/bash

kubectl alpha debug -it 'task-pv-pod' --image=busybox --target='task-pv-container'

kubectl exec --stdin --tty zookeeper-cdd8d6757-k5lqb  -- /bin/bash


```

dataDir=/var/lib/zookeeper/data
dataLogDir=/var/lib/zookeeper/log

