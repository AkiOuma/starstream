package template

var Usecase = `package usecase

import (
	%s	
)

// Interface
%s

// Implement
%s

var _ %s = (*%s)(nil)

// Constructor
%s

// Implement methods
%s
`

var UsecaseInterface = `type %s interface {
	%s
}
`

var UsecaseImplement = `type %s struct {
	repo repository.%s
	svc  service.%s
}
`
var UsecaseConstructor = `func New%s(
	repo repository.%s,
	svc service.%s,
) %s {
	return &%s{svc: svc, repo: repo}
}
`

var UsecaseImplMethod = `func (u *%s) %s(%s) (%s) {
	%s
}
`

var UsecaseRemoveTemplate = `	  result, err := u.Find%ss(ctx, querier)
	if err != nil {
		return err
	}
	root := make([]%s, 0, result.Count)
	for _, v := range result.Data {
		root = append(root, v.%s)
	}
	return u.repo.Remove%ss(ctx, root)
`
