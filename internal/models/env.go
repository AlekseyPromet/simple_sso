package models

type TypeEnv string

const (
	LocalEnv      TypeEnv = "local"
	DevelopEnv    TypeEnv = "develop"
	ProductionEnv TypeEnv = "production"
)
