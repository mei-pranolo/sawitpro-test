package repository

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	m "github.com/SawitProRecruitment/UserService/types"
	gomock "go.uber.org/mock/gomock"
)

func TestRepository_GetEstateByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		client sqlmock.Sqlmock
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantEstate m.Estate
		wantErr    bool
		mock       func(ctrl *gomock.Controller) sqlmock.Sqlmock
	}{
		{
			name: "when all good, return estate data",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx: context.Background(),
				id:  "aaa",
			},
			wantEstate: m.Estate{ID: "aaa", Length: 2, Width: 2},
			wantErr:    false,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				rows := sqlmock.NewRows([]string{"id", "length", "width"}).AddRow("aaa", 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM estate WHERE id = $1")).WithArgs("aaa").WillReturnRows(rows)
				return mock
			},
		},
		{
			name: "when database return error, return error",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx: context.Background(),
				id:  "aaa",
			},
			wantEstate: m.Estate{},
			wantErr:    true,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM estate WHERE id = $1")).WithArgs("aaa").WillReturnError(errors.New(""))
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockControl := gomock.NewController(t)
			if tt.mock != nil {
				tt.fields.client = tt.mock(mockControl)
			}
			r := &Repository{
				Db: db,
			}
			gotEstate, err := r.GetEstateByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetEstateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEstate, tt.wantEstate) {
				t.Errorf("Repository.GetEstateByID() = %v, want %v", gotEstate, tt.wantEstate)
			}
		})
	}
}

func TestRepository_CreateEstate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		client sqlmock.Sqlmock
	}

	type args struct {
		ctx    context.Context
		length int
		width  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  string
		wantErr bool
		mock    func(ctrl *gomock.Controller) sqlmock.Sqlmock
	}{
		{
			name: "when all good, return estate id",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:    context.Background(),
				length: 2,
				width:  2,
			},
			wantId:  "aaa",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("aaa")
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO estate (length, width) VALUES($1, $2) RETURNING id")).WithArgs(2, 2).WillReturnRows(rows)
				return mock
			},
		},
		{
			name: "when database return error, return error",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:    context.Background(),
				length: 2,
				width:  2,
			},
			wantId:  "",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO estate (length, width) VALUES($1, $2) RETURNING id")).WithArgs(2, 2).WillReturnError(errors.New("new"))
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockControl := gomock.NewController(t)
			if tt.mock != nil {
				tt.fields.client = tt.mock(mockControl)
			}
			r := &Repository{
				Db: db,
			}

			gotId, err := r.CreateEstate(tt.args.ctx, tt.args.length, tt.args.width)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Repository.CreateEstate() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestRepository_CreateTree(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		client sqlmock.Sqlmock
	}

	type args struct {
		ctx      context.Context
		estateID string
		tree     m.Tree
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  string
		wantErr bool
		mock    func(ctrl *gomock.Controller) sqlmock.Sqlmock
	}{
		{
			name: "when all good, return estate id",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 1, Y: 2, Height: 3},
			},
			wantId:  "aaa",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO tree (estate_id, x, y, height) VALUES($1, $2, $3, $4)")).WithArgs("aaa", 1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
				return mock
			},
		},
		{
			name: "when database return error, return error",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
				tree:     m.Tree{X: 1, Y: 2, Height: 3},
			},
			wantId:  "",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO tree (estate_id, x, y, height) VALUES($1, $2, $3, $4)")).WithArgs("aaa", 1, 2, 3).WillReturnError(errors.New("create tree"))
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockControl := gomock.NewController(t)
			if tt.mock != nil {
				tt.fields.client = tt.mock(mockControl)
			}
			r := &Repository{
				Db: db,
			}

			gotId, err := r.CreateTree(tt.args.ctx, tt.args.estateID, tt.args.tree)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Repository.CreateTree() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestRepository_GetTree(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		client sqlmock.Sqlmock
	}

	type args struct {
		ctx      context.Context
		estateID string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTrees []m.Tree
		wantErr   bool
		mock      func(ctrl *gomock.Controller) sqlmock.Sqlmock
	}{
		{
			name: "when all good, return tree data",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantTrees: []m.Tree{{X: 1, Y: 1, Height: 1}, {X: 2, Y: 2, Height: 2}},
			wantErr:   false,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				rows := sqlmock.NewRows([]string{"x", "y", "height"}).AddRow(1, 1, 1).AddRow(2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT x,y,height FROM tree WHERE estate_id = $1")).WithArgs("aaa").WillReturnRows(rows)
				return mock
			},
		},
		{
			name: "when database return error, return error",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantTrees: []m.Tree(nil),
			wantErr:   true,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT x,y,height FROM tree WHERE estate_id = $1")).WithArgs("aaa").WillReturnError(errors.New("tree"))
				return mock
			},
		},
		{
			name: "when some data is corrupted, skip it",
			fields: fields{
				client: mock,
			},
			args: args{
				ctx:      context.Background(),
				estateID: "aaa",
			},
			wantTrees: []m.Tree{{X: 1, Y: 1, Height: 1}, {X: 2, Y: 2, Height: 2}},
			wantErr:   false,
			mock: func(ctrl *gomock.Controller) sqlmock.Sqlmock {
				rows := sqlmock.NewRows([]string{"x", "y", "height"}).AddRow(1, 1, 1).AddRow(2, 2, 2).RowError(3, errors.New("row"))
				mock.ExpectQuery(regexp.QuoteMeta("SELECT x,y,height FROM tree WHERE estate_id = $1")).WithArgs("aaa").WillReturnRows(rows)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockControl := gomock.NewController(t)
			if tt.mock != nil {
				tt.fields.client = tt.mock(mockControl)
			}
			r := &Repository{
				Db: db,
			}

			gotTrees, err := r.GetTree(tt.args.ctx, tt.args.estateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTrees, tt.wantTrees) {
				t.Errorf("Repository.GetTree() = %v, want %v", gotTrees, tt.wantTrees)
			}
		})
	}
}
