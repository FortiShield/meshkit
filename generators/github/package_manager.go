package github

import (
	"net/url"

	"github.com/khulnasoft/meshkit/generators/models"
	"github.com/khulnasoft/meshkit/utils/walker"
)

type GitHubPackageManager struct {
	PackageName string
	SourceURL   string
}

func (ghpm GitHubPackageManager) GetPackage() (models.Package, error) {
	url, err := url.Parse(ghpm.SourceURL)
	if err != nil {
		err = walker.ErrCloningRepo(err)
		return nil, err
	}
	protocol := url.Scheme

	downloader := NewDownloaderForScheme(protocol, url, ghpm.PackageName)
	if downloader == nil {
		return nil, ErrGenerateGitHubPackage(err, ghpm.PackageName)
	}
	ghPackage, err := downloader.GetContent()
	if err != nil {
		return nil, ErrGenerateGitHubPackage(err, ghpm.PackageName)
	}
	return ghPackage, nil
}
