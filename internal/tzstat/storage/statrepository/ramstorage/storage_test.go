package ramstorage

import (
	"reflect"
	"sync"
	"testing"
)

func Test_storage_getEventsWithDatetime(t *testing.T) {
	type fields struct {
		mu   sync.RWMutex
		data []event
	}
	type args struct {
		datetimeFrom int64
		datetimeTo   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []event
	}{
		{
			name: "",
			fields: fields{
				mu:   sync.RWMutex{},
				data: nil,
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{},
		},
		{
			name: "",
			fields: fields{
				mu:   sync.RWMutex{},
				data: []event{},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{},
		},
		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 1,
						value:    1,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{},
		},
		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 12,
						value:    12,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 12,
					value:    12,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 15,
						value:    15,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 15,
					value:    15,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 12,
						value:    12,
					},
					{
						datetime: 15,
						value:    15,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 12,
					value:    12,
				},
				{
					datetime: 15,
					value:    15,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 13,
						value:    13,
					},
					{
						datetime: 15,
						value:    15,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 13,
					value:    13,
				},
				{
					datetime: 15,
					value:    15,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 13,
						value:    13,
					},
					{
						datetime: 14,
						value:    14,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 13,
					value:    13,
				},
				{
					datetime: 14,
					value:    14,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 13,
						value:    13,
					},
					{
						datetime: 14,
						value:    14,
					},
					{
						datetime: 17,
						value:    17,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 13,
					value:    13,
				},
				{
					datetime: 14,
					value:    14,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 13,
						value:    13,
					},
					{
						datetime: 13,
						value:    13,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 13,
					value:    13,
				},
				{
					datetime: 13,
					value:    13,
				},
			},
		},

		{
			name: "",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{
					{
						datetime: 13,
						value:    13,
					},
					{
						datetime: 13,
						value:    13,
					},
					{
						datetime: 14,
						value:    14,
					},
					{
						datetime: 17,
						value:    17,
					},
				},
			},
			args: args{
				datetimeFrom: 12,
				datetimeTo:   15,
			},
			want: []event{
				{
					datetime: 13,
					value:    13,
				},
				{
					datetime: 13,
					value:    13,
				},
				{
					datetime: 14,
					value:    14,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				mu:   tt.fields.mu,
				data: tt.fields.data,
			}
			if got := s.getEventsWithDatetime(tt.args.datetimeFrom, tt.args.datetimeTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEventsWithDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storage_insert(t *testing.T) {
	type fields struct {
		mu   sync.RWMutex
		data []event
	}
	type args struct {
		datetime int64
		value    int64
	}
	type result struct {
		data []event
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result result
	}{
		{
			name: "insert to nil store",
			fields: fields{
				mu:   sync.RWMutex{},
				data: nil,
			},
			args: args{
				datetime: 13,
				value:    13,
			},
			result: result{
				data: []event{{
					datetime: 13,
					value:    13,
				}},
			},
		},
		{
			name: "insert to empty store",
			fields: fields{
				mu:   sync.RWMutex{},
				data: []event{},
			},
			args: args{
				datetime: 13,
				value:    13,
			},
			result: result{
				data: []event{{
					datetime: 13,
					value:    13,
				}},
			},
		},

		{
			name: "insert to not empty store",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{{
					datetime: 12,
					value:    12,
				}},
			},
			args: args{
				datetime: 13,
				value:    13,
			},
			result: result{
				data: []event{
					{
						datetime: 12,
						value:    12,
					}, {
						datetime: 13,
						value:    13,
					}},
			},
		},

		{
			name: "insert to not empty store",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{{
					datetime: 12,
					value:    12,
				}},
			},
			args: args{
				datetime: 11,
				value:    11,
			},
			result: result{
				data: []event{
					{
						datetime: 11,
						value:    11,
					}, {
						datetime: 12,
						value:    12,
					}},
			},
		},

		{
			name: "insert to not empty store",
			fields: fields{
				mu: sync.RWMutex{},
				data: []event{{
					datetime: 12,
					value:    12,
				}, {
					datetime: 15,
					value:    15,
				}},
			},
			args: args{
				datetime: 13,
				value:    13,
			},
			result: result{
				data: []event{
					{
						datetime: 12,
						value:    12,
					}, {
						datetime: 13,
						value:    13,
					}, {
						datetime: 15,
						value:    15,
					}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				mu:   tt.fields.mu,
				data: tt.fields.data,
			}
			s.insert(tt.args.datetime, tt.args.value)

			if !reflect.DeepEqual(s.data, tt.result.data) {
				t.Errorf("insert() storage.data = %v, want %v\"", s.data, tt.result.data)
			}
		})
	}
}
