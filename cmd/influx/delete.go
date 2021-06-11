package main

import (
	"github.com/influxdata/influx-cli/v2/clients/delete"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newDeleteCmd() *cli.Command {
	var params delete.Params
	return &cli.Command{
		Name:        "delete",
		Usage:       "Delete points from InfluxDB",
		Description: "Delete points from InfluxDB, by specify start, end time and a sql like predicate string",
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:    "org-id",
				Usage:   "The ID of the organization that owns the bucket",
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization that owns the bucket",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:    "bucket-id",
				Usage:   "The ID of the bucket to delete from",
				EnvVars: []string{"INFLUX_BUCKET_ID"},
				Value:   &params.BucketID,
			},
			&cli.StringFlag{
				Name:        "bucket",
				Usage:       "The name of the bucket to delete from",
				EnvVars:     []string{"INFLUX_BUCKET_NAME"},
				Destination: &params.BucketName,
			},
			// NOTE: cli has a Timestamp flag we could use to parse the strings immediately on input,
			// but the help-text generation is broken for it.
			&cli.StringFlag{
				Name:        "start",
				Usage:       "The start time in RFC3339Nano format (ex: '2009-01-02T23:00:00Z')",
				Required:    true,
				Destination: &params.Start,
			},
			&cli.StringFlag{
				Name:        "stop",
				Usage:       "The stop time in RFC3339Nano format (ex: '2009-01-02T23:00:00Z')",
				Required:    true,
				Destination: &params.Stop,
			},
			&cli.StringFlag{
				Name:        "predicate",
				Usage:       "sql like predicate string (ex: 'tag1=\"v1\" and (tag2=123)')",
				Aliases:     []string{"p"},
				Destination: &params.Predicate,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			client := delete.Client{CLI: getCLI(ctx), DeleteApi: getAPI(ctx).DeleteApi}
			return client.Delete(ctx.Context, &params)
		},
	}
}