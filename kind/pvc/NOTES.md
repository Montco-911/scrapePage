

```
k exec -it task-pv-pod -- /bin/bash

kubectl alpha debug -it 'task-pv-pod' --image=busybox --target='task-pv-container'



```