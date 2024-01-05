// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_devops_build_pipeline", DevopsBuildPipelineDataSource())
	tfresource.RegisterDatasource("oci_devops_build_pipeline_stage", DevopsBuildPipelineStageDataSource())
	tfresource.RegisterDatasource("oci_devops_build_pipeline_stages", DevopsBuildPipelineStagesDataSource())
	tfresource.RegisterDatasource("oci_devops_build_pipelines", DevopsBuildPipelinesDataSource())
	tfresource.RegisterDatasource("oci_devops_build_run", DevopsBuildRunDataSource())
	tfresource.RegisterDatasource("oci_devops_build_runs", DevopsBuildRunsDataSource())
	tfresource.RegisterDatasource("oci_devops_connection", DevopsConnectionDataSource())
	tfresource.RegisterDatasource("oci_devops_connections", DevopsConnectionsDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_artifact", DevopsDeployArtifactDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_artifacts", DevopsDeployArtifactsDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_environment", DevopsDeployEnvironmentDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_environments", DevopsDeployEnvironmentsDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_pipeline", DevopsDeployPipelineDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_pipelines", DevopsDeployPipelinesDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_stage", DevopsDeployStageDataSource())
	tfresource.RegisterDatasource("oci_devops_deploy_stages", DevopsDeployStagesDataSource())
	tfresource.RegisterDatasource("oci_devops_deployment", DevopsDeploymentDataSource())
	tfresource.RegisterDatasource("oci_devops_deployments", DevopsDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_devops_project", DevopsProjectDataSource())
	tfresource.RegisterDatasource("oci_devops_projects", DevopsProjectsDataSource())
	tfresource.RegisterDatasource("oci_devops_repo_file_line", DevopsRepoFileLineDataSource())
	tfresource.RegisterDatasource("oci_devops_repositories", DevopsRepositoriesDataSource())
	tfresource.RegisterDatasource("oci_devops_repository", DevopsRepositoryDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_archive_content", DevopsRepositoryArchiveContentDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_author", DevopsRepositoryAuthorDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_authors", DevopsRepositoryAuthorsDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_commit", DevopsRepositoryCommitDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_commits", DevopsRepositoryCommitsDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_diff", DevopsRepositoryDiffDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_diffs", DevopsRepositoryDiffsDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_file_diff", DevopsRepositoryFileDiffDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_file_line", DevopsRepositoryFileLineDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_mirror_record", DevopsRepositoryMirrorRecordDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_mirror_records", DevopsRepositoryMirrorRecordsDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_object", DevopsRepositoryObjectDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_object_content", DevopsRepositoryObjectContentDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_path", DevopsRepositoryPathDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_paths", DevopsRepositoryPathsDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_ref", DevopsRepositoryRefDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_refs", DevopsRepositoryRefsDataSource())
	tfresource.RegisterDatasource("oci_devops_trigger", DevopsTriggerDataSource())
	tfresource.RegisterDatasource("oci_devops_triggers", DevopsTriggersDataSource())
	tfresource.RegisterDatasource("oci_devops_repository_mirrorrecord", DevopsRepositoryMirrorrecordDataSource())
}
