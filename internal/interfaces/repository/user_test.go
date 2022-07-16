package repository

import (
	"context"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/domain/model"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/google/go-cmp/cmp"
)

func Test_userRepositoryImpl_FindByID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		id  int
	}
	type fields struct {
		UserClient *ent.UserClient
	}
	type setupFieldsFunc func(t *testing.T, args args, want model.User) fields
	tests := map[string]struct {
		args        args
		want        model.User
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		"success": {
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.User{
				ID:    1,
				Name:  "test",
				Email: "test@example.com",
			},
			setupFields: func(t *testing.T, args args, want model.User) fields {
				uc := newEntClient(t).User
				insertUser(args.ctx, t, uc, args.id, want.Name, want.Email)

				return fields{uc}
			},
			wantErr: false,
		},
		"user not found": {
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.User{},
			setupFields: func(t *testing.T, args args, want model.User) fields {
				uc := newEntClient(t).User
				insertUser(args.ctx, t, uc, args.id+1, "test", "test@example.com")

				return fields{uc}
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t, tt.args, tt.want)
			r := NewUserRepository(f.UserClient)
			got, err := r.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}
