package payload

import (
	fmt "fmt"
)

func genPayload(nAudiences int) Payload {
	s := "abcdefojkljkds"
	p := Payload{
		WeboId: s,
		Action: Action_ADD,
	}
	aud := &Audience{AudienceId: s, DatatransferId: s}
	auds := make([]*Audience, 0, nAudiences)
	for i := 1; i <= nAudiences; i++ {
		auds = append(auds, aud)
	}
	p.Audiences = auds
	return p
}

func payloadsAreEqual(ref, comp Payload) error {
	if ref.WeboId != comp.WeboId {
		return fmt.Errorf("expected ProfileID '%s' received '%s'", ref.WeboId, comp.WeboId)
	}
	if ref.Action != comp.Action {
		return fmt.Errorf("expected Action '%s' received '%s'", ref.Action, comp.Action)
	}
	if len(ref.Audiences) != len(comp.Audiences) {
		return fmt.Errorf("expected %d audiences received %d", len(ref.Audiences), len(comp.Audiences))
	}
	var err error
	for ind, aud := range ref.Audiences {
		if err = audiencesAreEqual(*aud, *(comp.Audiences[ind])); err != nil {
			return fmt.Errorf("audience[%d]: %v", ind, err)
		}
	}
	return nil
}

func audiencesAreEqual(ref, comp Audience) error {
	if ref.AudienceId != comp.AudienceId {
		return fmt.Errorf("expected audience ID '%s' received '%s'", ref.AudienceId, comp.AudienceId)
	}
	if ref.DatatransferId != comp.DatatransferId {
		return fmt.Errorf("expected datatransfer ID '%s' received '%s'", ref.DatatransferId, comp.DatatransferId)
	}
	return nil
}
