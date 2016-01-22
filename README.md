This repository contains some simple examples to demo Docker on top of
SUSE Linux Enterprise.

## simple_webapp

This demo shows a simple web application running on top of either SLE12 or SLE11SP3.

The web application is written using Go, its sources are inside of simple_web/webapp_demo.

### Requirements

* Docker daemon running.
* Official SUSE Linux Enterprise Docker images installed.
* The Go compiler
* make

### Building

Enter the `simple_webapp` directory and just type: `make`.

### Running

Type:

`
docker run -p 8080:8080 demo-sles11sp3 /demo/webapp
`

and then point your browser to `http://localhost:8080`

