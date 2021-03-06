## How to run the Koji test

Firstly, you need to start the koji container:

```
sudo ./tools/run-koji-container.sh start
```

This command starts a kojihub instance available at
http://localhost:8080/kojihub . You can test that it started successfully
by running:
```
koji --server=http://localhost:8080/kojihub --user=osbuild --password=osbuildpass --authtype=password hello
```

Now, you can run the koji test using:
```
go test -v -tags koji_test ./internal/upload/koji
```

The test is run on each PR in the Github CI. See `.github/workflows/tests.yml`
for more details.

To stop and remove the koji container, use the following command:

```
sudo ./tools/koji/run-koji-container.sh stop
```
