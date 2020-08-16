# Docker cpu_shares playground
I wanted to verify that mechanics of docker's `cpu_shares` fits my usecase, so I came up with this proof of concept which confirmed my expectations:

1) Containers with low `cpu_shares` have unrestricted CPU access in absense of containers with higher `cpu_shares`
2) As container with higher `cpu_shares` appears, resources are distributed accordingly
3) If a container with higher `cpu_shares` does not utilize enough resources, more resources are available to the container with lower `cpu_shares`

![](./demo.gif)