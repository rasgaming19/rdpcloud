//go:build windows && amd64

package fileio

import (
	"context"

	fileioServicePb "github.com/s77rt/rdpcloud/proto/go/services/fileio"
	fileioApi "github.com/s77rt/rdpcloud/server/go/api/fileio"
)

func (s *Server) GetQuotaState(ctx context.Context, in *fileioServicePb.GetQuotaStateRequest) (*fileioServicePb.GetQuotaStateResponse, error) {
	quotaState, err := fileioApi.GetQuotaState(in.GetVolume())
	if err != nil {
		return nil, err
	}

	return &fileioServicePb.GetQuotaStateResponse{
		QuotaState: quotaState,
	}, nil
}

func (s *Server) SetQuotaState(ctx context.Context, in *fileioServicePb.SetQuotaStateRequest) (*fileioServicePb.SetQuotaStateResponse, error) {
	if err := fileioApi.SetQuotaState(in.GetVolume(), in.GetQuotaState()); err != nil {
		return nil, err
	}

	return &fileioServicePb.SetQuotaStateResponse{}, nil
}

func (s *Server) GetDefaultQuota(ctx context.Context, in *fileioServicePb.GetDefaultQuotaRequest) (*fileioServicePb.GetDefaultQuotaResponse, error) {
	defaultQuota, err := fileioApi.GetDefaultQuota(in.GetVolume())
	if err != nil {
		return nil, err
	}

	return &fileioServicePb.GetDefaultQuotaResponse{
		DefaultQuota: defaultQuota,
	}, nil
}

func (s *Server) SetDefaultQuota(ctx context.Context, in *fileioServicePb.SetDefaultQuotaRequest) (*fileioServicePb.SetDefaultQuotaResponse, error) {
	if err := fileioApi.SetDefaultQuota(in.GetVolume(), in.GetDefaultQuota()); err != nil {
		return nil, err
	}

	return &fileioServicePb.SetDefaultQuotaResponse{}, nil
}

func (s *Server) GetUsersQuotaEntries(ctx context.Context, in *fileioServicePb.GetUsersQuotaEntriesRequest) (*fileioServicePb.GetUsersQuotaEntriesResponse, error) {
	quotaEntries, err := fileioApi.GetUsersQuotaEntries(in.GetVolume())
	if err != nil {
		return nil, err
	}

	return &fileioServicePb.GetUsersQuotaEntriesResponse{
		QuotaEntries: quotaEntries,
	}, nil
}

func (s *Server) GetUserQuotaEntry(ctx context.Context, in *fileioServicePb.GetUserQuotaEntryRequest) (*fileioServicePb.GetUserQuotaEntryResponse, error) {
	quotaEntry, err := fileioApi.GetUserQuotaEntry(in.GetVolume(), in.GetUser())
	if err != nil {
		return nil, err
	}

	return &fileioServicePb.GetUserQuotaEntryResponse{
		QuotaEntry: quotaEntry,
	}, nil
}

func (s *Server) SetUserQuotaEntry(ctx context.Context, in *fileioServicePb.SetUserQuotaEntryRequest) (*fileioServicePb.SetUserQuotaEntryResponse, error) {
	if err := fileioApi.SetUserQuotaEntry(in.GetVolume(), in.GetUser(), in.GetQuotaEntry()); err != nil {
		return nil, err
	}

	return &fileioServicePb.SetUserQuotaEntryResponse{}, nil
}
