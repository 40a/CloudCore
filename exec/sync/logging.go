package sync

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Sirupsen/logrus"
	"github.com/CloudPac/CloudCore/db"
	"github.com/CloudPac/CloudCore/exec/types"
)

func start(t types.SyncJob) {
	t.Job.Status = "running"
	t.Job.Started = time.Now()

	d := bson.M{
		"$set": bson.M{
			"status":  t.Job.Status,
			"failed":  false,
			"started": t.Job.Started,
		},
	}

	if err := db.Jobs().UpdateId(t.Job.ID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Status": t.Job.Status,
			"Error":  err,
		}).Errorln("Failed to update job status")
	}
}

func status(t types.SyncJob, s string) {
	t.Job.Status = s
	d := bson.M{
		"$set": bson.M{
			"status": t.Job.Status,
		},
	}

	if err := db.Jobs().UpdateId(t.Job.ID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Status": t.Job.Status,
			"Error":  err,
		}).Errorln("Failed to update job status")
	}
}

func jobFail(t types.SyncJob) {
	t.Job.Status = "failed"
	t.Job.Finished = time.Now()
	t.Job.Failed = true

	//get elapsed time in minutes
	diff := t.Job.Finished.Sub(t.Job.Started)

	d := bson.M{
		"$set": bson.M{
			"status":          t.Job.Status,
			"failed":          t.Job.Failed,
			"finished":        t.Job.Finished,
			"elapsed":         diff.Minutes(),
			"result_stdout":   t.Job.ResultStdout,
			"job_explanation": t.Job.JobExplanation,
			"job_args":        t.Job.JobARGS,
			"job_env":         t.Job.JobENV,
			"job_cwd":         t.Job.JobCWD,
		},
	}

	if err := db.Jobs().UpdateId(t.Job.ID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Status": t.Job.Status,
			"Error":  err,
		}).Errorln("Failed to update job status")
	}

	updateProject(t)
}

func jobCancel(t types.SyncJob) {
	t.Job.Status = "canceled"
	t.Job.Finished = time.Now()
	t.Job.Failed = false

	//get elapsed time in minutes
	diff := t.Job.Finished.Sub(t.Job.Started)

	d := bson.M{
		"$set": bson.M{
			"status":          t.Job.Status,
			"cancel_flag":     true,
			"failed":          t.Job.Failed,
			"finished":        t.Job.Finished,
			"elapsed":         diff.Minutes(),
			"result_stdout":   "stdout capture is missing",
			"job_explanation": "Job Cancelled",
			"job_args":        t.Job.JobARGS,
			"job_env":         t.Job.JobENV,
			"job_cwd":         t.Job.JobCWD,
		},
	}

	if err := db.Jobs().UpdateId(t.Job.ID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Status": t.Job.Status,
			"Error":  err,
		}).Errorln("Failed to update job status")
	}

	updateProject(t)
}

func jobError(t types.SyncJob) {
	t.Job.Status = "error"
	t.Job.Finished = time.Now()
	t.Job.Failed = true

	//get elapsed time in minutes
	diff := t.Job.Finished.Sub(t.Job.Started)

	d := bson.M{
		"$set": bson.M{
			"status":          t.Job.Status,
			"failed":          t.Job.Failed,
			"finished":        t.Job.Finished,
			"elapsed":         diff.Minutes(),
			"result_stdout":   t.Job.ResultStdout,
			"job_explanation": t.Job.JobExplanation,
			"job_args":        t.Job.JobARGS,
			"job_env":         t.Job.JobENV,
			"job_cwd":         t.Job.JobCWD,
		},
	}

	if err := db.Jobs().UpdateId(t.Job.ID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Status": t.Job.Status,
			"Error":  err,
		}).Errorln("Failed to update job status")
	}

	updateProject(t)
}

func jobSuccess(t types.SyncJob) {
	t.Job.Status = "successful"
	t.Job.Finished = time.Now()
	t.Job.Failed = false

	//get elapsed time in minutes
	diff := t.Job.Finished.Sub(t.Job.Started)

	d := bson.M{
		"$set": bson.M{
			"status":          t.Job.Status,
			"failed":          t.Job.Failed,
			"finished":        t.Job.Finished,
			"elapsed":         diff.Minutes(),
			"result_stdout":   t.Job.ResultStdout,
			"job_explanation": t.Job.JobExplanation,
			"job_args":        t.Job.JobARGS,
			"job_env":         t.Job.JobENV,
			"job_cwd":         t.Job.JobCWD,
		},
	}

	if err := db.Jobs().UpdateId(t.Job.ID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Status": t.Job.Status,
			"Error":  err,
		}).Errorln("Failed to update job status")
	}

	updateProject(t)
}

func updateProject(t types.SyncJob) {
	d := bson.M{
		"$set": bson.M{
			"last_updated":       t.Job.Finished,
			"last_update_failed": t.Job.Failed,
			"status":             t.Job.Status,
		},
	}

	if err := db.Projects().UpdateId(t.ProjectID, d); err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Errorln("Failed to update project")
	}
}
