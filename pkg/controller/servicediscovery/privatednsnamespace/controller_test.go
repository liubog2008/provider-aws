package privatednsnamespace

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/test"

	svcapitypes "github.com/crossplane/provider-aws/apis/servicediscovery/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
	svcclient "github.com/crossplane/provider-aws/pkg/clients/servicediscovery"
	"github.com/crossplane/provider-aws/pkg/clients/servicediscovery/fake"
)

const (
	validOpID        string = "123"
	validNSID        string = "ns-id"
	validDescription string = "valid description"
	validArn         string = "arn:string"
)

type args struct {
	client svcclient.Client
	kube   client.Client
	cr     *svcapitypes.PrivateDNSNamespace
}

func TestObserve(t *testing.T) {
	type want struct {
		cr     *svcapitypes.PrivateDNSNamespace
		result managed.ExternalObservation
		err    error
	}

	cases := map[string]struct {
		args
		want
	}{
		"NewNoOpID": {
			args: args{
				client: &fake.MockServicediscoveryClient{},
				kube:   nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{},
				},
				result: managed.ExternalObservation{
					ResourceExists: false,
				},
			},
		},
		"NewOperationSubmitted": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("SUBMITTED"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists: true,
				},
			},
		},
		"NewOperationPending": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("PENDING"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists: true,
				},
			},
		},
		"NewOperationFailed": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("FAIL"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Unavailable()},
							},
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists: true,
				},
			},
		},
		"DeletingOperationFail": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("FAIL"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Deleting()},
							},
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Unavailable()},
							},
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists: false,
				},
			},
		},
		"NewOpDoneNSNotFound": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("SUCCESS"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
					MockGetNamespace: func(input *svcsdk.GetNamespaceInput) (*svcsdk.GetNamespaceOutput, error) {
						return &svcsdk.GetNamespaceOutput{}, awserr.New(svcsdk.ErrCodeNamespaceNotFound, "namespace not found", fmt.Errorf("err"))
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Unavailable()},
							},
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists: false,
				},
			},
		},
		"NewExternalNameNotSet": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("SUCCESS"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
					MockGetNamespace: func(input *svcsdk.GetNamespaceInput) (*svcsdk.GetNamespaceOutput, error) {
						if awsclient.StringValue(input.Id) != validNSID {
							return &svcsdk.GetNamespaceOutput{}, nil
						}
						return &svcsdk.GetNamespaceOutput{
							Namespace: &svcsdk.Namespace{
								Arn:         aws.String(validArn),
								Name:        aws.String(validNSID),
								Description: aws.String(validDescription),
							},
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
							Description: aws.String(validDescription),
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{"crossplane.io/external-name": validNSID},
					},
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
							Description: aws.String(validDescription),
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Available()},
							},
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists:          true,
					ResourceLateInitialized: true,
					ResourceUpToDate:        true,
				},
			},
		},
		"NewWithExternalName": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockGetOperation: func(input *svcsdk.GetOperationInput) (*svcsdk.GetOperationOutput, error) {
						if awsclient.StringValue(input.OperationId) != validOpID {
							return &svcsdk.GetOperationOutput{}, nil
						}
						return &svcsdk.GetOperationOutput{
							Operation: &svcsdk.Operation{
								Status:  aws.String("SUCCESS"),
								Targets: map[string]*string{"NAMESPACE": aws.String(validNSID)},
							},
						}, nil
					},
					MockGetNamespace: func(input *svcsdk.GetNamespaceInput) (*svcsdk.GetNamespaceOutput, error) {
						if awsclient.StringValue(input.Id) != validNSID {
							return &svcsdk.GetNamespaceOutput{}, nil
						}
						return &svcsdk.GetNamespaceOutput{
							Namespace: &svcsdk.Namespace{
								Arn:         aws.String(validArn),
								Name:        aws.String(validNSID),
								Description: aws.String(validDescription),
							},
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{"crossplane.io/external-name": validNSID},
					},
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{"crossplane.io/external-name": validNSID},
					},
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
							Description: aws.String(validDescription),
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Available()},
							},
						},
					},
				},
				result: managed.ExternalObservation{
					ResourceExists:          true,
					ResourceLateInitialized: true,
					ResourceUpToDate:        true,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := newExternal(tc.kube, tc.client, []option{useHooks})

			o, err := e.Observe(context.Background(), tc.args.cr)

			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.cr, tc.args.cr, test.EquateConditions(),
				cmpopts.IgnoreFields(v1.Condition{}, "LastTransitionTime"),
			); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.result, o); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type want struct {
		cr     *svcapitypes.PrivateDNSNamespace
		result managed.ExternalCreation
		err    error
	}

	cases := map[string]struct {
		args
		want
	}{
		"NewPrivateDNSNamespace": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockCreatePrivateDNSNamespace: func(input *svcsdk.CreatePrivateDnsNamespaceInput) (*svcsdk.CreatePrivateDnsNamespaceOutput, error) {
						return &svcsdk.CreatePrivateDnsNamespaceOutput{
							OperationId: aws.String(validOpID),
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
							Tags: []*svcapitypes.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
							Tags: []*svcapitypes.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Creating()},
							},
						},
					},
				},
				result: managed.ExternalCreation{},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := newExternal(tc.kube, tc.client, []option{useHooks})

			o, err := e.Create(context.Background(), tc.args.cr)

			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.cr, tc.args.cr, test.EquateConditions(),
				cmpopts.IgnoreFields(v1.Condition{}, "LastTransitionTime"),
			); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.result, o); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type want struct {
		cr  *svcapitypes.PrivateDNSNamespace
		err error
	}

	cases := map[string]struct {
		args
		want
	}{
		"NewPrivateDNSNamespace": {
			args: args{
				client: &fake.MockServicediscoveryClient{
					MockDeleteNamespace: func(input *svcsdk.DeleteNamespaceInput) (*svcsdk.DeleteNamespaceOutput, error) {
						return &svcsdk.DeleteNamespaceOutput{
							OperationId: aws.String(validOpID),
						}, nil
					},
				},
				kube: nil,
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
					},
				},
			},
			want: want{
				cr: &svcapitypes.PrivateDNSNamespace{
					Spec: svcapitypes.PrivateDNSNamespaceSpec{
						ForProvider: svcapitypes.PrivateDNSNamespaceParameters{
							Region: "eu-central-1",
							Name:   aws.String("test"),
							CustomPrivateDNSNamespaceParameters: svcapitypes.CustomPrivateDNSNamespaceParameters{
								VPC: aws.String("default"),
							},
						},
					},
					Status: svcapitypes.PrivateDNSNamespaceStatus{
						AtProvider: svcapitypes.PrivateDNSNamespaceObservation{
							OperationID: aws.String(validOpID),
						},
						ResourceStatus: xpv1.ResourceStatus{
							ConditionedStatus: xpv1.ConditionedStatus{
								Conditions: []xpv1.Condition{xpv1.Deleting()},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := newExternal(tc.kube, tc.client, []option{useHooks})

			err := e.Delete(context.Background(), tc.args.cr)

			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.cr, tc.args.cr, test.EquateConditions(),
				cmpopts.IgnoreFields(v1.Condition{}, "LastTransitionTime"),
			); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
