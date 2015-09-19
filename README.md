### continuous deployment service for containers

* register nodes to service
* deploy docker images automaticly (or manually)
* deploy strategies

#### Definitions

| Name | Explanation |
| ---- | ----------- |
| Node   | machine that running docker service |
| Region | indicate diffrent data center |
| Task   | sepecified service image |
| Strategies | auto-scaling, single node, or mass-deployment |

#### APIs

|   |   |   |
| --- | --- | --- |
| /region | /add | in: {desc, geo} out: {region_id} |
|         | /del | in: {region_id} |
| /node   | /add | in: {region_id, ip, port, key, desc} out: {node_id} |
|         | /del | in: {node_id} |
|         | /mod |   |
|         | /ls  |   |
| /task   | /add | in: {image_id, arguments, hooks} out: {task_id} |
|         | /ls  | { []task } |
