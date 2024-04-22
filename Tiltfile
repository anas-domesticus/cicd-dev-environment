### Main Tilt dev environment

load('ext://helm_resource', 'helm_resource', 'helm_repo')
update_settings (k8s_upsert_timeout_secs = 90)

docker_build('example', 'cmd/example', dockerfile = 'cmd/example/Dockerfile')

k8s_yaml(kustomize('k8s/overlays/dev/'))

k8s_resource('example-deployment', labels=['workload'])

