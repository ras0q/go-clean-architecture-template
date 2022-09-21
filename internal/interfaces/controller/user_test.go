package controller

import (
	"context"
	"net/http"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/domain"
	"github.com/Ras96/go-clean-architecture-template/internal/usecases"
	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository/mock_repository"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_userControllerImpl_GetUser(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		req *GetUserRequest
	}
	type fields struct {
		ur repository.UserRepository
	}
	type setupFieldsFunc func(t *testing.T, args args, _ GetUserResponse, _ int) fields
	tests := map[string]struct {
		args        args
		want        GetUserResponse
		want1       int
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		"success": {
			args: args{
				ctx: context.Background(),
				req: &GetUserRequest{
					ID: 1,
				},
			},
			want: GetUserResponse{
				ID:    1,
				Name:  "test",
				Email: "test@example.com",
			},
			want1: http.StatusOK,
			setupFields: func(t *testing.T, args args, want GetUserResponse, _ int) fields {
				ctrl := gomock.NewController(t)
				mockur := mock_repository.NewMockUserRepository(ctrl)
				mockur.
					EXPECT().
					FindByID(args.ctx, args.req.ID).
					Return(domain.User{
						ID:    args.req.ID,
						Name:  want.Name,
						Email: want.Email,
					}, nil)

				return fields{
					ur: mockur,
				}
			},
			wantErr: false,
		},
		"error: not found": {
			args: args{
				ctx: context.Background(),
				req: &GetUserRequest{
					ID: 1,
				},
			},
			want:  GetUserResponse{},
			want1: http.StatusNotFound,
			setupFields: func(t *testing.T, args args, _ GetUserResponse, _ int) fields {
				ctrl := gomock.NewController(t)
				mockur := mock_repository.NewMockUserRepository(ctrl)
				mockur.
					EXPECT().
					FindByID(args.ctx, args.req.ID).
					Return(domain.User{}, errors.New(usecases.ECNotFound))

				return fields{
					ur: mockur,
				}
			},
			wantErr: true,
		},
		"error: internal": {
			args: args{
				ctx: context.Background(),
				req: &GetUserRequest{
					ID: 1,
				},
			},
			want: GetUserResponse{},
			setupFields: func(t *testing.T, args args, _ GetUserResponse, _ int) fields {
				ctrl := gomock.NewController(t)
				mockur := mock_repository.NewMockUserRepository(ctrl)
				mockur.
					EXPECT().
					FindByID(args.ctx, args.req.ID).
					Return(domain.User{}, errors.New(usecases.ECInternal))

				return fields{
					ur: mockur,
				}
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t, tt.args, tt.want, tt.want1)
			c := NewUserController(f.ur)
			got, err := c.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func Test_userControllerImpl_PostUser(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		req *PostUserRequest
	}
	type fields struct {
		ur repository.UserRepository
	}
	type setupFieldsFunc func(t *testing.T, args args, _ PostUserResponse, _ int) fields
	tests := map[string]struct {
		args        args
		want        PostUserResponse
		want1       int
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		"success": {
			args: args{
				ctx: context.Background(),
				req: &PostUserRequest{
					Name:  "test",
					Email: "test@example.com",
				},
			},
			want: PostUserResponse{
				ID:    1,
				Name:  "test",
				Email: "test@example.com",
			},
			want1: http.StatusCreated,
			setupFields: func(t *testing.T, args args, _ PostUserResponse, _ int) fields {
				ctrl := gomock.NewController(t)
				mockur := mock_repository.NewMockUserRepository(ctrl)
				mockur.
					EXPECT().
					Create(args.ctx, &repository.CreateUserParams{
						Name:  args.req.Name,
						Email: args.req.Email,
					}).
					Return(domain.User{
						ID:    1,
						Name:  args.req.Name,
						Email: args.req.Email,
					}, nil)

				return fields{
					ur: mockur,
				}
			},
			wantErr: false,
		},
		"error: internal": {
			args: args{
				ctx: context.Background(),
				req: &PostUserRequest{
					Name:  "test",
					Email: "test@example.com",
				},
			},
			want: PostUserResponse{},
			setupFields: func(t *testing.T, args args, _ PostUserResponse, _ int) fields {
				ctrl := gomock.NewController(t)
				mockur := mock_repository.NewMockUserRepository(ctrl)
				mockur.
					EXPECT().
					Create(args.ctx, &repository.CreateUserParams{
						Name:  args.req.Name,
						Email: args.req.Email,
					}).
					Return(domain.User{}, errors.New(usecases.ECInternal))

				return fields{
					ur: mockur,
				}
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t, tt.args, tt.want, tt.want1)
			c := NewUserController(f.ur)
			got, err := c.PostUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}
