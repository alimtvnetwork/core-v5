package corecreator

type newCreator struct{}

func (it newCreator) CreatorMap(
	scopeName string,
	additionalMaps ...map[string]any,
) Creator {
	return Creator{}
}
