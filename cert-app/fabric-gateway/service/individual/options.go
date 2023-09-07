package individual

type individualOptions struct {
	avatar string   // An URL to the individual's avatar
	name   string   // Individualname
	org    string   // Organisation of the individual
	mspID  string   // MSPID of the individual
	roles  []string // The roles
}

type IndividualOption func(*individualOptions) error

// Individual Options Implementations
func WithAvatar(avatar string) IndividualOption {
	return func(o *individualOptions) error {
		o.avatar = avatar
		return nil
	}
}

// WithName option
func WithName(name string) IndividualOption {
	return func(o *individualOptions) error {
		o.name = name
		return nil
	}
}

func WithOrg(orgName string) IndividualOption {
	return func(o *individualOptions) error {
		o.org = orgName
		return nil
	}
}

func WithMspid(mspID string) IndividualOption {
	return func(o *individualOptions) error {
		o.mspID = mspID
		return nil
	}
}

func WithRoles(roles []string) IndividualOption {
	return func(o *individualOptions) error {
		o.roles = roles
		return nil
	}
}
