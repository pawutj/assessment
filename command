Run Dev
docker-compose -f docker-compose-dev.yml up --abort-on-container-exit  --exit-code-from go_server_dev

Run Test Sandbox
docker-compose -f docker-compose-test.yml up --exit-code-from test_sandbox

Run Deploy
docker-compose -f docker-compose-deploy.yml up --abort-on-container-exit  --exit-code-from go_server_deploy