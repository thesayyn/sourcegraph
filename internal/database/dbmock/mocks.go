package dbmock

//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i DB -o mock_db.go

//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i AccessTokenStore -o mock_access_tokens.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i RepoStore -o mock_repos.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i SavedSearchStore -o mock_saved_searches.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i NamespaceStore -o mock_namespaces.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i OrgMemberStore -o mock_org_members.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i OrgStore -o mock_orgs.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i UserStore -o mock_users.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i SettingsStore -o mock_settings.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i UserCredentialsStore -o mock_user_credentials.go
//go:generate ../../../dev/mockgen.sh github.com/sourcegraph/sourcegraph/internal/database -i UserEmailsStore -o mock_user_emails.go
