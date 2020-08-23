package reference

type ChainedResolver struct {
	resolvers []Resolver
}

func (r ChainedResolver) ResolveKey(ref string) (*Key, error) {
	for _, r := range r.resolvers {
		key, err := r.ResolveKey(ref)
		if err != nil {
			return nil, err
		}
		if key != nil {
			return key, nil
		}
	}
	return nil, nil
}

func (r ChainedResolver) ResolveCert(ref string) (*Cert, error) {
	for _, r := range r.resolvers {
		cert, err := r.ResolveCert(ref)
		if err != nil {
			return nil, err
		}
		if cert != nil {
			return cert, nil
		}
	}
	return nil, nil
}

func (r ChainedResolver) ResolveAny(ref string) (interface{}, error) {
	for _, r := range r.resolvers {
		any, err := ResolveAny(ref, r)
		if err != nil {
			return nil, err
		}
		if any != nil {
			return any, nil
		}
	}
	return nil, nil
}