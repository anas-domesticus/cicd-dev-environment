### Main Tilt dev environment
docker_build('example', 'cmd/example', dockerfile = 'cmd/example/Dockerfile')

k8s_yaml(kustomize('k8s/overlays/dev/example'))

k8s_resource('example-deployment', labels=['workload'])
