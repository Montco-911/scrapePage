

```

docker exec -it kind-control-plane  /bin/bash

k exec -it task-pv-pod -- /bin/bash

kubectl alpha debug -it 'task-pv-pod' --image=busybox --target='task-pv-container'

kubectl exec --stdin --tty zookeeper-5cdcd9f4f-mldzh -- /bin/bash


```

dataDir=/var/lib/zookeeper/data
dataLogDir=/var/lib/zookeeper/log

