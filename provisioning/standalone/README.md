# Standalone Applications

## What is a Standalone Application ?

**Standalone applications** are autonomous web services which are exposed by **http**. Those services have to be packaged in a **container**.
An example could be a **container** exposing a **static web page**. If the **container** includes **all the dependencies**, it can be considered a **standalone application**.

## How Kubernetes manages Standalone Applications ?

**Standalone applications** are just **instances** of a **template** whose **environmentType** is **Standalone**. Then instances are managed by the **instance operator**, which is a **controller** capable of creating what is needed by a **standalone application**.

## Why standalone applications ?

**Standalone applications** allow users to create their **templates**, without having to manage the **cluster infrastructure** (_ingres_, _service_, _deploy_, and _pvc_ are created and managed by the **instance operator**). So a **developer** can focus on the **container** without thinking about the **infrastructure**.

## How to create a standalone application ?

In order to create a container which can be used as standalone application, some simple rules have to be followed:

-   The **port** which the service is exposed has to be set using the environment variable **$CROWNLABS_LISTEN_PORT**.
-   The **user** running the container cannot be **root**.
-   The listening service has to use as **basepath** the environment variable **$CROWNLABS_BASE_PATH** or setup the **rewriteURL** inside the template (_see next paragraph_)
-   Must answer a **200** if a **GET request** is sent to **$CROWNLABS_BASE_PATH/** (_rewriteURL to false_) or **/** (_rewriteURL to true_).

### RewriteURL vs CROWNLABS_BASE_PATH

Let's make an example. We are deploying an application as **standalone application**, so we will create a template with a **yaml file** which look like this:

```yaml
apiVersion: crownlabs.polito.it/v1alpha2
kind: Template
metadata:
    name: my-app
    namespace: workspace-sid
spec:
    prettyName: My Application
    description: Just an example template
    environmentList:
        - name: my-app-environment
          environmentType: Standalone
          mode: Standard
          image: myapplicationimage
          resources:
              cpu: 2
              memory: 2G
              reservedCPUPercentage: 25
          rewriteURL: false
          persistent: false
    workspace.crownlabs.polito.it/WorkspaceRef:
        name: sid
    deleteAfter: 30d
```

Then we will create an **instance**, with a **yaml** or using the **crownlabs frontend**. The generated **url** to access the **instance** will look something like this: `https://crownlabs.polito.it/instance/4e46/app/`. So when we will send a request to this **url**, the **application** will receive a **request** for `https://crownlabs.polito.it/instance/4e46/app/` and not just `/` . So the **application** has to be aware of it (_see the example down below_).

In some cases our **containers** don't run software written by ourself. When is not possible to configure the application **basepath**, the **rewriteUrl** flag inside the **template** has to be set. This enables the **rewrite-url** feature inside the used **ingress**, which translates all **request's url** (towards our **application**) in `/`. So will be possible for the **application** to not set the **basepath** and work.

## Example

In this example we are going to learn how to deploy a simple echo server using **golang**, which will answer to a **GET http request** with a predefined message.

### Create the server

Create a **main.go** file which will contain the source code of the server and **copy/paste** the following code inside **main.go**

```go
import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

func main() {
	port := os.Args[1]
	basepath := os.Args[2]
	http.HandleFunc(fmt.Sprintf("%s/hello", basepath), hello)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
```

Let's test it running it with `go run 8888 /aaa`.

In the meanwhile open another terminal and launch the command `curl localhost:8888/aaa/hello`, and you will get `Hello!` as response.

### Build the container

Let's create a **Dockerfile** ` and **copy/paste** the following code inside of it.

```dockerfile
FROM golang:1.18
ENV CROWNLABS_LISTEN_PORT=8888
ENV CROWNLABS_BASE_PATH=/basepath
COPY main.go main.go
RUN go build main.go
RUN useradd -ms /bin/bash myuser
USER myuser
CMD exec ./main ${CROWNLABS_LISTEN_PORT} ${CROWNLABS_BASE_PATH}
```

Build the container with `docker build -t ${YOUR_DOCKERHUB_USERNAME}/echotest` and push it on dockerhub with `docker push ${YOUR_DOCKERHUB_USERNAME}/echotest`.

### Create the template

Create a file called **template.yaml** and **copy/paste** the following code:

```yaml
apiVersion: crownlabs.polito.it/v1alpha2
kind: Template
metadata:
    name: echotest
    namespace: workspace-sid # a workspace where you have the permission to create templates
spec:
    deleteAfter: never
    description: An http server used for testing
    environmentList:
        - environmentType: Standalone
          guiEnabled: true
          image: YOUR_DOCKERHUB_USERNAME/echotest
          mode: Standard
          name: default
          persistent: false
          resources:
              cpu: 1
              memory: 1000M
              reservedCPUPercentage: 50
          rewriteURL: false
    prettyName: Echo HTTP Server
    workspace.crownlabs.polito.it/WorkspaceRef:
        name: sid # choosen workspace name
```

And **apply** it with `kubectl apply -f template.yaml`

### Create an instance and try it

Open **CrownLabs dashboard** and create a new instance.

<img align="left" src="./docs/img.png"><br>

Click on **connect** and you will see an `Hello!` message.
