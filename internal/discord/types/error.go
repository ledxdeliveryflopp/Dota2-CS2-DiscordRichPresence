package types

type DotaPresenceError struct {
	Presence DotaPresence
	Err      error
}

func (e *DotaPresenceError) SetErrors(presence *DotaPresence, err error) {
	e.Presence.State = presence.State
	e.Presence.Details = presence.Details
	e.Presence.MainImage = presence.MainImage
	e.Presence.HeroReadableName = presence.HeroReadableName
	e.Presence.SmallImage = presence.SmallImage
	e.Err = err
	return
}

type CsGoPresenceError struct {
	Presence CsGoPresence
	Err      error
}

func (c *CsGoPresenceError) SetErrors(presence *CsGoPresence, err error) {
	c.Presence.State = presence.State
	c.Presence.Details = presence.Details
	c.Err = err
	return
}
