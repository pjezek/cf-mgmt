package isosegment

import (
	"github.com/pivotalservices/cf-mgmt/config"
	"github.com/xchapter7x/lo"
)

// Segment represents a Cloud Foundry isolation segment.
type Segment struct {
	Name string
	GUID string
}

// Updater performs the required updates to acheive the desired state wrt isolation segments.
type Updater struct {
	Cfg config.Reader

	DryRun  bool // print the actions that would be taken, make no changes
	CleanUp bool // delete/restrict access to any iso segments not identified in the config

	cc manager
}

// Ensure creates any isolation segments that do not yet exist,
// and optionally removes unneeded isolation segments.
func (u *Updater) Ensure() error {
	desired, err := u.allDesiredSegments()
	if err != nil {
		return err
	}
	current, err := u.cc.GetIsolationSegments()
	if err != nil {
		return err
	}

	c := classify(desired, current)
	return c.update("", u.create, u.delete)
}

// Entitle ensures that each org is entitled to the isolation segments it needs to use.
func (u *Updater) Entitle() error {
	spaces, err := u.Cfg.GetSpaceConfigs()
	if err != nil {
		return err
	}

	orgs, err := u.Cfg.GetOrgConfigs()
	if err != nil {
		return err
	}

	// build up a list of segments required by each org (grouped by org name)
	// this includes segments used by all of the orgs spaces, as well as the
	// org's default segment
	sm := make(map[string][]Segment)
	for _, space := range spaces {
		if s := space.IsoSegment; s != "" {
			sm[space.Org] = append(sm[space.Org], Segment{Name: s})
		}
	}
	for _, org := range orgs {
		if s := org.DefaultIsoSegment; s != "" {
			sm[org.Org] = append(sm[org.Org], Segment{Name: s})
		}
	}

	for org, desiredSegments := range sm {
		currentSegments, err := u.cc.EntitledIsolationSegments(org)
		if err != nil {
			return err
		}

		c := classify(desiredSegments, currentSegments)
		err = c.update(org, u.entitle, u.revoke)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Updater) create(s *Segment, _ string) error {
	if u.DryRun {
		lo.G.Info("[dry-run]: create segment", s.Name)
		return nil
	}
	return u.cc.CreateIsolationSegment(s.Name)
}

func (u *Updater) delete(s *Segment, _ string) error {
	if !u.CleanUp {
		return nil
	}
	if u.DryRun {
		lo.G.Infof("[dry-run]: delete segment %s (%s)", s.Name, s.GUID)
		return nil
	}
	return u.cc.DeleteIsolationSegment(s.Name)
}

func (u *Updater) entitle(s *Segment, orgName string) error {
	if u.DryRun {
		lo.G.Infof("[dry-run]: entitle org %s to iso segment %s", orgName, s.Name)
		return nil
	}
	return u.cc.EnableOrgIsolation(orgName, s.Name)
}

func (u *Updater) revoke(s *Segment, orgName string) error {
	if !u.CleanUp {
		return nil
	}
	if u.DryRun {
		lo.G.Infof("[dry-run]: revoke iso segment %s from org %s", s.Name, orgName)
		return nil
	}
	return u.cc.RevokeOrgIsolation(orgName, s.Name)
}

// allDesiredSegments iterates through the cf-mgmt configuration for all
// orgs and spaces and builds the complete set of isolation segments that
// should exist
func (u *Updater) allDesiredSegments() ([]Segment, error) {
	orgs, err := u.Cfg.GetOrgConfigs()
	if err != nil {
		return nil, err
	}
	spaces, err := u.Cfg.GetSpaceConfigs()
	if err != nil {
		return nil, err
	}

	segments := make(map[string]struct{})
	for _, org := range orgs {
		if org.DefaultIsoSegment != "" {
			segments[org.DefaultIsoSegment] = struct{}{}
		}
	}
	for _, space := range spaces {
		if space.IsoSegment != "" {
			segments[space.IsoSegment] = struct{}{}
		}
	}

	result := make([]Segment, 0, len(segments))
	for k := range segments {
		result = append(result, Segment{Name: k})
	}
	return result, nil
}
