package packageB

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -generate types,skip-prune --package=packageB -o externalref.gen.go spec.yaml
