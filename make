docker-compose -f docker-compose-dev.yml up --abort-on-container-exit  --exit-code-from go_server_dev
docker-compose -f docker-compose-test.yml up --exit-code-from test_sandbox