### Use with OpenStack

Example:

```
 terraformer import openstack --resources=compute,networking --regions=RegionOne
```

List of supported OpenStack services:

*   `blockstorage`
    * `openstack_blockstorage_volume_v1`
    * `openstack_blockstorage_volume_v2`
    * `openstack_blockstorage_volume_v3`
*   `compute`
    * `openstack_compute_instance_v2`
*   `networking`
    * `openstack_networking_secgroup_v2`
    * `openstack_networking_secgroup_rule_v2`




The following vairbales should be expported 


export OS_AUTH_URL=http://10.102.192.221:5000/v3
export OS_PROJECT_ID=c2e747a5c1874ed981c3e6bb7a513dec
export OS_PROJECT_NAME="admin"
export OS_USER_DOMAIN_NAME="default"
export OS_PROJECT_DOMAIN_ID="default"
export OS_DOMAIN_ID="default"
export OS_USERNAME="admin"
export OS_PASSWORD="password"
export OS_REGION_NAME="RegionOne"
export OS_IDENTITY_API_VERSION=3
