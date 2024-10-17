// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	v1alpha1 "github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BandwidthAssociate.
func (mg *BandwidthAssociate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BandwidthID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BandwidthIDRef,
		Selector:     mg.Spec.ForProvider.BandwidthIDSelector,
		To: reference.To{
			List:    &BandwidthList{},
			Managed: &Bandwidth{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BandwidthID")
	}
	mg.Spec.ForProvider.BandwidthID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BandwidthIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EIPID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.EIPIDRef,
		Selector:     mg.Spec.ForProvider.EIPIDSelector,
		To: reference.To{
			List:    &EIPList{},
			Managed: &EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EIPID")
	}
	mg.Spec.ForProvider.EIPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EIPIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BandwidthID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.BandwidthIDRef,
		Selector:     mg.Spec.InitProvider.BandwidthIDSelector,
		To: reference.To{
			List:    &BandwidthList{},
			Managed: &Bandwidth{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BandwidthID")
	}
	mg.Spec.InitProvider.BandwidthID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BandwidthIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EIPID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.EIPIDRef,
		Selector:     mg.Spec.InitProvider.EIPIDSelector,
		To: reference.To{
			List:    &EIPList{},
			Managed: &EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EIPID")
	}
	mg.Spec.InitProvider.EIPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EIPIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Eipv3Associate.
func (mg *Eipv3Associate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PublicipID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PublicipIDRef,
		Selector:     mg.Spec.ForProvider.PublicipIDSelector,
		To: reference.To{
			List:    &EIPList{},
			Managed: &EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PublicipID")
	}
	mg.Spec.ForProvider.PublicipID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PublicipIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PublicipID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.PublicipIDRef,
		Selector:     mg.Spec.InitProvider.PublicipIDSelector,
		To: reference.To{
			List:    &EIPList{},
			Managed: &EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PublicipID")
	}
	mg.Spec.InitProvider.PublicipID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PublicipIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InternetGateway.
func (mg *InternetGateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetList{},
			Managed: &v1alpha1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetList{},
			Managed: &v1alpha1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}
