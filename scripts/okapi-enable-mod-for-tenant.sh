curl -w '\n' -X POST -D -   \
    -H "Content-type: application/json"   \
    -d @../descriptors/TenantModuleDescriptor.json \
    http://localhost:9130/_/proxy/tenants/diku/modules