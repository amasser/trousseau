package trousseau

import (
	"fmt"

	"github.com/crowdmob/goamz/aws"
	"github.com/oleiade/trousseau/pkg/remote/gist"
	"github.com/oleiade/trousseau/pkg/remote/s3"
	"github.com/oleiade/trousseau/pkg/remote/ssh"
)

// Uploader is an interface representing the capacity to upload
// files to remote services or servers.
type Uploader interface {
	Upload(path string) error
}

// Downloader is an interface representing the capacity to upload
// files to remote services or servers.
type Downloader interface {
	Download(path string) error
}

type UploadDownloader interface {
	Uploader
	Downloader
}

// S3Remote allows uploading and downloading files to and from Amazon S3 service.
// It implements the UploadDownloader interface.
type S3Remote struct {
	storage *s3.S3Storage
}

// NewS3Remote generates a S3Remote
func NewS3Remote(region, accessKey, secretKey, bucket string) (*S3Remote, error) {
	AWSRegion, ok := aws.Regions[region]
	if !ok {
		return nil, fmt.Errorf("Invalid aws region supplied %s", region)
	}

	uploader := &S3Remote{
		storage: s3.NewS3Storage(
			aws.Auth{AccessKey: accessKey, SecretKey: secretKey},
			bucket,
			AWSRegion,
		),
	}

	return uploader, nil
}

// Upload executes the whole process of pushing
// the trousseau data store file to s3 remote storage
// using the provided environment.
func (s *S3Remote) Upload(path string) error {
	err := s.storage.Connect()
	if err != nil {
		return fmt.Errorf("Unable to connect to S3")
	}

	err = s.storage.Push(GetStorePath(), path)
	if err != nil {
		return err
	}

	return nil
}

// Download executes the whole process of pulling
// the trousseau data store file from s3 remote storage
// using the provided environment.
func (s *S3Remote) Download(path string) error {
	err := s.storage.Connect()
	if err != nil {
		return fmt.Errorf("Unable to connect to S3")
	}

	err = s.storage.Pull(path, GetStorePath())
	if err != nil {
		return err
	}

	return nil
}

// SCPRemote allows uploading and download files to and from a remote server using SSH.
// It implements the UploadDownloader interface.
type SCPRemote struct {
	storage *ssh.ScpStorage
}

// NewSCPRemote generates a SCPRemote
func NewSCPRemote(host, port, user, password, privateKey string) *SCPRemote {
	return &SCPRemote{
		storage: ssh.NewScpStorage(
			host,
			port,
			user,
			password,
			privateKey,
		),
	}
}

// Upload executes the whole process of pushing
// the trousseau data store file to scp remote storage
// using the provided environment.
func (s *SCPRemote) Upload(path string) (err error) {
	err = s.storage.Connect()
	if err != nil {
		return err
	}

	err = s.storage.Push(GetStorePath(), path)
	if err != nil {
		return err
	}

	return nil
}

// Download executes the whole process of downloading
// the trousseau data store file to scp remote storage
// using the provided environment.
func (s *SCPRemote) Download(path string) (err error) {
	err = s.storage.Connect()
	if err != nil {
		return err
	}

	err = s.storage.Pull(path, GetStorePath())
	if err != nil {
		return err
	}

	return nil
}

// GistRemote allows uploading and downloading files to and from Github's Gist service.
// It implements the UploadDownloader interface.
type GistRemote struct {
	storage *gist.GistStorage
}

// NewGistRemote generates a GistRemote
func NewGistRemote(user, token string) *GistRemote {
	return &GistRemote{
		storage: gist.NewGistStorage(user, token),
	}
}

// Upload executes the whole process of pushing
// the trousseau data store file to gist remote storage
// using the provided dsn informations.
func (g *GistRemote) Upload(path string) error {
	g.storage.Connect()

	err := g.storage.Push(GetStorePath(), path)
	if err != nil {
		return err
	}

	return nil
}

// Download executes the whole process of downloading
// the trousseau data store file to gist remote storage
// using the provided dsn informations.
func (g *GistRemote) Download(path string) error {
	g.storage.Connect()

	err := g.storage.Pull(path, GetStorePath())
	if err != nil {
		return err
	}

	return nil
}