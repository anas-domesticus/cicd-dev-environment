tilt:
	ctlptl apply -f ctlptl_config.yaml
	tilt up -f Tiltfile

tilt-ci:
	ctlptl apply -f ctlptl_config.yaml
	tilt up -f Tiltfile.ci

tilt-cd:
	ctlptl apply -f ctlptl_config.yaml
	tilt up -f Tiltfile.cd

tilt-clean:
	ctlptl delete cluster kind-tilt

argo-pass:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo