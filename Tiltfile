docker_prune_settings(
	disable = False,
	max_age_mins = 360,
	num_builds = 0,
	interval_hrs = 1,
	keep_recent = 2
)
docker_build(
	"adgayle/book-library",
	"./",
	dockerfile="Dockerfile.book-library"
)

load('ext://namespace', 'namespace_yaml')
k8s_yaml(namespace_yaml('applications'), allow_duplicates=False)

k8s_yaml(
	helm(
		"./helm/charts/book-library",
		name="book-library",
		namespace="applications",
		values="./helm/charts/book-library/values.yaml"
	)
)

k8s_resource(
	new_name="book-library-aux",
  objects=[
		"book-library:ServiceAccount:applications",
		"book-library:Ingress:applications"
	]
)
