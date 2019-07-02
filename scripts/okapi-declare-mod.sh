curl -w '\n' -X POST -D -   \
   -H "Content-type: application/json"   \
   -d @../descriptors/ModuleDescriptor.json \
   http://localhost:9130/_/proxy/modules
