package atlasproject

import (
	"context"
	"errors"

	mdbv1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/controller/atlas"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/controller/workflow"
	"go.mongodb.org/atlas/mongodbatlas"
)

// ensureProjectExists creates the project if it doesn't exist yet. Returns the project ID
func ensureProjectExists(ctx *workflow.Context, connection atlas.Connection, project *mdbv1.AtlasProject) (string, workflow.Result) {
	client, err := atlas.Client(connection, ctx.Log)
	if err != nil {
		return "", workflow.Terminate(workflow.Internal, err.Error())
	}
	// Try to find the project
	p, _, err := client.Projects.GetOneProjectByName(context.Background(), project.Spec.Name)
	if err != nil {
		var apiError *mongodbatlas.ErrorResponse
		if errors.As(err, &apiError) && apiError.ErrorCode == atlas.NotInGroup {
			// Project doesn't exist? Try to create it
			p = &mongodbatlas.Project{
				OrgID: connection.OrgID,
				Name:  project.Spec.Name,
			}
			if p, _, err = client.Projects.Create(context.Background(), p); err != nil {
				return "", workflow.Terminate(workflow.ProjectNotCreatedInAtlas, err.Error())
			}
			ctx.Log.Infow("Created Atlas Project", "name", project.Spec.Name, "id", p.ID)
		} else {
			return "", workflow.Terminate(workflow.ProjectNotCreatedInAtlas, err.Error())
		}
	}

	if p == nil || p.ID == "" {
		ctx.Log.Error("Project or its project ID are empty")
		return "", workflow.Terminate(workflow.Internal, "")
	}
	return p.ID, workflow.OK()
}
