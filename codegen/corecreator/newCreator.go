package corecreator

type newCreator struct{}

func (it newCreator) CreatorMap(
	scopeName string,
	additionalMaps ...map[string]interface{},
) Creator {

}
