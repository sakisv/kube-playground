# Step 1

The first step was to get a small service running that I could access from the outside
world on port 80

To do that I:

1. Wrote a small app in go that would respond to http requests and pushed it to docker in
   a public repo
1. Created a deployment for this image (miniserver.yml)
1. Exposed this deployment with a [NodePort][] service (miniserver-service.yml)

At this point I could access it with `localhost:NODE_PORT` (or `ip_address:NODE_PORT`)

The next step was to remove the PORT from the url:

1. I created an ingress service (ingress.yml)
1. I applied the [mandatory yaml][] to create an nginx ingress controller

After this I could just access my service on `ip_address`


[NodePort]: https://kubernetes.io/docs/concepts/services-networking/service/#nodeport
[mandatory yaml]: https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml   
