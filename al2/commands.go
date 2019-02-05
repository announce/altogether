package al2

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:  "jira-attachment-stats",
		Usage: "Shows Jira's attachment statistics",
		Description: `
    Shows statistics of attachments uploaded to Jira.
    Note that the command only can retrieve information accessible to the specified user.
`,
		//Action: jira.PrintAttachmentStats,
		Flags: adminFlags,
	},
	{
		Name:  "confluence-attachment-stats",
		Usage: "Shows Confluence's attachment statistics",
		Description: `
    Shows statistics of attachments uploaded to Confluence.
    Note that the command only can retrieve information accessible to the specified user.
`,
		//Action: confluence.PrintAttachmentStats,
		Flags: adminFlags,
	},
}

var adminFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "dry-run",
		Usage:  "Specify your cloud site's URL such as 'https://example.atlassian.net'",
		EnvVar: "DRY_RUN",
	},
	cli.StringFlag{
		Name:   "admin-email",
		Usage:  "Specify an email address of site administrator which you can check at https://id.atlassian.com/manage-profile/email",
		EnvVar: "AM_ADMIN_EMAIL",
	},
	cli.StringFlag{
		Name:   "api-token",
		Usage:  "Specify an API token issued at https://id.atlassian.com/manage/api-tokens",
		EnvVar: "AM_API_TOKEN",
	},
}
