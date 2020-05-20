package config

type AliasConfig struct {
	ConfigMap
}

// TODO at what point should the alias top level be added? Lazily on first write, I think;
// but then we aren't initializing the default aliases that we want. also want to ensure that we
// don't re-add the default aliases if people delete them.
//
// I think I'm not going to worry about it for now; config setup will add the aliases and existing
// users can take some step later on to add them if they want them. we'll otherwise lazily create
// the aliases: section on first write.

func (a *AliasConfig) Exists(alias string) bool {
	if a.Empty() {
		return false
	}
	value, _ := a.GetStringValue(alias)

	return value != ""
}

func (a *AliasConfig) Get(alias string) string {
	value, _ := a.GetStringValue(alias)

	return value
}

func (a *AliasConfig) Add(alias, expansion string) error {
	if a.Root == nil {
		// a.Root = initAliasConfig() // something like this?
		// TODO create mapping node
		// TODO initialize aliases key in config
	}

	// TODO add k/v nodes
	// TODO write config

	// TODO how to do all of this without reference to parent?

	return nil
}

func (a *AliasConfig) Delete(alias string) error {
	// TODO when we get to gh alias delete
	return nil
}
