tilt:
	ctlptl apply -f ctlptl_config.yaml
	tilt up -f Tiltfile

tilt-ci:
	ctlptl apply -f ctlptl_config.yaml
	tilt up -f Tiltfile.ci

tilt-clean:
	ctlptl delete cluster kind-tilt
