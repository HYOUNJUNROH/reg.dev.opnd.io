```
docker compsoe pull
docker compose up -d
```
reg.dev.opnd.io/gc/bite-web-api:release
jc21/nginx-proxy-manager:2.9.18
reg.dev.opnd.io/gc/backend-admin:main
reg.dev.opnd.io/gc/frontend-admin:main
postgres:12-alpine 
reg.dev.opnd.io/gc/bite-web-app:release
reg.dev.opnd.io/gc/bite-web-app:main-eng
reg.dev.opnd.io/gc/bite-web-api:main-eng
reg.dev.opnd.io/gc/backend-admin:main-eng

docker save reg.dev.opnd.io/gc/bite-web-api:release > bite-web-app-main.tar
docker save jc21/nginx-proxy-manager:2.9.18 > nginx-proxy-manager.tar
docker save reg.dev.opnd.io/gc/backend-admin:main > backend-admin-main.tar
docker save reg.dev.opnd.io/gc/frontend-admin:main > frontend-admin-main.tar
docker save postgres:12-alpine > postgres-12-alpine.tar
docker save reg.dev.opnd.io/gc/bite-web-app:release > bite-web-app-release.tar
docker save reg.dev.opnd.io/gc/bite-web-app:main-eng > bite-web-app-main-eng.tar
docker save reg.dev.opnd.io/gc/bite-web-api:main-eng > bite-web-api-main-eng.tar


bite-web-app-main.tar
nginx-proxy-manager.tar