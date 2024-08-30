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
