package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "webhook plugin"
	app.Usage = "webhook plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "webhook",
			Usage:  "webhook url",
			EnvVar: "PLUGIN_WEBHOOK",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "token",
			EnvVar: "PLUGIN_TOKEN",
		},
		cli.BoolFlag{
			Name:   "skip-verify",
			Usage:  "skip tls verification",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},
		cli.StringFlag{
			Name:   "repo.scm",
			Usage:  "repository scm",
			EnvVar: "DRONE_REPO_SCM",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.link",
			Usage:  "repository link",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "repo.avatar",
			Usage:  "repository avatar",
			EnvVar: "DRONE_REPO_AVATAR",
		},
		cli.StringFlag{
			Name:   "repo.branch",
			Usage:  "repository branch",
			EnvVar: "DRONE_REPO_BRANCH",
		},
		cli.BoolFlag{
			Name:   "repo.private",
			Usage:  "repository private",
			EnvVar: "DRONE_REPO_PRIVATE",
		},
		cli.BoolFlag{
			Name:   "repo.trusted",
			Usage:  "repository trusted",
			EnvVar: "DRONE_REPO_TRUSTED",
		},
		cli.StringFlag{
			Name:   "commit.url",
			Usage:  "git commit url",
			EnvVar: "DRONE_COMMIT_URL",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "git commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "git commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "build.deploy",
			Usage:  "build deployment target",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.Int64Flag{
			Name:   "job.number",
			Usage:  "job number",
			EnvVar: "DRONE_JOB_NUMBER",
		},
		cli.StringFlag{
			Name:   "job.status",
			Usage:  "job status",
			EnvVar: "DRONE_JOB_STATUS",
		},
		cli.StringFlag{
			Name:   "job.error",
			Usage:  "job error",
			EnvVar: "DRONE_JOB_ERROR",
		},
		cli.Int64Flag{
			Name:   "job.exit.code",
			Usage:  "job exit code",
			EnvVar: "DRONE_JOB_EXIT_CODE",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.Int64Flag{
			Name:   "job.finished",
			Usage:  "job finished",
			EnvVar: "DRONE_JOB_FINISHED",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Repo: Repo{
			Scm:      c.String("repo.scm"),
			Owner:    c.String("repo.owner"),
			Name:     c.String("repo.name"),
			Link:     c.String("repo.link"),
			Avatar:   c.String("repo.avatar"),
			Branch:   c.String("repo.branch"),
			Private:  c.Bool("repo.private"),
			Trusted:  c.Bool("repo.trusted"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Link:     c.String("build.link"),
			Deploy:   c.String("build.deploy"),
			Created:  c.Int64("build.created"),
			Started:  c.Int64("build.started"),
			Finished: c.Int64("build.finished"),
			Url:      c.String("commit.url"),
			Commit:   c.String("commit.sha"),
			Ref:      c.String("commit.ref"),
			Branch:   c.String("commit.branch"),
			Clink:    c.String("commit.link"),
			Message:  c.String("commit.message"),
			Author:   c.String("commit.author"),
			Email:    c.String("commit.author.email"),
			Avatar:   c.String("commit.author.avatar"),
		},
		Job: Job{
			Number:   c.Int("job.number"),
			Status:   c.String("job.status"),
			Error:    c.String("job.error"),
			Code:     c.Int("job.exit.code"),
			Started:  c.Int64("job.started"),
			Finished: c.Int64("job.finished"),
		},
		Config: Config{
			Webhook:    c.String("webhook"),
			Token:      c.String("token"),
			SkipVerify: c.Bool("skip-verify"),
		},
	}

	return plugin.Exec()
}
